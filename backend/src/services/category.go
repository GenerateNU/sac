package services

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type CategoryServiceInterface interface {
	CreateCategory(params models.CreateCategoryRequestBody) (*models.Category, error)
	GetCategories() (*[]models.Category, error)
	GetCategory(id string) (*models.Category, error)
	UpdateCategory(id string, params models.UpdateCategoryRequestBody) (*models.Category, error)
	DeleteCategory(id string) error
 }

type CategoryService struct {
	DB *gorm.DB
}

func (c *CategoryService) CreateCategory(params models.CreateCategoryRequestBody) (*models.Category, error) {
	if err := utilities.ValidateData(params); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "failed to validate the data")
	}

	category := models.Category{
		Name: params.Name,
	}

	category.Name = cases.Title(language.English).String(category.Name)

	return transactions.CreateCategory(c.DB, category)
}

func (c *CategoryService) GetCategories() (*[]models.Category, error) {
	return transactions.GetCategories(c.DB)
}

func (c *CategoryService) GetCategory(id string) (*models.Category, error) {
	uintId, err := utilities.ValidateID(id)

	if err != nil {
		return nil, err
	}

	return transactions.GetCategory(c.DB, *uintId)
}

func (c *CategoryService) UpdateCategory(id string, params models.UpdateCategoryRequestBody) (*models.Category, error) {
	uintId, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if err := utilities.ValidateData(params); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "failed to validate the data")
	}

	category := models.Category{
		Name: params.Name,
	}

	category.Name = cases.Title(language.English).String(category.Name)

	return transactions.UpdateCategory(c.DB, *uintId, category)
}

func (c *CategoryService) DeleteCategory(id string) error {
	uintId, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return transactions.DeleteCategory(c.DB, *uintId)
}
