package services

import (
	"strings"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	CreateUser(userBody models.CreateUserRequestBody) (*models.User, *errors.Error)
	GetUsers(limit string, page string) ([]models.User, *errors.Error)
	GetUser(id string) (*models.User, *errors.Error)
	UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, *errors.Error)
	DeleteUser(id string) *errors.Error
	GetAllUsers() ([]models.User, error)
	GetUser(id uint) (*models.User, error)
	Register(userBody models.CreateUserResponseBody) (*models.User, error)
	Login(userBody models.LoginUserResponseBody) (*models.User, error)
}

type UserService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (u *UserService) CreateUser(userBody models.CreateUserRequestBody) (*models.User, *errors.Error) {
	if err := u.Validate.Struct(userBody); err != nil {
		return nil, &errors.FailedToValidateUser
	}

	user, err := utilities.MapRequestToModel(userBody, &models.User{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	passwordHash, err := auth.ComputePasswordHash(userBody.Password)
	if err != nil {
		return nil, &errors.FailedToComputePasswordHash
	}

	user.Email = strings.ToLower(userBody.Email)
	user.PasswordHash = *passwordHash

	return transactions.CreateUser(u.DB, user)
}

func (u *UserService) GetUsers(limit string, page string) ([]models.User, *errors.Error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)

	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)

	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	offset := (*pageAsInt - 1) * *limitAsInt

	return transactions.GetUsers(u.DB, *limitAsInt, offset)
}

func (u *UserService) GetUser(id string) (*models.User, *errors.Error) {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetUser(u.DB, *idAsUint)
}

func (u *UserService) UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, *errors.Error) {
	idAsUint, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
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
		return nil, &errors.FailedToMapRequestToModel
	}

	user.PasswordHash = *passwordHash

	return transactions.UpdateUser(u.DB, *idAsUint, *user)
}

func (u *UserService) DeleteUser(id string) *errors.Error {
	idAsInt, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return transactions.DeleteUser(u.DB, *idAsInt)
}

func (u *UserService) GetUser(id uint) (*models.User, error) {
	return transactions.GetUser(u.DB, id)
}

// Registers a user
func (u *UserService) Register(userBody models.CreateUserResponseBody) (*models.User, error) {
	if err := utilities.ValidateData(userBody); err != nil {
		return nil, err
	}

	passwordHash, err := auth.ComputePasswordHash(userBody.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Role:         models.Student,
		NUID:         userBody.NUID,
		FirstName:    userBody.FirstName,
		LastName:     userBody.LastName,
		Email:        userBody.Email,
		PasswordHash: *passwordHash,
		College:      models.College(userBody.College),
		Year:         models.Year(userBody.Year),
	}

	return transactions.CreateUser(u.DB, user)
}

func (u *UserService) Login(userBody models.LoginUserResponseBody) (*models.User, error) {
	if err := utilities.ValidateData(userBody); err != nil {
		return nil, err
	}

	// check if user exists
	user, err := transactions.GetUserByEmail(u.DB, userBody.Email)
	if err != nil {
		return nil, err
	}

	correct, err := auth.ComparePasswordAndHash(userBody.Password, user.PasswordHash)

	if err != nil {
		return nil, err
	}

	if !correct {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "incorrect password")
	}

	return user, nil
}
