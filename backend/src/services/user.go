package services

import (
	"backend/src/models"
	"backend/src/transactions"

	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
}

type UserService struct {
	DB *gorm.DB
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}
