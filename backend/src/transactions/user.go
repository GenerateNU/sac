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

func GetUser(db *gorm.DB, id uint) (*models.User, error) {
	var user models.User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return &user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return &user, nil
}

func CreateUser(db *gorm.DB, user models.User) (*models.User, error) {
	if err := db.Create(&user).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "failed to create user")
	}

	return &user, nil
}