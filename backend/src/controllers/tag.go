package controllers

import (
	"github.com/GenerateNU/sac/backend/src/errors"
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

// GetAllTags godoc

// CreateTag creates a new tag.
//
// @Summary		Retrieve all tags
// @Description	Retrieves all tags
// @ID			get-all-tags
// @Tags      	tag
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	    []models.Tag
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/tags/  [post]
func (t *TagController) CreateTag(c *fiber.Ctx) error {
	var tagBody models.TagRequestBody

	if err := c.BodyParser(&tagBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	dbTag, err := t.tagService.CreateTag(tagBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(&dbTag)
}

// GetTag godoc
//
// @Summary		Retrieve a tag
// @Description	Retrieves a tag
// @ID			get-tag
// @Tags      	tag
// @Produce		json
// @Param		tagID	path	string	true	"Tag ID"
// @Success		200	  {object}	    models.Tag
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/tags/{tagID}/  [get]
func (t *TagController) GetTag(c *fiber.Ctx) error {
	tag, err := t.tagService.GetTag(c.Params("tagID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&tag)
}

// UpdateTag godoc
//
// @Summary		Update a tag
// @Description	Updates a tag
// @ID			update-tag
// @Tags      	tag
// @Accept		json
// @Produce		json
// @Param		tagID	path	string	true	"Tag ID"
// @Param		tag	body	models.TagRequestBody	true	"Tag"
// @Success		200	  {object}	  models.Tag
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/tags/{tagID}/  [patch]
func (t *TagController) UpdateTag(c *fiber.Ctx) error {
	var tagBody models.TagRequestBody

	if err := c.BodyParser(&tagBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	tag, err := t.tagService.UpdateTag(c.Params("tagID"), tagBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&tag)
}

// DeleteTag godoc
//
// @Summary		Delete a tag
// @Description	Deletes a tag
// @ID			delete-tag
// @Tags      	tag
// @Produce		json
// @Param		tagID	path	string	true	"Tag ID"
// @Success		204	  {string}	    utilities.SuccessResponse
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/tags/{tagID}/  [delete]
func (t *TagController) DeleteTag(c *fiber.Ctx) error {
	err := t.tagService.DeleteTag(c.Params("tagID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
