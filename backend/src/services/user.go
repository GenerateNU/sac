package services

import (
	"backend/src/models"
	"backend/src/transactions"

	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	UpdateUser(id string, user models.User) (models.User, error)
}

type UserService struct {
	DB *gorm.DB
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

// Updates a user
func (u *UserService) UpdateUser(id string, payload models.User) (models.User, error) {

	//TODO: validation
	return transactions.UpdateUser(u.DB, id, payload)
}