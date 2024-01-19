package services

import (
	"backend/src/models"
	"backend/src/transactions"
	"backend/src/utilities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryServiceInterface interface {
	CreateCategory(category models.Category) (*models.Category, error)
}

type CategoryService struct {
	DB *gorm.DB
}

func (c *CategoryService) CreateCategory(category models.Category) (*models.Category, error) {
	if err := utilities.ValidateData(category); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "failed to validate the data")
	}

	return transactions.CreateCategory(c.DB, category)
}
