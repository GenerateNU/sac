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
	CreateTag(categoryID string, tagBody models.TagRequestBody) (*models.Tag, *errors.Error)
	GetTag(categoryID string, tagID string) (*models.Tag, *errors.Error)
	UpdateTag(categoryID string, tagID string, tagBody models.TagRequestBody) (*models.Tag, *errors.Error)
	DeleteTag(categoryID string, tagID string) *errors.Error
}

type TagService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewTagService(db *gorm.DB, validate *validator.Validate) *TagService {
	return &TagService{DB: db, Validate: validate}
}

func (t *TagService) CreateTag(categoryID string, tagBody models.TagRequestBody) (*models.Tag, *errors.Error) {
	categoryIDAsUUID, idErr := utilities.ValidateID(categoryID)

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

	tag.CategoryID = *categoryIDAsUUID

	return transactions.CreateTag(t.DB, *tag)
}

func (t *TagService) GetTag(categoryID string, tagID string) (*models.Tag, *errors.Error) {
	categoryIDAsUUID, idErr := utilities.ValidateID(categoryID)

	if idErr != nil {
		return nil, idErr
	}

	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

	if idErr != nil {
		return nil, idErr
	}

	}

	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

	if idErr != nil {
		return nil, idErr
	}

	return transactions.GetTag(t.DB, *categoryIDAsUUID, *tagIDAsUUID)
}

func (t *TagService) UpdateTag(categoryID string, tagID string, tagBody models.TagRequestBody) (*models.Tag, *errors.Error) {
	categoryIDAsUUID, idErr := utilities.ValidateID(categoryID)

	if idErr != nil {
		return nil, idErr
	}

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

	tag.CategoryID = *categoryIDAsUUID

	return transactions.UpdateTag(t.DB, *tagIDAsUUID, *tag)
}

func (t *TagService) DeleteTag(categoryID string, tagID string) *errors.Error {
	categoryIDAsUUID, idErr := utilities.ValidateID(categoryID)

	if idErr != nil {
		return idErr
	}

	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

	if idErr != nil {
		return idErr
	}

	return transactions.DeleteTag(t.DB, *categoryIDAsUUID, *tagIDAsUUID)
}
