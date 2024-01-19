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
		return nil, fiber.ErrInternalServerError
	}

	return users, nil
}

func GetUser(db *gorm.DB, id uint) (*models.User, error) {
	var user models.User

	if err := db.Omit("password_hash").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		} else {
			return nil, fiber.ErrInternalServerError
		}
	}

	return &user, nil
}

func UpdateUser(db *gorm.DB, id uint, user models.User) (*models.User, error) {
	var existingUser models.User

	err := db.First(&existingUser, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		} else {
			return nil, fiber.ErrInternalServerError
		}
	}

	if err := db.Model(&existingUser).Updates(&user).Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return &existingUser, nil
}

func DeleteUser(db *gorm.DB, id string) error {
	var deletedUser models.User

	result := db.Where("id = ?", id).Delete(&deletedUser)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unable to delete user")
	}
	return nil
}
