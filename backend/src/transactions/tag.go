package transactions

import (
	"backend/src/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTag(db *gorm.DB, tag models.Tag) (*models.Tag, error) {
	if err := db.Create(&tag).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "failed to create tag")
	}
	return &tag, nil
}

func GetTag(db *gorm.DB, id uint) (*models.Tag, error) {
	var tag models.Tag

	if err := db.First(&tag, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid tag id")
		} else {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "unable to retrieve tag")
		}
	}

	return &tag, nil
}
