package services

import (
	"strings"

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
	CreateUser(userBody models.CreateUserRequestBody) (*models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, error)
	DeleteUser(id string) error
}

type UserService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

func (u *UserService) CreateUser(userBody models.CreateUserRequestBody) (*models.User, error) {
	if err := u.Validate.Struct(userBody); err != nil {
		return nil, fiber.ErrBadRequest
	}

	user, err := utilities.MapResponseToModel(userBody, &models.User{})
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	passwordHash, err := auth.ComputePasswordHash(userBody.Password)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	user.Email = strings.ToLower(userBody.Email)
	user.PasswordHash = *passwordHash

	return transactions.CreateUser(u.DB, user)
}

func (u *UserService) GetUser(id string) (*models.User, error) {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	return transactions.GetUser(u.DB, *idAsUint)
}

func (u *UserService) UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, error) {
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

// Delete user with a specific id
func (u *UserService) DeleteUser(id string) error {
	idAsInt, err := utilities.ValidateID(id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return transactions.DeleteUser(u.DB, *idAsInt)
}
