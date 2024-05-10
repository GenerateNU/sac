package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/entities/models"

	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, category models.Category) (*models.Category, error) {
	if err := db.Create(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, utilities.ErrDuplicate
		}
		return nil, err
	}

	return &category, nil
}

func GetCategories(db *gorm.DB, limit int, page int) ([]models.Category, error) {
	var categories []models.Category

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategory(db *gorm.DB, id uuid.UUID) (*models.Category, error) {
	var category models.Category

	if err := db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &category, nil
}

func UpdateCategory(db *gorm.DB, id uuid.UUID, category models.Category) (*models.Category, error) {
	if err := db.Model(&models.Category{}).Where("id = ?", id).Updates(category).First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &category, nil
}

func DeleteCategory(db *gorm.DB, id uuid.UUID) error {
	if result := db.Delete(&models.Category{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}

	return nil
}
