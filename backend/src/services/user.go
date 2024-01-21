package services

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUser(string) (*models.User, error)
}

type UserService struct {
	DB *gorm.DB
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

func (u *UserService) GetUser(userID string) (*models.User, error) {
	if integer, integerErr := strconv.Atoi(userID); integerErr != nil || integer <= 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "id must be a positive integer")
	}

	return transactions.GetUser(u.DB, userID)
}
