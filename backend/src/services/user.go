package services

import (
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
}

type UserService struct {
	DB *gorm.DB
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

// temporary
func createUserFromRequestBody(userBody models.CreateUserRequestBody) (models.User, error) {
	// TL DAVID -- some validation still needs to be done but depends on design

	validate := validator.New()
	validate.RegisterValidation("neu_email", utilities.ValidateEmail)
	validate.RegisterValidation("password", utilities.ValidatePassword)
	if err := validate.Struct(userBody); err != nil {
		return models.User{}, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User
	user.NUID = userBody.NUID
	user.FirstName = userBody.FirstName
	user.LastName = userBody.LastName
	user.Email = userBody.Email
	// TODO: hash
	user.PasswordHash = userBody.Password
	user.College = models.College(userBody.College)
	user.Year = models.Year(userBody.Year)

	return user, nil
}

func (u *UserService) CreateUser(userBody models.CreateUserRequestBody) (*models.User, error) {
	user, err := createUserFromRequestBody(userBody)
	if err != nil {
		return nil, err
	}

	return transactions.CreateUser(u.DB, &user)
}
