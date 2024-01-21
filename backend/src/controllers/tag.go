package controllers

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"

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
	var tagBody models.CreateTagRequestBody

	if err := c.BodyParser(&tagBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to process the request")
	}

	dbTag, err := t.tagService.CreateTag(tagBody)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&dbTag)
}

// GetTag godoc
//
// @Summary		Gets a tag
// @Description	Returns a tag
// @ID			get-tag
// @Tags      	tag
// @Produce		json
// @Param		id	path	int	true	"Tag ID"
// @Success		200	  {object}    models.Tag
// @Failure     400   {string}    string "failed to validate id"
// @Failure     404   {string}    string "faied to find tag"
// @Failure     500   {string}    string "failed to retrieve tag"
// @Router		/api/v1/tags/{id}  [get]
func (t *TagController) GetTag(c *fiber.Ctx) error {
	tag, err := t.tagService.GetTag(c.Params("id"))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&tag)
}
