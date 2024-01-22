package services

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	DeleteUser(id string) error
	GetUser(id string) (*models.User, error)
	UpdateUser(id string, userBody models.UserRequestBody) (*models.User, error)
}

type UserService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

// Delete user with a specific id
func (u *UserService) DeleteUser(id string) error {
	idAsInt, err := utilities.ValidateID(id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return transactions.DeleteUser(u.DB, *idAsInt)
}

func (u *UserService) GetUser(userID string) (*models.User, error) {
	idAsUint, err := utilities.ValidateID(userID)

	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	return transactions.GetUser(u.DB, *idAsUint)
}

// Updates a user
func (u *UserService) UpdateUser(id string, userBody models.UserRequestBody) (*models.User, error) {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	if err := u.Validate.Struct(userBody); err != nil {
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

	return transactions.UpdateUser(u.DB, *idAsUint, *user)
}
