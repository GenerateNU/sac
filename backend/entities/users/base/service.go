package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errs"
	"github.com/garrettladley/fiberpaginate"

	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type UserServiceInterface interface {
	GetUsers(pageInfo fiberpaginate.PageInfo) ([]models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id string, userBody UpdateUserRequestBody) (*models.User, error)
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
		return nil, errs.ErrAtLeastOne
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

func (u *UserService) DeleteUser(id string) error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return DeleteUser(u.DB, *idAsUUID)
}
