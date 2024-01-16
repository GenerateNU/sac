package services

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id string, user models.User) (models.User, error)
}

type UserService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

func (u *UserService) GetUser(userID string) (*models.User, error) {
	idAsUint, err := utilities.ValidateID(userID)

	if err != nil {
		return nil, err
	}

	return transactions.GetUser(u.DB, *idAsUint)
}

// Updates a user
func (u *UserService) UpdateUser(id string, payload models.User) (models.User, error) {

	//TODO: validation
	return transactions.UpdateUser(u.DB, id, payload)
}
