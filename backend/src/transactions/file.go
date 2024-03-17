package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFile(db *gorm.DB, ownerID uuid.UUID, ownerType string, fileInfo models.FileInfo) (*models.File, *errors.Error) {
	file := &models.File{
		OwnerID:   ownerID,
		OwnerType: ownerType,
		FileName:  fileInfo.FileName,
		FileType:  fileInfo.FileType,
		FileSize:  fileInfo.FileSize,
		FileURL:   fileInfo.FileURL,
		ObjectKey: fileInfo.ObjectKey,
	}

	if err := db.Create(file).Error; err != nil {
		return nil, &errors.FailedToCreateFileInDB
	}

	return file, nil
}

func DeleteFile(db *gorm.DB, fileID uuid.UUID) *errors.Error {
	if err := db.Delete(&models.File{}, fileID).Error; err != nil {
		return &errors.FailedToDeleteFile
	}

	return nil
}

func UpdateFile(db *gorm.DB, fileID uuid.UUID, fileInfo models.FileInfo) (*models.File, *errors.Error) {
	var existingFile models.File

	err := db.First(&existingFile, fileID).Error
	if err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.FileNotFound
		} else {
			return nil, &errors.FailedToGetFile
		}
	}

	if err := db.Model(&existingFile).Updates(models.File{
		FileName: fileInfo.FileName,
		FileType: fileInfo.FileType,
		FileSize: fileInfo.FileSize,
		FileURL:  fileInfo.FileURL,
		ObjectKey: fileInfo.ObjectKey,
	}).Error; err != nil {
		return nil, &errors.FailedToUpdateFile
	}

	return &existingFile, nil
}

func GetFile(db *gorm.DB, fileID uuid.UUID) (*models.File, *errors.Error) {
	var file models.File
	if err := db.First(&file, fileID).Error; err != nil {
		return nil, &errors.FailedToGetFile
	}

	return &file, nil
}