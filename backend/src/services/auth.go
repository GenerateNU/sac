package services

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	GetRole(id string) (*models.UserRole, *errors.Error)
	Me(id string) (*models.User, *errors.Error)
	Login(userBody models.LoginUserResponseBody) (*models.User, *errors.Error)
}

type AuthService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewAuthService(db *gorm.DB, validate *validator.Validate) *AuthService {
	return &AuthService{
		DB:       db,
		Validate: validate,
	}
}

func (a *AuthService) Me(id string) (*models.User, *errors.Error) {
	idAsUint, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}
	
	user, err := transactions.GetUser(a.DB, *idAsUint)
	if err != nil {
		return nil, &errors.UserNotFound
	}

	return user, nil
}

func (a *AuthService) Login(userBody models.LoginUserResponseBody) (*models.User, *errors.Error) {
	if err := a.Validate.Struct(userBody); err != nil {
		return nil, &errors.FailedToValidateUser
	}

	// check if user exists
	user, err := transactions.GetUserByEmail(a.DB, userBody.Email)
	if err != nil {
		return nil, &errors.UserNotFound
	}

	correct, passwordErr := auth.ComparePasswordAndHash(userBody.Password, user.PasswordHash)
	if passwordErr != nil {
		return nil, &errors.FailedToValidateUser
	}

	if !correct {
		return nil, &errors.FailedToValidateUser
	}

	return user, nil
}

func (a *AuthService) GetRole(id string) (*models.UserRole, *errors.Error) {
	idAsUint, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	user, err := transactions.GetUser(a.DB, *idAsUint)
	if err != nil {
		return nil, &errors.UserNotFound
	}

	role := models.UserRole(user.Role)

	return &role, nil
}