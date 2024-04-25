package base

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
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
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return &errors.FileNotFound
		} else {
			return &errors.FailedToDeleteFile
		}
	}

	return nil
}

func UpdateFile(db *gorm.DB, fileID uuid.UUID, fileInfo models.FileInfo) (*models.File, *errors.Error) {
	existingFile, err := GetFile(db, fileID)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&existingFile).Updates(models.File{
		FileName:  fileInfo.FileName,
		FileType:  fileInfo.FileType,
		FileSize:  fileInfo.FileSize,
		FileURL:   fileInfo.FileURL,
		ObjectKey: fileInfo.ObjectKey,
	}).Error; err != nil {
		return nil, &errors.FailedToUpdateFile
	}

	return existingFile, nil
}

func GetFiles(db *gorm.DB, limit, page int) ([]models.File, *errors.Error) {
	var files []models.File

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Find(&files).Error; err != nil {
		return nil, &errors.FailedToGetFiles
	}

	return files, nil
}

func GetFile(db *gorm.DB, fileID uuid.UUID) (*models.File, *errors.Error) {
	var file models.File
	if err := db.First(&file, fileID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.FileNotFound
		} else {
			return nil, &errors.FailedToGetFile
		}
	}

	return &file, nil
}
