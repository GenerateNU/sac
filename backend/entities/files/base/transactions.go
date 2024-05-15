package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/garrettladley/fiberpaginate"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFile(db *gorm.DB, ownerID uuid.UUID, ownerType string, fileInfo models.FileInfo) (*models.File, error) {
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
		return nil, err
	}

	return file, nil
}

func DeleteFile(db *gorm.DB, fileID uuid.UUID) error {
	if err := db.Delete(&models.File{}, fileID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utilities.ErrNotFound
		}
		return err
	}

	return nil
}

func UpdateFile(db *gorm.DB, fileID uuid.UUID, fileInfo models.FileInfo) (*models.File, error) {
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
		return nil, err
	}

	return existingFile, nil
}

func GetFiles(db *gorm.DB, pageInfo fiberpaginate.PageInfo) ([]models.File, error) {
	var files []models.File
	if err := db.Scopes(utilities.IntoScope(pageInfo, db)).Find(&files).Error; err != nil {
		return nil, err
	}

	return files, nil
}

func GetFile(db *gorm.DB, fileID uuid.UUID) (*models.File, error) {
	var file models.File
	if err := db.First(&file, fileID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &file, nil
}
