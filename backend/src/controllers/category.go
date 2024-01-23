package controllers

import (
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
func (t *CategoryController) CreateCategory(c *fiber.Ctx) error {
	var categoryBody models.CategoryRequestBody

	if err := c.BodyParser(&categoryBody); err != nil {
		return errors.Error{StatusCode: fiber.StatusBadRequest, Message: "failed to process the request"}.FiberError(c)
	}

	newCategory, err := t.categoryService.CreateCategory(categoryBody)

	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(newCategory)
}
