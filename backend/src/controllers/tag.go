package controllers

import (
	"backend/src/models"
	"backend/src/services"
	"backend/src/utilities"

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
// @Failure     400   {string}    string "Failed to process the request"
// @Failure     400   {string}    string "Failed to validate the data"
// @Failure     500   {string}    string "Failed to create tag"
// @Router		/api/v1/tags/  [post]
func (t *TagController) CreateTag(c *fiber.Ctx) error {
	var tag models.Tag

	if err := c.BodyParser(&tag); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to process the request")
	}

	if err := utilities.ValidateData(c, tag); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to validate the data")
	}

	dbTag, err := t.tagService.CreateTag(tag)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create tag")
	}

	return c.Status(fiber.StatusCreated).JSON(&dbTag)
}
