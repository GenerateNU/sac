package controllers

import (
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

type UserFollowerController struct {
	userFollowerService services.UserFollowerServiceInterface
}

func NewUserFollowerController(userFollowerService services.UserFollowerServiceInterface) *UserFollowerController {
	return &UserFollowerController{userFollowerService: userFollowerService}
}

// CreateFollowing godoc
//
// @Summary		Follow a club
// @Description	Follow a club
// @ID			create-following
// @Tags      	user-follower
// @Accept		json
// @Produce		json
// @Param		userID		path	string	true	"User ID"
// @Param		clubID		path	string	true	"Club ID"
// @Success		201	  {object}	    utilities.SuccessResponse
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/{userID}/follower/{clubID}/  [post]
func (uf *UserFollowerController) CreateFollowing(c *fiber.Ctx) error {
	err := uf.userFollowerService.CreateFollowing(c.Params("userID"), c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}
	return utilities.FiberMessage(c, fiber.StatusCreated, "Successfully followed club")
}

// DeleteFollowing godoc
//
// @Summary		Unfollow a club
// @Description	Unfollow a club
// @ID			delete-following
// @Tags      	user-follower
// @Accept		json
// @Produce		json
// @Param		userID		path	string	true	"User ID"
// @Param		clubID		path	string	true	"Club ID"
// @Success		204	  {object}	    utilities.SuccessResponse
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/{userID}/follower/{clubID}/  [delete]
func (uf *UserFollowerController) DeleteFollowing(c *fiber.Ctx) error {
	err := uf.userFollowerService.DeleteFollowing(c.Params("userID"), c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// GetAllFollowing godoc
//
// @Summary		Retrieve all clubs a user is following
// @Description	Retrieves all clubs a user is following
// @ID			get-following
// @Tags      	user-follower
// @Produce		json
// @Param		userID	path	string	true	"User ID"
// @Success		200	  {object}	    []models.Club
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/{userID}/follower/  [get]
func (uf *UserFollowerController) GetFollowing(c *fiber.Ctx) error {
	clubs, err := uf.userFollowerService.GetFollowing(c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.Status(fiber.StatusOK).JSON(clubs)
}
