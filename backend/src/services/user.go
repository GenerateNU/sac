package services

import (
	"errors"

	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

type UserServiceInterface interface {
	ValidateUserParams(params types.UserParams, noEmptyFields bool) error
	CreateUserFromParams(params types.UserParams) models.User
	GetAllUsers() ([]models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id string, params types.UserParams) (models.User, error)
}

type UserService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

// Validates the fields of params, and returns an error if any field is invalid. If noEmptyFields is true,
// then an error will be thrown if any field is missing.
func (u *UserService) ValidateUserParams(params types.UserParams, noEmptyFields bool) error {
	// check for empty fields
	if noEmptyFields {
		if params.NUID == "" {
			return errors.New("nuid is missing")
		}
		if params.FirstName == "" {
			return errors.New("first name is missing")
		}
		if params.LastName == "" {
			return errors.New("last name is missing")
		}
		if params.Email == "" {
			return errors.New("email is missing")
		}
		if params.Password == "" {
			return errors.New("password is missing")
		}
		if params.College == "" {
			return errors.New("college is missing")
		}
		if params.Year == 0 {
			return errors.New("year is missing")
		}
	}

	// run validation rules
	if err := utilities.ValidateNUID(params.NUID); params.NUID != "" && err != nil {
		return err
	}
	if err := utilities.ValidateEmail(params.Email); params.Email != "" && err != nil {
		return err
	}
	if err := utilities.ValidatePassword(params.Password); params.Password != "" && err != nil {
		return err
	}
	if err := utilities.ValidateCollege(params.College); params.College != "" && err != nil {
		return err
	}
	if err := utilities.ValidateYear(params.Year); params.Year != 0 && err != nil {
		return err
	}

	return nil
}

// Creates a models.User from params. This *does not* interact with the database at all; the value will need to be
// passed to gorm.Db.Create(interface{}) for it to be persisted.
func (u *UserService) CreateUserFromParams(params types.UserParams) models.User {
	var user models.User
	user.NUID = params.NUID
	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Email = params.Email
	// TODO: hash
	user.PasswordHash = params.Password
	user.College = models.College(params.College)
	user.Year = models.Year(params.Year)

	return user
}

// Gets all users (including soft deleted users) for testing
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
func (u *UserService) UpdateUser(id string, params types.UserParams) (models.User, error) {
	if err := u.ValidateUserParams(params, false); err != nil {
		return models.User{}, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return transactions.UpdateUser(u.DB, id, u.CreateUserFromParams(params))
}
