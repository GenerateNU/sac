package transactions

import (
	"backend/src/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, category models.Category) (*models.Category, error) {
	if err := db.Create(&category).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "failed to create category")
	}
	return &category, nil
}

func GetCategory(db *gorm.DB, id uint) (*models.Category, error) {
	var category models.Category
	
	if err := db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid category id")
		} else {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "unable to retrieve category")
		}
	}

	return &category, nil
}