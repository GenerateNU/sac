package base

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	categoryService CategoryServiceInterface
}

func NewCategoryController(categoryService CategoryServiceInterface) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

// CreateCategory godoc
//
// @Summary		Creates a category
// @Description	Creates a category
// @ID			create-category
// @Tags      	category
// @Accept		json
// @Produce		json
// @Param		categoryBody	body	models.CategoryRequestBody	true	"Category Body"
// @Success		201	  {object}	  models.Category
// @Failure     400   {string}    error
// @Failure     401   {string}    error
// @Failure     404   {string}    error
// @Failure     409   {string}    error
// @Failure     500   {string}    error
// @Router		/categories/  [post]
func (cat *CategoryController) CreateCategory(c *fiber.Ctx) error {
	var categoryBody models.CategoryRequestBody

	if err := c.BodyParser(&categoryBody); err != nil {
		return utilities.InvalidJSON()
	}

	newCategory, err := cat.categoryService.CreateCategory(categoryBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(newCategory)
}

// GetCategories godoc
//
// @Summary		Retrieve all categories
// @Description	Retrieves all categories
// @ID			get-categories
// @Tags      	category
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	    []models.Category
// @Failure     400   {string}      error
// @Failure     404   {string}      error
// @Failure     500   {string}      error
// @Router		/categories/  [get]
func (cat *CategoryController) GetCategories(c *fiber.Ctx) error {
	categories, err := cat.categoryService.GetCategories(c.Query("limit", constants.DEFAULT_LIMIT_STRING), c.Query("page", constants.DEFAULT_PAGE_STRING))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(&categories)
}

// GetCategory godoc
//
// @Summary		Retrieve a category
// @Description	Retrieves a category
// @ID			get-category
// @Tags      	category
// @Produce		json
// @Param		categoryID	path	string	true	"Category ID"
// @Success		200	  {object}	    models.Category
// @Failure     400   {string}      error
// @Failure     404   {string}      error
// @Failure     500   {string}      error
// @Router		/categories/{categoryID}/  [get]
func (cat *CategoryController) GetCategory(c *fiber.Ctx) error {
	category, err := cat.categoryService.GetCategory(c.Params("categoryID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(&category)
}

// DeleteCategory godoc
//
// @Summary		Deletes a category
// @Description	Deletes a category
// @ID			delete-category
// @Tags      	category
// @Produce		json
// @Param		categoryID	path	string	true	"Category ID"
// @Success		204	  {string}	    utilities.SuccessResponse
// @Failure     400   {string}      error
// @Failure     401   {string}      error
// @Failure     404   {string}      error
// @Failure     500   {string}      error
// @Router		/categories/{categoryID}/  [delete]
func (cat *CategoryController) DeleteCategory(c *fiber.Ctx) error {
	if err := cat.categoryService.DeleteCategory(c.Params("categoryID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// UpdateCategory godoc
//
// @Summary		Updates a category
// @Description	Updates a category
// @ID			update-category
// @Tags      	category
// @Accept		json
// @Produce		json
// @Param		categoryID	path	string	true	"Category ID"
// @Param		categoryBody	body	models.CategoryRequestBody	true	"Category Body"
// @Success		200	  {object}	    models.Category
// @Failure     400   {string}      error
// @Failure     401   {string}      error
// @Failure     404   {string}      error
// @Failure     500   {string}      error
// @Router		/categories/{categoryID}/  [patch]
func (cat *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	var category models.CategoryRequestBody
	if err := c.BodyParser(&category); err != nil {
		return utilities.InvalidJSON()
	}

	updatedCategory, err := cat.categoryService.UpdateCategory(c.Params("categoryID"), category)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(updatedCategory)
}
