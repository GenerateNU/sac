package base

import (
	"errors"
	"fmt"
	"strings"

	"github.com/GenerateNU/sac/backend/auth"
	authBase "github.com/GenerateNU/sac/backend/entities/auth/base"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/middleware"

	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type UserServiceInterface interface {
	GetUsers(pagination middleware.Pagination) ([]models.User, error)
	GetMe(id string) (*models.User, error)
	GetUser(id string) (*models.User, error)
	CreateUser(userBody models.CreateUserRequestBody) (*models.User, *auth.Token, error)
	UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, error)
	UpdatePassword(id string, passwordBody models.UpdatePasswordRequestBody) error
	DeleteUser(id string) error
}

type UserService struct {
	types.ServiceParams
	AuthService authBase.AuthServiceInterface
}

func NewUserService(serviceParams types.ServiceParams, authService authBase.AuthServiceInterface) UserServiceInterface {
	return &UserService{
		serviceParams,
		authService,
	}
}

func (u *UserService) CreateUser(userBody models.CreateUserRequestBody) (*models.User, *auth.Token, error) {
	if err := utilities.Validate(u.Validate, userBody); err != nil {
		return nil, nil, err
	}

	if err := auth.ValidatePassword(userBody.Password); err != nil {
		return nil, nil, err
	}

	user, err := utilities.MapRequestToModel(userBody, &models.User{})
	if err != nil {
		return nil, nil, err
	}

	passwordHash, err := auth.ComputeHash(userBody.Password)
	if err != nil {
		return nil, nil, err
	}

	user.Email = strings.ToLower(userBody.Email)
	user.PasswordHash = *passwordHash

	if err := u.Integrations.Email.SendWelcomeEmail(fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email); err != nil {
		return nil, nil, fmt.Errorf("failed to send welcome email: %w", err)
	}

	user, err = CreateUser(u.DB, user)
	if err != nil {
		return nil, nil, err
	}

	if err := u.AuthService.SendCode(user.Email); err != nil {
		return nil, nil, err
	}

	_, tokens, err := u.AuthService.Login(models.LoginUserResponseBody{Email: user.Email, Password: userBody.Password})
	if err != nil {
		return nil, nil, err
	}

	return user, tokens, nil
}

func (u *UserService) GetUsers(pagination middleware.Pagination) ([]models.User, error) {
	return GetUsers(u.DB, pagination)
}

func (u *UserService) GetMe(id string) (*models.User, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	user, err := users.GetUser(u.DB, *idAsUUID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetUser(id string) (*models.User, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return users.GetUser(u.DB, *idAsUUID)
}

func (u *UserService) UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, error) {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	if utilities.AtLeastOne(userBody, models.UpdateUserRequestBody{}) {
		return nil, errors.New("no fields to update")
	}

	if err := utilities.Validate(u.Validate, userBody); err != nil {
		return nil, err
	}

	user, err := utilities.MapRequestToModel(userBody, &models.User{})
	if err != nil {
		return nil, err
	}

	return UpdateUser(u.DB, *idAsUUID, *user)
}

func (u *UserService) UpdatePassword(id string, passwordBody models.UpdatePasswordRequestBody) error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	if err := utilities.Validate(u.Validate, passwordBody); err != nil {
		return err
	}

	if err := auth.ValidatePassword(passwordBody.OldPassword); err != nil {
		return err
	}

	if err := auth.ValidatePassword(passwordBody.NewPassword); err != nil {
		return err
	}

	passwordHash, err := GetUserPasswordHash(u.DB, *idAsUUID)
	if err != nil {
		return err
	}

	correct, passwordErr := auth.CompareHash(passwordBody.OldPassword, *passwordHash)
	if passwordErr != nil || !correct {
		return err
	}

	hash, hashErr := auth.ComputeHash(passwordBody.NewPassword)
	if hashErr != nil {
		return err
	}

	return users.UpdatePassword(u.DB, *idAsUUID, *hash)
}

func (u *UserService) DeleteUser(id string) error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return DeleteUser(u.DB, *idAsUUID)
}
