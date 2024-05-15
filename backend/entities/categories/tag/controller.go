package tag

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/gofiber/fiber/v2"
)

type CategoryTagController struct {
	categoryTagService CategoryTagServiceInterface
}

func NewCategoryTagController(categoryTagService CategoryTagServiceInterface) *CategoryTagController {
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
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/categories/{categoryID}/tags/  [get]
func (ct *CategoryTagController) GetTagsByCategory(c *fiber.Ctx) error {
	pagination, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPagination
	}

	tags, err := ct.categoryTagService.GetTagsByCategory(c.Params("categoryID"), *pagination)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(&tags)
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
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/categories/{categoryID}/tags/{tagID}/  [get]
func (ct *CategoryTagController) GetTagByCategory(c *fiber.Ctx) error {
	tag, err := ct.categoryTagService.GetTagByCategory(c.Params("categoryID"), c.Params("tagID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(&tag)
}
