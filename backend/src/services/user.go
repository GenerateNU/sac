package services

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"

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

func (u *UserService) GetUser(id string) (*models.User, error) {
	return transactions.GetUser(u.DB, id)
}
