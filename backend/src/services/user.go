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
	GetUser(id uint) (*models.User, error)
	Register(userBody models.CreateUserResponseBody) (*models.User, error)
	Login(userBody models.LoginUserResponseBody) (*models.User, error)
}

type UserService struct {
	DB *gorm.DB
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
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
