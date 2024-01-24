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
	GetCategories(limit string, page string) ([]models.Category, *errors.Error)
	GetCategory(id string) (*models.Category, *errors.Error)
	UpdateCategory(id string, params models.CategoryRequestBody) (*models.Category, *errors.Error)
	DeleteCategory(id string) *errors.Error
}

type CategoryService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (c *CategoryService) CreateCategory(categoryBody models.CategoryRequestBody) (*models.Category, *errors.Error) {
	if err := c.Validate.Struct(categoryBody); err != nil {
		return nil, &errors.FailedToValidateCategory
	}

	category, err := utilities.MapRequestToModel(categoryBody, &models.Category{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	category.Name = cases.Title(language.English).String(category.Name)

	return transactions.CreateCategory(c.DB, *category)
}

func (c *CategoryService) GetCategories(limit string, page string) ([]models.Category, *errors.Error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)

	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)

	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	offset := (*pageAsInt - 1) * *limitAsInt

	return transactions.GetCategories(c.DB, *limitAsInt, offset)
}

func (c *CategoryService) GetCategory(id string) (*models.Category, *errors.Error) {
	uintId, err := utilities.ValidateID(id)

	if err != nil {
		return nil, err
	}

	return transactions.GetCategory(c.DB, *uintId)
}

func (c *CategoryService) UpdateCategory(id string, categoryBody models.CategoryRequestBody) (*models.Category, *errors.Error) {
	idAsUint, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	if err := c.Validate.Struct(categoryBody); err != nil {
		return nil, &errors.FailedToValidateTag
	}

	category, err := utilities.MapRequestToModel(categoryBody, &models.Category{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	category.Name = cases.Title(language.English).String(category.Name)

	return transactions.UpdateCategory(c.DB, *idAsUint, *category)
}

func (c *CategoryService) DeleteCategory(id string) *errors.Error {
	idAsUInt, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return transactions.DeleteCategory(c.DB, *idAsUInt)
}
