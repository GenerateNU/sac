package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"

	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, category models.Category) (*models.Category, *errors.Error) {
	if err := db.Create(&category).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.CategoryAlreadyExists}
		} else {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToCreateCategory}
		}
	}

	return &category, nil
}

func GetCategory(db *gorm.DB, id uint) (*models.Category, *errors.Error) {
	var category models.Category

	if err := db.First(&category, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.CategoryNotFound}
		} else {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetCategory}
		}
	}

	return &category, nil
}
