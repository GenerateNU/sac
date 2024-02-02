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
// @Summary		Create a category
// @Description	Creates a category that is used to group tags
// @ID			create-category
// @Tags      	category
// @Produce		json
// @Success		201	  {object}	  models.Category
// @Failure		400	  {string}	  string "failed to process the request"
// @Failure 	400	  {string}	  string "failed to validate data"
// @Failure		400	  {string}	  string "category with that name already exists"
// @Failure     500   {string}    string "failed to create category"
// @Router		/api/v1/category/  [post]
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
// @Description	Retrieves all existing categories
// @ID			get-categories
// @Tags      	category
// @Produce		json
// @Success		200	  {object}	  []models.Category
// @Failure     500   {string}    string "unable to retrieve categories"
// @Router		/api/v1/category/  [get]
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
// @Description	Retrieve a category by its ID
// @ID			get-category
// @Tags      	category
// @Produce		json
// @Success		200	  {object}	  models.Category
// @Failure 	400   {string}    string "failed to validate id"
// @Failure     404   {string}    string "faied to find category"
// @Failure     500   {string}    string "failed to retrieve category"
// @Router		/api/v1/category/{id}  [get]
func (cat *CategoryController) GetCategory(c *fiber.Ctx) error {
	category, err := cat.categoryService.GetCategory(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&category)
}

// DeleteCategory godoc
//
// @Summary		Delete a category
// @Description	Delete a category by ID
// @ID			delete-category
// @Tags      	category
// @Produce		json
// @Success		204	  {object}
// @Failure 	400   {string}    string "failed to validate id"
// @Failure     404   {string}    string "failed to find category"
// @Failure     500   {string}    string "failed to delete category"
// @Router		/api/v1/category/{id}  [delete]
func (cat *CategoryController) DeleteCategory(c *fiber.Ctx) error {
	if err := cat.categoryService.DeleteCategory(c.Params("id")); err != nil {
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
// @Produce		json
// @Success		200	  {object}	  models.Category
// @Failure 	400   {string}    string "failed to validate id"
// @Failure     404   {string}    string "failed to find category"
// @Failure     500   {string}    string "failed to update category"
// @Router		/api/v1/category/{id}  [patch]
func (cat *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	var category models.CategoryRequestBody

	if err := c.BodyParser(&category); err != nil {
		return errors.FailedToValidateCategory.FiberError(c)
	}

	updatedCategory, err := cat.categoryService.UpdateCategory(c.Params("id"), category)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedCategory)
}
