package transactions

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User

	if err := db.Unscoped().Omit("password_hash").Find(&users).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "failed to get all users")
	}

	return users, nil
}

func UpdateUser(db *gorm.DB, id string, user models.User) (models.User, error) {
	var existingUser models.User

	if err := db.First(&existingUser, id).Error; err != nil {
		return models.User{}, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if err := db.Model(&existingUser).Updates(&user).Error; err != nil {
		return models.User{}, fiber.NewError(fiber.StatusBadRequest, "Failed to update user")
	}

	return existingUser, nil
}
