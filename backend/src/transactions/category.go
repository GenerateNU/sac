package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"

	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, category models.Category) (*models.Category, *errors.Error) {
	var existingCategory models.Category

	if err := db.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
		if !stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: "failed to create category"}
		}
	} else {
		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: "category already exists"}
	}

	if err := db.Create(&category).Error; err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: "failed to create category"}
	}

	return &category, nil
}

func GetCategory(db *gorm.DB, id uint) (*models.Category, *errors.Error) {
	var category models.Category

	if err := db.First(&category, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.Error{StatusCode: fiber.StatusNotFound, Message: "category not found"}
		} else {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: "failed to get category"}
		}
	}

	return &category, nil
}
