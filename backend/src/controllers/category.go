package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	categoryService services.CategoryServiceInterface
}

func NewCategoryController(categoryService services.CategoryServiceInterface) *CategoryController {
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
// @Failure     400   {string}    errors.Error
// @Failure     401   {string}    errors.Error
// @Failure     404   {string}    errors.Error
// @Failure     500   {string}    errors.Error
// @Router		/categories/  [post]
func (cat *CategoryController) CreateCategory(c *fiber.Ctx) error {
	var categoryBody models.CategoryRequestBody

	if err := c.BodyParser(&categoryBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	newCategory, err := cat.categoryService.CreateCategory(categoryBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(newCategory)
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
// @Failure     400   {string}      errors.Error
// @Failure     404   {string}      errors.Error
// @Failure     500   {string}      errors.Error
// @Router		/categories/  [get]
func (cat *CategoryController) GetCategories(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	categories, err := cat.categoryService.GetCategories(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&categories)
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
// @Failure     400   {string}      errors.Error
// @Failure     404   {string}      errors.Error
// @Failure     500   {string}      errors.Error
// @Router		/categories/{categoryID}/  [get]
func (cat *CategoryController) GetCategory(c *fiber.Ctx) error {
	category, err := cat.categoryService.GetCategory(c.Params("categoryID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&category)
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
// @Failure     400   {string}      errors.Error
// @Failure     401   {string}      errors.Error
// @Failure     404   {string}      errors.Error
// @Failure     500   {string}      errors.Error
// @Router		/categories/{categoryID}/  [delete]
func (cat *CategoryController) DeleteCategory(c *fiber.Ctx) error {
	if err := cat.categoryService.DeleteCategory(c.Params("categoryID")); err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
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
// @Failure     400   {string}      errors.Error
// @Failure     401   {string}      errors.Error
// @Failure     404   {string}      errors.Error
// @Failure     500   {string}      errors.Error
// @Router		/categories/{categoryID}/  [patch]
func (cat *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	var category models.CategoryRequestBody

	if err := c.BodyParser(&category); err != nil {
		return errors.FailedToValidateCategory.FiberError(c)
	}

	updatedCategory, err := cat.categoryService.UpdateCategory(c.Params("categoryID"), category)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedCategory)
}
