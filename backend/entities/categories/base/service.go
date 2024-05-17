package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CategoryServiceInterface interface {
	CreateCategory(categoryBody models.CategoryRequestBody) (*models.Category, error)
	GetCategories(pageInfo fiberpaginate.PageInfo) ([]models.Category, error)
	GetCategory(id string) (*models.Category, error)
	UpdateCategory(id string, params models.CategoryRequestBody) (*models.Category, error)
	DeleteCategory(id string) error
}

type CategoryService struct {
	types.ServiceParams
}

func NewCategoryService(params types.ServiceParams) CategoryServiceInterface {
	return &CategoryService{params}
}

func (c *CategoryService) CreateCategory(categoryBody models.CategoryRequestBody) (*models.Category, error) {
	if err := utilities.Validate(c.Validate, categoryBody); err != nil {
		return nil, err
	}

	category, err := utilities.MapJsonTags(categoryBody, &models.Category{})
	if err != nil {
		return nil, err
	}

	category.Name = cases.Title(language.English).String(category.Name)

	return CreateCategory(c.DB, *category)
}

func (c *CategoryService) GetCategories(pageInfo fiberpaginate.PageInfo) ([]models.Category, error) {
	return GetCategories(c.DB, pageInfo)
}

func (c *CategoryService) GetCategory(id string) (*models.Category, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return GetCategory(c.DB, *idAsUUID)
}

func (c *CategoryService) UpdateCategory(id string, categoryBody models.CategoryRequestBody) (*models.Category, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(c.Validate, categoryBody); err != nil {
		return nil, err
	}

	category, err := utilities.MapJsonTags(categoryBody, &models.Category{})
	if err != nil {
		return nil, err
	}

	category.Name = cases.Title(language.English).String(category.Name)

	return UpdateCategory(c.DB, *idAsUUID, *category)
}

func (c *CategoryService) DeleteCategory(id string) error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return DeleteCategory(c.DB, *idAsUUID)
}
