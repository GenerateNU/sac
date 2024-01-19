package services

import (
	"backend/src/models"
	"backend/src/transactions"
	"backend/src/utilities"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

	category.Name = cases.Title(language.English).String(category.Name)

	return transactions.CreateCategory(c.DB, category)
}
