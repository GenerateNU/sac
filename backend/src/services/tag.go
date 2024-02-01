package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TagServiceInterface interface {
	CreateTag(tagBody models.TagRequestBody) (*models.Tag, *errors.Error)
	GetTag(id string) (*models.Tag, *errors.Error)
	UpdateTag(id string, tagBody models.TagRequestBody) (*models.Tag, *errors.Error)
	DeleteTag(id string) *errors.Error
	GetTagClubs(id string) ([]models.Club, *errors.Error)
}

type TagService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (t *TagService) CreateTag(tagBody models.TagRequestBody) (*models.Tag, *errors.Error) {
	if err := t.Validate.Struct(tagBody); err != nil {
		return nil, &errors.FailedToValidateTag
	}

	tag, err := utilities.MapRequestToModel(tagBody, &models.Tag{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.CreateTag(t.DB, *tag)
}

func (t *TagService) GetTag(id string) (*models.Tag, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return transactions.GetTag(t.DB, *idAsUUID)
}

func (t *TagService) UpdateTag(id string, tagBody models.TagRequestBody) (*models.Tag, *errors.Error) {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	if err := t.Validate.Struct(tagBody); err != nil {
		return nil, &errors.FailedToValidateTag
	}

	tag, err := utilities.MapRequestToModel(tagBody, &models.Tag{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.UpdateTag(t.DB, *idAsUUID, *tag)
}

func (t *TagService) DeleteTag(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteTag(t.DB, *idAsUUID)
}

func (t *TagService) GetTagClubs(id string) ([]models.Club, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetTagClubs(t.DB, *idAsUUID)
}