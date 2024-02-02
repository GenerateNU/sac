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
	GetTag(tagID string) (*models.Tag, *errors.Error)
	UpdateTag(tagID string, tagBody models.TagRequestBody) (*models.Tag, *errors.Error)
	DeleteTag(tagID string) *errors.Error
}

type TagService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewTagService(db *gorm.DB, validate *validator.Validate) *TagService {
	return &TagService{DB: db, Validate: validate}
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

func (t *TagService) GetTag(tagID string) (*models.Tag, *errors.Error) {
	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

	if idErr != nil {
		return nil, idErr
	}

	return transactions.GetTag(t.DB, *tagIDAsUUID)
}

func (t *TagService) UpdateTag(tagID string, tagBody models.TagRequestBody) (*models.Tag, *errors.Error) {
	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

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

	return transactions.UpdateTag(t.DB, *tagIDAsUUID, *tag)
}

func (t *TagService) DeleteTag(tagID string) *errors.Error {
	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

	if idErr != nil {
		return idErr
	}

	return transactions.DeleteTag(t.DB, *tagIDAsUUID)
}
