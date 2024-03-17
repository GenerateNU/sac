package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
)

type FileServiceInterface interface {
	GetFile(fileID string) (*models.File, *errors.Error)
	CreateFile(fileBody models.CreateFileRequestBody) (*models.File, *errors.Error)
	DeleteFile(fileID string) *errors.Error
}

type FileService struct {
	AWSProvider *files.AWSProvider
}
