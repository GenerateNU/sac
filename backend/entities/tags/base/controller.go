package base

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
)

type TagController struct {
	tagService TagServiceInterface
}

func NewTagController(tagService TagServiceInterface) *TagController {
	return &TagController{tagService: tagService}
}

// GetTags godoc
//
// @Summary		Retrieve all tags
// @Description	Retrieves all tags
// @ID			get-all-tags
// @Tags      	tag
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	    []models.Tag
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/tags  [get]
func (t *TagController) GetTags(c *fiber.Ctx) error {
	tags, err := t.tagService.GetTags()
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(&tags)
}

// CreateTag godoc
//
// @Summary		Create a tag
// @Description	Creates a tag
// @ID			create-tag
// @Tags      	tag
// @Accept		json
// @Produce		json
// @Param		tagBody	body	CreateTagRequestBody	true	"Tag Body"
// @Success		201	  {object}	  models.Tag
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/tags/  [post]
func (t *TagController) CreateTag(c *fiber.Ctx) error {
	var tagBody CreateTagRequestBody

	if err := c.BodyParser(&tagBody); err != nil {
		return utilities.InvalidJSON()
	}

	dbTag, err := t.tagService.CreateTag(tagBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(&dbTag)
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
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/tags/{tagID}/  [get]
func (t *TagController) GetTag(c *fiber.Ctx) error {
	tag, err := t.tagService.GetTag(c.Params("tagID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(&tag)
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
// @Param		tag	body	UpdateTagRequestBody	true	"Tag"
// @Success		200	  {object}	  models.Tag
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/tags/{tagID}/  [patch]
func (t *TagController) UpdateTag(c *fiber.Ctx) error {
	var tagBody UpdateTagRequestBody

	if err := c.BodyParser(&tagBody); err != nil {
		return utilities.InvalidJSON()
	}

	tag, err := t.tagService.UpdateTag(c.Params("tagID"), tagBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(&tag)
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
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/tags/{tagID}/  [delete]
func (t *TagController) DeleteTag(c *fiber.Ctx) error {
	err := t.tagService.DeleteTag(c.Params("tagID"))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
