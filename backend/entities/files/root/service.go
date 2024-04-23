package root

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/file"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
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

	return GetFiles(f.DB, *limitAsInt, *pageAsInt)
}

func (f *FileService) GetFile(fileID string) (*models.File, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(fileID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return GetFile(f.DB, *idAsUUID)
}

func (f *FileService) CreateFile(fileBody *models.CreateFileRequestBody, fileHeader *multipart.FileHeader) (*models.File, *errors.Error) {
	if err := f.Validate.Struct(fileBody); err != nil {
		return nil, &errors.FailedToValidateFile
	}

	fileInfo, err := f.Deps.File.UploadFile(fileBody.OwnerType, fileHeader, []file.FileType{file.ALL})
	if err != nil {
		return nil, &errors.FailedToUploadFile
	}

	file, err := CreateFile(f.DB, fileBody.OwnerID, fileBody.OwnerType, *fileInfo)
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

	return DeleteFile(f.DB, *idAsUUID)
}
