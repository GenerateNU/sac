package controllers

import (
	"backend/src/models"
	"backend/src/services"

	"github.com/gofiber/fiber/v2"
)

type TagController struct {
	tagService services.TagServiceInterface
}

func NewTagController(tagService services.TagServiceInterface) *TagController {
	return &TagController{tagService: tagService}
}

// CreateTag godoc
//
// @Summary		Creates a tag
// @Description	Creates a tag
// @ID			create-tag
// @Tags      	tag
// @Accept		json
// @Produce		json
// @Success		201	  {object}    models.Tag
// @Failure     400   {string}    string "failed to process the request"
// @Failure     400   {string}    string "failed to validate the data"
// @Failure     500   {string}    string "failed to create tag"
// @Router		/api/v1/tags/  [post]
func (t *TagController) CreateTag(c *fiber.Ctx) error {
	var tag models.TagCreateRequestBody

	if err := c.BodyParser(&tag); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to process the request")
	}

	partialTag := models.Tag{
		Name:       tag.Name,
		CategoryID: tag.CategoryID,
	}

	dbTag, err := t.tagService.CreateTag(partialTag)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&dbTag)
}
