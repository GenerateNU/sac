package base

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/file"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type FileServiceInterface interface {
	GetFiles(pageInfo fiberpaginate.PageInfo) ([]models.File, error)
	GetFile(fileID string) (*models.File, error)
	CreateFile(fileBody *CreateFileRequestBody, formFile *multipart.FileHeader) (*models.File, error)
	DeleteFile(fileID string) error
}

type FileService struct {
	types.ServiceParams
}

func NewFileService(serviceParams types.ServiceParams) FileServiceInterface {
	return &FileService{serviceParams}
}

func (f *FileService) GetFiles(pageInfo fiberpaginate.PageInfo) ([]models.File, error) {
	return GetFiles(f.DB, pageInfo)
}

func (f *FileService) GetFile(fileID string) (*models.File, error) {
	idAsUUID, err := utilities.ValidateID(fileID)
	if err != nil {
		return nil, err
	}

	return GetFile(f.DB, *idAsUUID)
}

func (f *FileService) CreateFile(fileBody *CreateFileRequestBody, fileHeader *multipart.FileHeader) (*models.File, error) {
	if err := utilities.Validate(f.Validate, fileBody); err != nil {
		return nil, err
	}

	fileInfo, err := f.Integrations.File.UploadFile(fileBody.OwnerType, fileHeader, []file.FileType{file.ALL})
	if err != nil {
		return nil, err
	}

	file, err := CreateFile(f.DB, fileBody.OwnerID, fileBody.OwnerType, *fileInfo)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *FileService) DeleteFile(fileID string) error {
	idAsUUID, err := utilities.ValidateID(fileID)
	if err != nil {
		return err
	}

	return DeleteFile(f.DB, *idAsUUID)
}
