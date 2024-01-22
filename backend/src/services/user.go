package services

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	DeleteUser(id string) (error)
	GetUser(id string) (*models.User, error)
}

type UserService struct {
	DB *gorm.DB
}

func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

// Delete user with a specific id
func (u *UserService) DeleteUser(id string) (error) {
	idAsInt, err := utilities.ValidateID(id);
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return transactions.DeleteUser(u.DB, *idAsInt)
}

func (u *UserService) GetUser(userID string) (*models.User, error) {
	idAsUint, err := utilities.ValidateID(userID)

	if err != nil {
		return nil, err
	}

	return transactions.GetUser(u.DB, *idAsUint)
}
