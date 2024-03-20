package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserTagServiceInterface interface {
	GetUserTags(id string) ([]models.Tag, *errors.Error)
	CreateUserTags(id string, tagIDs models.CreateUserTagsBody) ([]models.Tag, *errors.Error)
	DeleteUserTag(id string, tagID string) *errors.Error
}

type UserTagService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserTagService(db *gorm.DB, validate *validator.Validate) *UserTagService {
	return &UserTagService{DB: db, Validate: validate}
}

func (u *UserTagService) GetUserTags(id string) ([]models.Tag, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return transactions.GetUserTags(u.DB, *idAsUUID)
}

func (u *UserTagService) CreateUserTags(id string, tagIDs models.CreateUserTagsBody) ([]models.Tag, *errors.Error) {
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
	if err != nil {
		return nil, err
	}

	// Update the user to reflect the new tags:
	return transactions.CreateUserTags(u.DB, *idAsUUID, tags)
}

func (u *UserTagService) DeleteUserTag(id string, tagID string) *errors.Error {
	// Validate the userID:
	userIDAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	tagIDAsUUID, err := utilities.ValidateID(tagID)
	if err != nil {
		return err
	}

	return transactions.DeleteUserTag(u.DB, *userIDAsUUID, *tagIDAsUUID)
}
