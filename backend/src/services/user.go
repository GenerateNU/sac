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
	UpdateUser(id string, params models.UpdateUserRequestBody) (*models.User, error)
}

type UserService struct {
	DB *gorm.DB
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

// Updates a user
func (u *UserService) UpdateUser(id string, params models.UpdateUserRequestBody) (*models.User, error) {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	validate.RegisterValidation("neu_email", utilities.ValidateEmail)
	validate.RegisterValidation("password", utilities.ValidatePassword)
	if err := validate.Struct(params); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	passwordHash, err := auth.ComputePasswordHash(params.Password)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "failed to hash password")
	}

	user, err := utilities.MapResponseToModel(params, &models.User{})
	if err != nil {
		return nil, err
	}

	user.PasswordHash = *passwordHash

	return transactions.UpdateUser(u.DB, *idAsUint, *user)
}
