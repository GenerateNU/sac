package controllers

import (
	"backend/src/models"
	"backend/src/services"

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
// @Success		200	  {object}	  []models.Category
// @Failure 	400	  {string}	  string "Failed to process the request"
// @Failure 	400	  {string}	  string "Failed to validate data"
// @Failure     500   {string}    string "Failed to create category"
// @Router		/api/v1/category/  [post]
func (t *CategoryController) CreateCategory(c *fiber.Ctx) error {
	var categoryBody models.CategoryPartial

	// Process the category data from the request body:
	if err := c.BodyParser(&categoryBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to process the request")
	}

	// Transform the accepted data into a category:
	category := models.Category {
		Name: categoryBody.Name,
	}


	newCategory, err := t.categoryService.CreateCategory(category)
	if err != nil { 
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(newCategory)
}
