package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/auth"
	authEntities "github.com/GenerateNU/sac/backend/entities/auth"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type UserServiceInterface interface {
	GetUsers(pageInfo fiberpaginate.PageInfo) ([]models.User, error)
	GetMe(id uuid.UUID) (*models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id string, userBody UpdateUserRequestBody) (*models.User, error)
	UpdatePassword(id string, passwordBody authEntities.UpdatePasswordRequestBody) error
	DeleteUser(id string) error
}

type UserService struct {
	types.ServiceParams
}

func NewUserService(serviceParams types.ServiceParams) UserServiceInterface {
	return &UserService{serviceParams}
}

func (u *UserService) GetUsers(pageInfo fiberpaginate.PageInfo) ([]models.User, error) {
	return GetUsers(u.DB, pageInfo)
}

func (u *UserService) GetMe(id uuid.UUID) (*models.User, error) {
	return users.GetUser(u.DB, id)
}

func (u *UserService) GetUser(id string) (*models.User, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return users.GetUser(u.DB, *idAsUUID)
}

func (u *UserService) UpdateUser(id string, userBody UpdateUserRequestBody) (*models.User, error) {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	if utilities.AtLeastOne(userBody, UpdateUserRequestBody{}) {
		return nil, errors.New("no fields to update")
	}

	if err := utilities.Validate(u.Validate, userBody); err != nil {
		return nil, err
	}

	user, err := utilities.MapJsonTags(userBody, &models.User{})
	if err != nil {
		return nil, err
	}

	return UpdateUser(u.DB, *idAsUUID, *user)
}

func (u *UserService) UpdatePassword(id string, passwordBody authEntities.UpdatePasswordRequestBody) error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	if err := utilities.Validate(u.Validate, passwordBody,
		*utilities.NewMaybeError("old_password", auth.ValidatePassword(passwordBody.OldPassword)),
		*utilities.NewMaybeError("new_password", auth.ValidatePassword(passwordBody.NewPassword))); err != nil {
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
