package services

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	UpdateUser(id string, userBody models.UserRequestBody) (*models.User, error)
}

type UserService struct {
	DB *gorm.DB
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

// Updates a user
func (u *UserService) UpdateUser(id string, userBody models.UserRequestBody) (*models.User, error) {
	if err := utilities.ValidateData(userBody); err != nil {
		return nil, fiber.ErrBadRequest
	}

	passwordHash, err := auth.ComputePasswordHash(userBody.Password)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	user, err := utilities.MapResponseToModel(userBody, &models.User{})
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	user.PasswordHash = *passwordHash

	return transactions.UpdateUser(u.DB, *user)
}
