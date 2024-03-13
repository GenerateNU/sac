package services

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/src/aws"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileServiceInterface interface {
	CreateFile(fileBody *models.CreateFileRequestBody, fileHeader *multipart.FileHeader) (*models.File, *errors.Error)
	// DeleteFile(fileID string) *errors.Error
	// GetFile(fileID string) (*models.File, *errors.Error)
}

type FileService struct {
	DB       *gorm.DB
	Validate *validator.Validate
	AWSClient *aws.AWSClient
}

func NewFileService(db *gorm.DB, validate *validator.Validate, client *aws.AWSClient) FileServiceInterface {
	return &FileService{DB: db, Validate: validate, AWSClient: client}
}

// validate size
// validate file type
// validate file name

func (f *FileService) CreateFile(fileBody *models.CreateFileRequestBody, fileHeader *multipart.FileHeader) (*models.File, *errors.Error) {
	if err := f.Validate.Struct(fileBody); err != nil {
		return nil, &errors.FailedToValidateFile
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, &errors.FailedToOpenFile
	}

	fmt.Println(fileBody.OwnerID)
	fmt.Println(fileBody.OwnerType)	
	fmt.Println(fileHeader.Filename)
	fmt.Println(fileHeader.Size)
	
	if fileHeader.Size > 5000000 {
		return nil, &errors.InvalidFileSize
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, &errors.FailedToReadFile
	}

	defer file.Close()

	fileName := generateUniqueFileName(fileHeader.Filename)

	fileURL, clientErr := f.AWSClient.UploadFile(fileBytes, fileName, "test")
	if clientErr != nil {
		return nil, clientErr
	}

	return transactions.CreateFile(f.DB, fileBody.OwnerID, fileBody.OwnerType, *fileURL)
}

func generateUniqueFileName(fileName string) string {
	return fmt.Sprintf("%s-%s", uuid.New().String(), fileName)
}