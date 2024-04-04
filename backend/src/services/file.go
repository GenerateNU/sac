package services

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/file"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type FileServiceInterface interface {
	GetFiles(limit string, page string) ([]models.File, *errors.Error)
	GetFile(fileID string) (*models.File, *errors.Error)
	CreateFile(fileBody *models.CreateFileRequestBody, formFile *multipart.FileHeader) (*models.File, *errors.Error)
	DeleteFile(fileID string) *errors.Error
}

type FileService struct {
	types.ServiceParams
}

func NewFileService(serviceParams types.ServiceParams) FileServiceInterface {
	return &FileService{serviceParams}
}

func (f *FileService) GetFiles(limit string, page string) ([]models.File, *errors.Error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	return transactions.GetFiles(f.DB, *limitAsInt, *pageAsInt)
}

func (f *FileService) GetFile(fileID string) (*models.File, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(fileID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetFile(f.DB, *idAsUUID)
}

// TODO: might not be something we want to do (if you delete a file will it remove its owner?)
func (f *FileService) CreateFile(fileBody *models.CreateFileRequestBody, fileHeader *multipart.FileHeader) (*models.File, *errors.Error) {
	if err := f.Validate.Struct(fileBody); err != nil {
		return nil, &errors.FailedToValidateFile
	}

	fileInfo, err := f.Deps.File.UploadFile(fileBody.OwnerType, fileHeader, []file.FileTypes{file.ALL})
	if err != nil {
		return nil, &errors.FailedToUploadFile
	}

	file, err := transactions.CreateFile(f.DB, fileBody.OwnerID, fileBody.OwnerType, *fileInfo)
	if err != nil {
		return nil, &errors.FailedToCreateFile
	}

	return file, nil
}

func (f *FileService) DeleteFile(fileID string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(fileID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteFile(f.DB, *idAsUUID)
}
