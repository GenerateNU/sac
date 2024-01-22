package transactions

import (
	"errors"

	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User

	if err := db.Omit("password_hash").Find(&users).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "failed to get all users")
	}

	return users, nil
}

func DeleteUser(db *gorm.DB, id uint) error {
	var deletedUser models.User

	result := db.Where("id = ?", id).Delete(&deletedUser)
	if result.RowsAffected == 0 {
		if result.Error == nil {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
		}
	}
	return nil
}

func GetUser(db *gorm.DB, id uint) (*models.User, error) {
	var user models.User

	if err := db.Omit("password_hash").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "failed to find tag")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "failed to get user")
	}

	return &user, nil
}
