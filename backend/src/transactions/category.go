package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"

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


func GetCategory(db *gorm.DB, id uint) (*models.Category, *errors.Error) {
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

func UpdateCategory(db *gorm.DB, id uint, category models.Category) (*models.Category, error) {
	var existingCategory models.Category

	if err := db.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "failed to update category")
		}
	} else {
		if existingCategory.ID != id {
			return nil, fiber.NewError(fiber.StatusBadRequest, "category with that name already exists")
		}	
	}
	
	if err := db.Model(&models.Category{}).Where("id = ?", id).Updates(category).First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "failed to find category")
		} else {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "failed to update category")
		}
	}

	return &category, nil
}

func DeleteCategory(db *gorm.DB, id uint) error {
	if result := db.Delete(&models.Category{}, id); result.RowsAffected == 0 {
		if result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to delete category")
		} else {
			return fiber.NewError(fiber.StatusNotFound, "failed to find category")
		}
	}

	return nil
}
