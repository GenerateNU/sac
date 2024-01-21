package services

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TagServiceInterface interface {
	CreateTag(tagBody models.CreateTagRequestBody) (*models.Tag, error)
	GetTag(id string) (*models.Tag, error)
	UpdateTag(id string, tagBody models.UpdateTagRequestBody) error
	DeleteTag(id string) error
}

type TagService struct {
	DB *gorm.DB
}

func (t *TagService) CreateTag(tagBody models.CreateTagRequestBody) (*models.Tag, error) {
	tag := models.Tag{
		Name:       tagBody.Name,
		CategoryID: tagBody.CategoryID,
	}

	if err := utilities.ValidateData(tag); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "failed to validate the data")
	}

	return transactions.CreateTag(t.DB, tag)
}

func (t *TagService) GetTag(id string) (*models.Tag, error) {
	idAsUint, err := utilities.ValidateID(id)

	if err != nil {
		return nil, err
	}

	return transactions.GetTag(t.DB, *idAsUint)
}

func (t *TagService) UpdateTag(id string, tagBody models.UpdateTagRequestBody) error {
	idAsUint, err := utilities.ValidateID(id)

	if err != nil {
		return err
	}

	tag := models.Tag{
		Name:       tagBody.Name,
		CategoryID: tagBody.CategoryID,
	}

	if err := utilities.ValidateData(tag); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to validate the data")
	}

	return transactions.UpdateTag(t.DB, *idAsUint, tag)
}

func (t *TagService) DeleteTag(id string) error {
	idAsUint, err := utilities.ValidateID(id)

	if err != nil {
		return err
	}

	return transactions.DeleteTag(t.DB, *idAsUint)
}
