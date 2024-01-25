package services

import (
	"fmt"
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
	GetUsers(limit string, page string) ([]models.User, *errors.Error)
	GetUser(id string) (*models.User, *errors.Error)
	CreateUser(userBody models.CreateUserRequestBody) (*models.User, *errors.Error)
	UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, *errors.Error)
	DeleteUser(id string) *errors.Error
	Login(userBody models.LoginUserResponseBody) (*models.User, *errors.Error)
}

type UserService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserService(db *gorm.DB, validate *validator.Validate) *UserService {
	return &UserService{DB: db, Validate: validate}
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

func (u *UserService) Login(userBody models.LoginUserResponseBody) (*models.User, *errors.Error) {
	if err := u.Validate.Struct(userBody); err != nil {
		return nil, &errors.FailedToValidateUser
	}

	fmt.Println("hit 1")
	
	// check if user exists
	user, err := transactions.GetUserByEmail(u.DB, userBody.Email)
	if err != nil {
		return nil, &errors.UserNotFound
	}

	fmt.Println("hit 2")

	correct, passwordErr := auth.ComparePasswordAndHash(userBody.Password, user.PasswordHash)
	if passwordErr != nil {
		return nil, &errors.FailedToValidateUser
	}

	fmt.Println("hit 3")

	if !correct {
		return nil, &errors.FailedToValidateUser
	}

	fmt.Println("hit 4")

	return user, nil
}
