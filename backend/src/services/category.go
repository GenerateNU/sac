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

// Gets all users (including soft deleted users) for testing
func (c *CategoryService) CreateCategory(category models.Category) (*models.Category, error) {
	// Validate the data based on the schema:
	if err := utilities.ValidateData(category); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Failed to validate the data")
	}
	
	return transactions.CreateCategory(c.DB, category);
}
