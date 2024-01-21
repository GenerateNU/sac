package controllers

import (
	"fmt"

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
	var categoryBody models.CreateCategoryRequestBody

	if err := c.BodyParser(&categoryBody); err != nil {
		fmt.Print(err)
		return fiber.NewError(fiber.StatusBadRequest, "failed to process the request")
	}

	category := models.Category{
		Name: categoryBody.Name,
	}

	newCategory, err := t.categoryService.CreateCategory(category)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(newCategory)
}
