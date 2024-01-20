package services

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TagServiceInterface interface {
	CreateTag(partialTag models.CreateTagRequestBody) (*models.Tag, error)
	GetTag(id string) (*models.Tag, error)
}

type TagService struct {
	DB *gorm.DB
}

func (t *TagService) CreateTag(partialTag models.CreateTagRequestBody) (*models.Tag, error) {
	tag := models.Tag{
		Name:       partialTag.Name,
		CategoryID: partialTag.CategoryID,
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
