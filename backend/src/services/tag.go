package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TagServiceInterface interface {
	CreateTag(tagBody models.TagRequestBody) (*models.Tag, *errors.Error)
	GetTag(id string) (*models.Tag, *errors.Error)
	UpdateTag(id string, tagBody models.TagRequestBody) (*models.Tag, *errors.Error)
	DeleteTag(id string) *errors.Error
}

type TagService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (t *TagService) CreateTag(tagBody models.TagRequestBody) (*models.Tag, *errors.Error) {
	if err := t.Validate.Struct(tagBody); err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: "failed to validate tag"}
	}

	tag, err := utilities.MapResponseToModel(tagBody, &models.Tag{})
	if err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: "failed to create tag"}
	}

	return transactions.CreateTag(t.DB, *tag)
}

func (t *TagService) GetTag(id string) (*models.Tag, *errors.Error) {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: "failed to validate id"}
	}

	return transactions.GetTag(t.DB, *idAsUint)
}

func (t *TagService) UpdateTag(id string, tagBody models.TagRequestBody) (*models.Tag, *errors.Error) {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: "failed to validate id"}
	}

	if err := t.Validate.Struct(tagBody); err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: "failed to validate tag"}
	}

	tag, err := utilities.MapResponseToModel(tagBody, &models.Tag{})
	if err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: "failed to update tag"}
	}

	return transactions.UpdateTag(t.DB, *idAsUint, *tag)
}

func (t *TagService) DeleteTag(id string) *errors.Error {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.Error{StatusCode: fiber.StatusBadRequest, Message: "failed to validate id"}
	}

	return transactions.DeleteTag(t.DB, *idAsUint)
}
