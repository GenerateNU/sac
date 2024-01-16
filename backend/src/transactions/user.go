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

func UpdateUser(db *gorm.DB, id string, payload models.User) (models.User, error) {
	var existingUser models.User

	if err := db.First(&existingUser, id).Error; err != nil {
		return models.User{}, err
	}

	if err := db.Model(&existingUser).Updates(&payload).Error; err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}
