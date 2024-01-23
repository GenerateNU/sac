package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"

	"github.com/go-playground/validator/v10"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"gorm.io/gorm"
)

type CategoryServiceInterface interface {
	CreateCategory(categoryBody models.CategoryRequestBody) (*models.Category, *errors.Error)
}

type CategoryService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (c *CategoryService) CreateCategory(categoryBody models.CategoryRequestBody) (*models.Category, *errors.Error) {
	if err := c.Validate.Struct(categoryBody); err != nil {
		return nil, &errors.FailedToValidateCategory
	}

	category, err := utilities.MapResponseToModel(categoryBody, &models.Category{})
	if err != nil {
		return nil, &errors.FailedToMapResposeToModel
	}

	category.Name = cases.Title(language.English).String(category.Name)

	return transactions.CreateCategory(c.DB, *category)
}
