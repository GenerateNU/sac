package services

import (
	"backend/src/models"
	"backend/src/transactions"
	"backend/src/utilities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TagServiceInterface interface {
	CreateTag(tag models.Tag) (*models.Tag, error)
}

type TagService struct {
	DB *gorm.DB
}

func (t *TagService) CreateTag(tag models.Tag) (*models.Tag, error) {
	if err := utilities.ValidateData(tag); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "failed to validate the data")
	}

	return transactions.CreateTag(t.DB, tag)
}
