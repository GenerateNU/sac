package services

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryServiceInterface interface {
	CreateCategory(categoryBody models.CategoryRequestBody) (*models.Category, error)
}

type CategoryService struct {
	DB *gorm.DB
}

func (c *CategoryService) CreateCategory(categoryBody models.CategoryRequestBody) (*models.Category, error) {
	if err := utilities.ValidateData(categoryBody); err != nil {
		return nil, fiber.ErrBadRequest
	}

	category, err := utilities.MapResponseToModel(categoryBody, &models.Category{})
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	category.Name = cases.Title(language.English).String(category.Name)

	return transactions.CreateCategory(c.DB, *category)
}
