package services

import (
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
	CreateUser(userBody models.CreateUserRequestBody) (*models.User, *errors.Error)
	GetUsers(limit string, page string) ([]models.User, *errors.Error)
	GetUser(id string) (*models.User, *errors.Error)
	UpdateUser(id string, userBody models.UpdateUserRequestBody) (*models.User, *errors.Error)
	DeleteUser(id string) *errors.Error
	GetUserTags(id string) ([]models.Tag, *errors.Error)
	CreateUserTags(id string, tagIDs models.CreateUserTagsBody) ([]models.Tag, *errors.Error)
}

type UserService struct {
	DB       *gorm.DB
	Validate *validator.Validate
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

	return transactions.UpdateUser(u.DB, *idAsUUID, *user)
}

func (u *UserService) DeleteUser(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return transactions.DeleteUser(u.DB, *idAsUUID)
}

func (u *UserService) GetUserTags(id string) ([]models.Tag, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return transactions.GetUserTags(u.DB, *idAsUUID)
}

func (u *UserService) CreateUserTags(id string, tagIDs models.CreateUserTagsBody) ([]models.Tag, *errors.Error) {
	// Validate the id:
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if err := u.Validate.Struct(tagIDs); err != nil {
		return nil, &errors.FailedToValidateUserTags
	}

	// Retrieve a list of valid tags from the ids:
	tags, err := transactions.GetTagsByIDs(u.DB, tagIDs.Tags)

	// Update the user to reflect the new tags:
	return transactions.CreateUserTags(u.DB, *idAsUUID, tags)
}
