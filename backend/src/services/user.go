package services

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	CreateUserFromParams(params types.UserParams) (models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(id string, params types.UserParams) (models.User, error)
}

type UserService struct {
	DB *gorm.DB
}

// Creates a models.User from params. This *does not* interact with the database at all; the value will need to be
// passed to gorm.Db.Create(interface{}) for it to be persisted.
func (u *UserService) CreateUserFromParams(params types.UserParams) (models.User, error) {
	validate := validator.New()
	validate.RegisterValidation("neu_email", utilities.ValidateEmail)
	validate.RegisterValidation("password", utilities.ValidatePassword)
	if err := validate.Struct(params); err != nil {
		return models.User{}, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User
	user.NUID = params.NUID
	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Email = params.Email
	// TODO: hash
	user.PasswordHash = params.Password
	user.College = models.College(params.College)
	user.Year = models.Year(params.Year)

	return user, nil
}

// Gets all users (including soft deleted users) for testing
func (u *UserService) GetAllUsers() ([]models.User, error) {
	return transactions.GetAllUsers(u.DB)
}

// Updates a user
func (u *UserService) UpdateUser(id string, params types.UserParams) (models.User, error) {
	user, err := u.CreateUserFromParams(params)
	if err != nil {
		return models.User{}, err
	}

	return transactions.UpdateUser(u.DB, id, user)
}
