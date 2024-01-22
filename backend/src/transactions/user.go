package transactions

import (
	"errors"
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

func UpdateUser(db *gorm.DB, id uint, user models.User) (*models.User, error) {
	var existingUser models.User

	err := db.First(&existingUser, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
		} else {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "database error")
		}
	}

	if err := db.Model(&existingUser).Updates(&user).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "database error")
	}

	return &existingUser, nil
}
