package tag

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type UserTagController struct {
	userTagService UserTagServiceInterface
}

func NewUserTagController(userTagService UserTagServiceInterface) *UserTagController {
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
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/users/{userID}/tags/  [get]
func (ut *UserTagController) GetUserTags(c *fiber.Ctx) error {
	tags, err := ut.userTagService.GetUserTags(c.Params("userID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(&tags)
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
// @Param		userTagsBody	body	CreateUserTagsBody	true	"User Tags Body"
// @Success		201	  {object}	  []models.Tag
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/users/{userID}/tags/  [post]
func (ut *UserTagController) CreateUserTags(c *fiber.Ctx) error {
	var requestBody CreateUserTagsBody

	if err := c.BodyParser(&requestBody); err != nil {
		return utilities.InvalidJSON()
	}

	tags, err := ut.userTagService.CreateUserTags(c.Params("userID"), requestBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(&tags)
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
// @Success		201
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/users/{userID}/tags/  [delete]
func (ut *UserTagController) DeleteUserTag(c *fiber.Ctx) error {
	err := ut.userTagService.DeleteUserTag(c.Params("userID"), c.Params("tagID"))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
