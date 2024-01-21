package transactions

import (
	"errors"

	"github.com/GenerateNU/sac/backend/src/models"

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
			return nil, fiber.NewError(fiber.StatusNotFound, "failed to find tag")
		} else {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "failed to retrieve tag")
		}
	}

	return &tag, nil
}

func DeleteTag(db *gorm.DB, id uint) error {
	if result := db.Delete(&models.Tag{}, id); result.RowsAffected == 0 {
		if result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to delete tag")
		} else {
			return fiber.NewError(fiber.StatusNotFound, "failed to find tag")
		}
	}

	return nil
}
