package services

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/src/aws"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileServiceInterface interface {
	CreateFile(fileBody *models.CreateFileRequestBody, fileHeader *multipart.FileHeader) (*models.File, *errors.Error)
	DeleteFile(fileID string) *errors.Error
	GetFileInfo(fileID string) (*models.File, *errors.Error)
}

type FileService struct {
	DB       *gorm.DB
	Validate *validator.Validate
	AWSClient *aws.AWSClient
}

func NewFileService(db *gorm.DB, validate *validator.Validate, client *aws.AWSClient) FileServiceInterface {
	return &FileService{DB: db, Validate: validate, AWSClient: client}
}

func (f *FileService) GetFileInfo(fileID string) (*models.File, *errors.Error) {
	fileIdAsUUID, err := utilities.ValidateID(fileID)
	if err != nil {
		return nil, &errors.FailedToValidateFile
	}

	return transactions.GetFile(f.DB, *fileIdAsUUID)
}

// TODO: specific file type validation
func (f *FileService) CreateFile(fileBody *models.CreateFileRequestBody, fileHeader *multipart.FileHeader) (*models.File, *errors.Error) {
	if err := f.Validate.Struct(fileBody); err != nil {
		return nil, &errors.FailedToValidateFile
	}
	
	if fileHeader.Size > 5000000 {
		return nil, &errors.InvalidFileSize
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, &errors.FailedToOpenFile
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, &errors.FailedToReadFile
	}

	defer file.Close()

	fileName := generateUniqueFileName(fileHeader.Filename)

	fileURL, clientErr := f.AWSClient.UploadFile(fileBytes, fileName, fileBody.OwnerType)
	if clientErr != nil {
		return nil, clientErr
	}

	fileInfo := models.FileInfo{
		FileName: fileHeader.Filename,
		FileType: fileHeader.Header.Get("Content-Type"),
		FileSize: int(fileHeader.Size),
		FileURL: *fileURL,
		ObjectKey: fileName,
	}

	return transactions.CreateFile(f.DB, fileBody.OwnerID, fileBody.OwnerType, fileInfo)
}

func (f *FileService) DeleteFile(fileID string) *errors.Error {
	fileIdAsUUID, err := utilities.ValidateID(fileID)
	if err != nil {
		return &errors.FailedToValidateFile
	}

	file, err := transactions.GetFile(f.DB, *fileIdAsUUID)
	if err != nil {
		return err
	}

	clientErr := f.AWSClient.DeleteFile(file.FileURL)
	if clientErr != nil {
		return clientErr
	}

	return transactions.DeleteFile(f.DB, *fileIdAsUUID)
}

func generateUniqueFileName(fileName string) string {
	return fmt.Sprintf("%s-%s", uuid.New().String(), fileName)
}