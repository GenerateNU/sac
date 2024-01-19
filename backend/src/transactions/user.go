package transactions

import (
	"backend/src/models"

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
