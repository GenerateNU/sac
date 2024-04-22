package services

import (
	"fmt"
	"strings"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/models"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type UserServiceInterface interface {
	GetUsers(limit string, page string) ([]models.User, *errors.Error)
	GetMe(id string) (*models.User, *errors.Error)
	GetUser(id string) (*models.User, *errors.Error)
	CreateUser(userBody models.CreateUserRequestBody) (*models.User, *auth.Token, *errors.Error)
	UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, *errors.Error)
	UpdatePassword(id string, passwordBody models.UpdatePasswordRequestBody) *errors.Error
	DeleteUser(id string) *errors.Error
}

type UserService struct {
	types.ServiceParams
}

func NewUserService(serviceParams types.ServiceParams) UserServiceInterface {
	return &UserService{serviceParams}
}

func (u *UserService) CreateUser(userBody models.CreateUserRequestBody) (*models.User, *auth.Token, *errors.Error) {
	if err := u.Validate.Struct(userBody); err != nil {
		return nil, nil, &errors.FailedToValidateUser
	}

	if pwordValErr := auth.ValidatePassword(userBody.Password); pwordValErr != nil {
		return nil, nil, pwordValErr
	}

	user, err := utilities.MapRequestToModel(userBody, &models.User{})
	if err != nil {
		return nil, nil, &errors.FailedToMapRequestToModel
	}

	passwordHash, err := auth.ComputeHash(userBody.Password)
	if err != nil {
		return nil, nil, &errors.FailedToComputePasswordHash
	}

	user.Email = strings.ToLower(userBody.Email)
	user.PasswordHash = *passwordHash

	emailErr := u.Deps.Email.SendWelcomeEmail(fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email)
	if emailErr != nil {
		return nil, nil, &errors.FailedToSendEmail
	}

	user, userErr := transactions.CreateUser(u.DB, user)
	if userErr != nil {
		return nil, nil, userErr
	}

	authService := NewAuthService(u.ServiceParams)
	authErr := authService.SendCode(user.Email)
	if authErr != nil {
		return nil, nil, authErr
	}

	_, tokens, authErr := authService.Login(models.LoginUserResponseBody{Email: user.Email, Password: userBody.Password})
	if authErr != nil {
		return nil, nil, authErr
	}

	return user, tokens, nil
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

func (u *UserService) GetMe(id string) (*models.User, *errors.Error) {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	user, err := transactions.GetUser(u.DB, *idAsUUID)
	if err != nil {
		return nil, err
	}

	return user, nil
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

func (u *UserService) UpdatePassword(id string, passwordBody models.UpdatePasswordRequestBody) *errors.Error {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return idErr
	}

	if err := u.Validate.Struct(passwordBody); err != nil {
		return &errors.FailedToValidateUpdatePasswordBody
	}

	if pwordValErr := auth.ValidatePassword(passwordBody.OldPassword); pwordValErr != nil {
		return pwordValErr
	}

	if pwordValErr := auth.ValidatePassword(passwordBody.NewPassword); pwordValErr != nil {
		return pwordValErr
	}

	passwordHash, err := transactions.GetUserPasswordHash(u.DB, *idAsUUID)
	if err != nil {
		return err
	}

	correct, passwordErr := auth.CompareHash(passwordBody.OldPassword, *passwordHash)
	if passwordErr != nil || !correct {
		return &errors.FailedToValidateUser
	}

	hash, hashErr := auth.ComputeHash(passwordBody.NewPassword)
	if hashErr != nil {
		return &errors.FailedToValidateUser
	}

	return transactions.UpdatePassword(u.DB, *idAsUUID, *hash)
}

func (u *UserService) DeleteUser(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return transactions.DeleteUser(u.DB, *idAsUUID)
}
