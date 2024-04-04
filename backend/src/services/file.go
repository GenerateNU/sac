package services

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/types"
)

type FileServiceInterface interface {
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

func (f *FileService) GetFile(fileID string) (*models.File, *errors.Error) {
	return nil, nil
}

func (f *FileService) CreateFile(fileBody *models.CreateFileRequestBody, formFile *multipart.FileHeader) (*models.File, *errors.Error) {
	return nil, nil
}

func (f *FileService) DeleteFile(fileID string) *errors.Error {
	return nil
}
