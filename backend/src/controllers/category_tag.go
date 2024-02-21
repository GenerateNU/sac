package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/services"

	"github.com/gofiber/fiber/v2"
)

type CategoryTagController struct {
	categoryTagService services.CategoryTagServiceInterface
}

func NewCategoryTagController(categoryTagService services.CategoryTagServiceInterface) *CategoryTagController {
	return &CategoryTagController{categoryTagService: categoryTagService}
}

// GetTagsByCategory godoc
//
// @Summary		Retrieve all tags by category
// @Description	Retrieves all tags associated with a category
// @ID			get-tags-by-category
// @Tags      	category-tag
// @Produce		json
// @Param		categoryID	path	string	true	"Category ID"
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	    []models.Tag
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/category/{categoryID}/tags  [get]
func (ct *CategoryTagController) GetTagsByCategory(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	tags, err := ct.categoryTagService.GetTagsByCategory(c.Params("categoryID"), c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&tags)
}

// GetTagByCategory godoc
//
// @Summary		Retrieve a tag by category
// @Description	Retrieves a tag associated with a category
// @ID			get-tag-by-category
// @Tags      	category-tag
// @Produce		json
// @Param		categoryID	path	string	true	"Category ID"
// @Param		tagID		path	string	true	"Tag ID"
// @Success		200	  {object}	    models.Tag
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/category/{categoryID}/tags/{tagID}  [get]
func (ct *CategoryTagController) GetTagByCategory(c *fiber.Ctx) error {
	tag, err := ct.categoryTagService.GetTagByCategory(c.Params("categoryID"), c.Params("tagID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&tag)
}
