package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type TagServiceInterface interface {
	GetTags() ([]models.Tag, *errors.Error)
	CreateTag(tagBody models.CreateTagRequestBody) (*models.Tag, *errors.Error)
	GetTag(id string) (*models.Tag, *errors.Error)
	UpdateTag(id string, tagBody models.UpdateTagRequestBody) (*models.Tag, *errors.Error)
	DeleteTag(id string) *errors.Error
}

type TagService struct {
	types.ServiceParams
}

func NewTagService(serviceParams types.ServiceParams) TagServiceInterface {
	return &TagService{serviceParams}
}

func (t *TagService) CreateTag(tagBody models.CreateTagRequestBody) (*models.Tag, *errors.Error) {
	if err := t.Validate.Struct(tagBody); err != nil {
		return nil, &errors.FailedToValidateTag
	}

	tag, err := utilities.MapRequestToModel(tagBody, &models.Tag{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.CreateTag(t.DB, *tag)
}

func (t *TagService) GetTags() ([]models.Tag, *errors.Error) {
	return transactions.GetTags(t.DB)
}

func (t *TagService) GetTag(tagID string) (*models.Tag, *errors.Error) {
	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

	if idErr != nil {
		return nil, idErr
	}

	return transactions.GetTag(t.DB, *tagIDAsUUID)
}

func (t *TagService) UpdateTag(tagID string, tagBody models.UpdateTagRequestBody) (*models.Tag, *errors.Error) {
	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

	if idErr != nil {
		return nil, idErr
	}

	if utilities.AtLeastOne(tagBody, models.UpdateTagRequestBody{}) {
		return nil, &errors.FailedToValidateTag
	}

	if err := t.Validate.Struct(tagBody); err != nil {
		return nil, &errors.FailedToValidateTag
	}

	tag, err := utilities.MapRequestToModel(tagBody, &models.Tag{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.UpdateTag(t.DB, *tagIDAsUUID, *tag)
}

func (t *TagService) DeleteTag(tagID string) *errors.Error {
	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

	if idErr != nil {
		return idErr
	}

	return transactions.DeleteTag(t.DB, *tagIDAsUUID)
}
