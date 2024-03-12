package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/src/models"

	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, category models.Category) (*models.Category, *errors.Error) {
	if err := db.Create(&category).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &errors.CategoryAlreadyExists
		} else {
			return nil, &errors.FailedToCreateCategory
		}
	}

	return &category, nil
}

func GetCategories(db *gorm.DB, limit int, page int) ([]models.Category, *errors.Error) {
	var categories []models.Category

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, &errors.FailedToGetCategories
	}

	return categories, nil
}

func GetCategory(db *gorm.DB, id uuid.UUID) (*models.Category, *errors.Error) {
	var category models.Category

	if err := db.First(&category, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.CategoryNotFound
		} else {
			return nil, &errors.FailedToGetCategory
		}
	}

	return &category, nil
}

func UpdateCategory(db *gorm.DB, id uuid.UUID, category models.Category) (*models.Category, *errors.Error) {
	if err := db.Model(&models.Category{}).Where("id = ?", id).Updates(category).First(&category, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.TagNotFound
		} else {
			return nil, &errors.FailedToUpdateTag
		}
	}

	return &category, nil
}

func DeleteCategory(db *gorm.DB, id uuid.UUID) *errors.Error {
	if result := db.Delete(&models.Category{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.CategoryNotFound
		} else {
			return &errors.FailedToDeleteCategory
		}
	}

	return nil
}
