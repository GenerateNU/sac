package services

import (
	"backend/src/models"

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
	var users []models.User
	if err := u.DB.Unscoped().Omit("password").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil

}
