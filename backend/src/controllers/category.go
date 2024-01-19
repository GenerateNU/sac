package controllers

import (
	"backend/src/models"
	"backend/src/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		return fiber.NewError(fiber.StatusBadRequest, "failed to process the request")
	}

	titleCasedCategoryName := cases.Title(language.English).String(categoryBody.Name)

	category := models.Category{
		Name: titleCasedCategoryName,
	}

	newCategory, err := t.categoryService.CreateCategory(category)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(newCategory)
}
