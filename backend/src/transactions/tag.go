package transactions

import (
	"errors"

	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTag(db *gorm.DB, tag models.Tag) (*models.Tag, error) {
	if err := db.Create(&tag).Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}
	
	return &tag, nil
}

func GetTag(db *gorm.DB, id uint) (*models.Tag, error) {
	var tag models.Tag

	if err := db.First(&tag, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		} else {
			return nil, fiber.ErrInternalServerError
		}
	}

	return &tag, nil
}

func UpdateTag(db *gorm.DB, id uint, tag models.Tag) (*models.Tag, error) {
	if err := db.Model(&models.Tag{}).Where("id = ?", id).Updates(tag).First(&tag, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		} else {
			return nil, fiber.ErrInternalServerError
		}
	}

	return &tag, nil

}

func DeleteTag(db *gorm.DB, id uint) error {
	if result := db.Delete(&models.Tag{}, id); result.RowsAffected == 0 {
		if result.Error != nil {
			return fiber.ErrInternalServerError
		} else {
			return fiber.ErrNotFound
		}
	}

	return nil
}
