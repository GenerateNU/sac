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
	DeleteUser(id string) (error)
}

type UserService struct {
	DB *gorm.DB
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

// Delete user with a specific id
func (u *UserService) DeleteUser(id string) (error) {
	idAsInt, err := strconv.Atoi(id)
	if idAsInt < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "failed to validate id")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to validate id")
	}
	return transactions.DeleteUser(u.DB, uint(idAsInt))
}
