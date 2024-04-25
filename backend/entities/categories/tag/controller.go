package tag

import (
	"github.com/GenerateNU/sac/backend/constants"

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
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/categories/{categoryID}/tags/  [get]
func (ct *CategoryTagController) GetTagsByCategory(c *fiber.Ctx) error {
	tags, err := ct.categoryTagService.GetTagsByCategory(c.Params("categoryID"), c.Query("limit", constants.DEFAULT_LIMIT_STRING), c.Query("page", constants.DEFAULT_PAGE_STRING))
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
// @Router		/categories/{categoryID}/tags/{tagID}/  [get]
func (ct *CategoryTagController) GetTagByCategory(c *fiber.Ctx) error {
	tag, err := ct.categoryTagService.GetTagByCategory(c.Params("categoryID"), c.Params("tagID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&tag)
}
