package services

import (
	"strings"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type UserServiceInterface interface {
	CreateUser(userBody models.CreateUserRequestBody) (*models.User, *errors.Error)
	GetUsers(limit string, page string) ([]models.User, *errors.Error)
	GetUser(id string) (*models.User, *errors.Error)
	UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, *errors.Error)
	DeleteUser(id string) *errors.Error
}

type UserService struct {
	types.ServiceParams
}

func NewUserService(serviceParams types.ServiceParams) *UserService {
	return &UserService{serviceParams}
}

func (u *UserService) CreateUser(userBody models.CreateUserRequestBody) (*models.User, *errors.Error) {
	if err := u.Validate.Struct(userBody); err != nil {
		return nil, &errors.FailedToValidateUser
	}

	user, err := utilities.MapRequestToModel(userBody, &models.User{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	passwordHash, err := auth.ComputeHash(userBody.Password)
	if err != nil {
		return nil, &errors.FailedToComputePasswordHash
	}

	user.Email = strings.ToLower(userBody.Email)
	user.PasswordHash = *passwordHash

	// send email creation event to email service
	// email.SendWelcomeEmail(user.Name, user.Email)

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

	return transactions.GetUsers(u.DB, *limitAsInt, *pageAsInt)
}

func (u *UserService) GetUser(id string) (*models.User, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetUser(u.DB, *idAsUUID)
}

func (u *UserService) UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, *errors.Error) {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	if utilities.AtLeastOne(userBody, models.UpdateUserRequestBody{}) {
		return nil, &errors.FailedToValidateUser
	}

	if err := u.Validate.Struct(userBody); err != nil {
		return nil, &errors.FailedToValidateUser
	}

	user, err := utilities.MapRequestToModel(userBody, &models.User{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.UpdateUser(u.DB, *idAsUUID, *user)
}

func (u *UserService) DeleteUser(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return transactions.DeleteUser(u.DB, *idAsUUID)
}
