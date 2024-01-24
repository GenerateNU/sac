package services

import (
	"strings"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, *errors.Error)
	CreateUser(userBody models.CreateUserRequestBody) (*models.User, *errors.Error)
	GetUser(id string) (*models.User, *errors.Error)
	UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, *errors.Error)
	DeleteUser(id string) *errors.Error
}

type UserService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, *errors.Error) {
	return transactions.GetAllUsers(u.DB)
}

func (u *UserService) CreateUser(userBody models.CreateUserRequestBody) (*models.User, *errors.Error) {
	if err := u.Validate.Struct(userBody); err != nil {
		return nil, &errors.FailedToValidateUser
	}

	user, err := utilities.MapRequestToModel(userBody, &models.User{})
	if err != nil {
		return nil, &errors.FailedToMapResposeToModel
	}

	passwordHash, err := auth.ComputePasswordHash(userBody.Password)
	if err != nil {
		return nil, &errors.FailedToComputePasswordHash
	}

	user.Email = strings.ToLower(userBody.Email)
	user.PasswordHash = *passwordHash

	return transactions.CreateUser(u.DB, user)
}

func (u *UserService) GetUser(id string) (*models.User, *errors.Error) {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetUser(u.DB, *idAsUint)
}

func (u *UserService) UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, *errors.Error) {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	if err := u.Validate.Struct(userBody); err != nil {
		return nil, &errors.FailedToValidateUser
	}

	passwordHash, err := auth.ComputePasswordHash(userBody.Password)
	if err != nil {
		return nil, &errors.FailedToComputePasswordHash
	}

	user, err := utilities.MapRequestToModel(userBody, &models.User{})
	if err != nil {
		return nil, &errors.FailedToMapResposeToModel
	}

	user.PasswordHash = *passwordHash

	return transactions.UpdateUser(u.DB, *idAsUint, *user)
}

// Delete user with a specific id
func (u *UserService) DeleteUser(id string) *errors.Error {
	idAsInt, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteUser(u.DB, *idAsInt)
}
