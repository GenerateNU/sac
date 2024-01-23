package services

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TagServiceInterface interface {
	CreateTag(tagBody models.TagRequestBody) (*models.Tag, error)
	GetTag(id string) (*models.Tag, error)
	UpdateTag(id string, tagBody models.TagRequestBody) (*models.Tag, error)
	DeleteTag(id string) error
}

type TagService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (t *TagService) CreateTag(tagBody models.TagRequestBody) (*models.Tag, error) {
	if err := t.Validate.Struct(tagBody); err != nil {
		return nil, fiber.ErrBadRequest
	}

	tag, err := utilities.MapResponseToModel(tagBody, &models.Tag{})
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return transactions.CreateTag(t.DB, *tag)
}

func (t *TagService) GetTag(id string) (*models.Tag, error) {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	return transactions.GetTag(t.DB, *idAsUint)
}

func (t *TagService) UpdateTag(id string, tagBody models.TagRequestBody) (*models.Tag, error) {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	if err := t.Validate.Struct(tagBody); err != nil {
		return nil, fiber.ErrBadRequest
	}

	tag, err := utilities.MapResponseToModel(tagBody, &models.Tag{})
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return transactions.UpdateTag(t.DB, *idAsUint, *tag)
}

func (t *TagService) DeleteTag(id string) error {
	idAsUint, err := utilities.ValidateID(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	return transactions.DeleteTag(t.DB, *idAsUint)
}
