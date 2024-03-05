package controllers

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type UserTagController struct {
	userTagService services.UserTagServiceInterface
}

func NewUserTagController(userTagService services.UserTagServiceInterface) *UserTagController {
	return &UserTagController{userTagService: userTagService}
}

// GetUserTags godoc
//
// @Summary		Retrieve all tags for a user
// @Description	Retrieves all tags associated with a user
// @ID			get-tags-by-user
// @Tags      	user-tag
// @Produce		json
// @Param		userID	path	string	true	"User ID"
// @Success		200	  {object}	    []models.Tag
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/{userID}/tags/  [get]
func (ut *UserTagController) GetUserTags(c *fiber.Ctx) error {
	tags, err := ut.userTagService.GetUserTags(c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(&tags)
}

// CreateUserTags godoc
//
// @Summary		Create user tags
// @Description	Creates tags for a user
// @ID			create-user-tags
// @Tags      	user-tag
// @Accept		json
// @Produce		json
// @Param		userID		path	string	true	"User ID"
// @Param		userTagsBody	body	models.CreateUserTagsBody	true	"User Tags Body"
// @Success		201	  {object}	  []models.Tag
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/users/{userID}/tags/  [post]
func (ut *UserTagController) CreateUserTags(c *fiber.Ctx) error {
	var requestBody models.CreateUserTagsBody
	if err := c.BodyParser(&requestBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	tags, err := ut.userTagService.CreateUserTags(c.Params("userID"), requestBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(&tags)
}

// DeleteUserTag godoc
//
// @Summary		Create user tags
// @Description	Creates tags for a user
// @ID			create-user-tags
// @Tags      	user-tag
// @Accept		json
// @Produce		json
// @Param		userID		path	string	true	"User ID"
// @Success		201	  {object}	  []models.Tag
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/users/{userID}/tags/  [post]
func (ut *UserTagController) DeleteUserTag(c *fiber.Ctx) error {
	err := ut.userTagService.DeleteUserTag(c.Params("userID"), c.Params("tagID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
