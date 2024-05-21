package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/tags"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type TagServiceInterface interface {
	GetTags() ([]models.Tag, error)
	CreateTag(tagBody CreateTagRequestBody) (*models.Tag, error)
	GetTag(id string) (*models.Tag, error)
	UpdateTag(id string, tagBody UpdateTagRequestBody) (*models.Tag, error)
	DeleteTag(id string) error
}

type TagService struct {
	types.ServiceParams
}

func NewTagService(serviceParams types.ServiceParams) TagServiceInterface {
	return &TagService{serviceParams}
}

func (t *TagService) CreateTag(tagBody CreateTagRequestBody) (*models.Tag, error) {
	if err := utilities.Validate(t.Validate, tagBody); err != nil {
		return nil, err
	}

	tag, err := utilities.MapJsonTags(tagBody, &models.Tag{})
	if err != nil {
		return nil, err
	}

	return CreateTag(t.DB, *tag)
}

func (t *TagService) GetTags() ([]models.Tag, error) {
	return GetTags(t.DB)
}

func (t *TagService) GetTag(tagID string) (*models.Tag, error) {
	tagIDAsUUID, err := utilities.ValidateID(tagID)
	if err != nil {
		return nil, err
	}

	return tags.GetTag(t.DB, *tagIDAsUUID)
}

func (t *TagService) UpdateTag(tagID string, tagBody UpdateTagRequestBody) (*models.Tag, error) {
	tagIDAsUUID, err := utilities.ValidateID(tagID)
	if err != nil {
		return nil, err
	}

	if utilities.AtLeastOne(tagBody, UpdateTagRequestBody{}) {
		return nil, errors.New("at least one field must be present")
	}

	if err := utilities.Validate(t.Validate, tagBody); err != nil {
		return nil, err
	}

	tag, err := utilities.MapJsonTags(tagBody, &models.Tag{})
	if err != nil {
		return nil, err
	}

	return UpdateTag(t.DB, *tagIDAsUUID, *tag)
}

func (t *TagService) DeleteTag(tagID string) error {
	tagIDAsUUID, err := utilities.ValidateID(tagID)
	if err != nil {
		return err
	}

	return DeleteTag(t.DB, *tagIDAsUUID)
}
