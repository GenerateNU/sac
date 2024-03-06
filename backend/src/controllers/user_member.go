package controllers

import (
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

type UserMemberController struct {
	clubMemberService services.UserMemberServiceInterface
}

func NewUserMemberController(clubMemberService services.UserMemberServiceInterface) *UserMemberController {
	return &UserMemberController{clubMemberService: clubMemberService}
}

// CreateMembership godoc
//
// @Summary		Join a club
// @Description	Join a club
// @ID			create-membership
// @Tags      	user-member
// @Accept		json
// @Produce		json
// @Param		userID		path	string	true	"User ID"
// @Param		clubID		path	string	true	"Club ID"
// @Success		201	  {object}	    utilities.SuccessResponse
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/{userID}/member/{clubID}/  [post]
func (um *UserMemberController) CreateMembership(c *fiber.Ctx) error {
	err := um.clubMemberService.CreateMembership(c.Params("userID"), c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return utilities.FiberMessage(c, fiber.StatusCreated, "Successfully joined club")
}

// DeleteMembership godoc
//
// @Summary		Leave a club
// @Description	Leave a club
// @ID			delete-membership
// @Tags      	user-member
// @Accept		json
// @Produce		json
// @Param		userID		path	string	true	"User ID"
// @Param		clubID		path	string	true	"Club ID"
// @Success		204	  {object}	    utilities.SuccessResponse
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/{userID}/member/{clubID}/  [delete]
func (um *UserMemberController) DeleteMembership(c *fiber.Ctx) error {
	err := um.clubMemberService.DeleteMembership(c.Params("userID"), c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetMembership godoc
//
// @Summary		Retrieve all clubs a user is a member of
// @Description	Retrieves all clubs a user is a member of
// @ID			get-membership
// @Tags      	user-member
// @Produce		json
// @Param		userID	path	string	true	"User ID"
// @Success		200	  {object}	    []models.Club
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/users/{userID}/member/  [get]
func (um *UserMemberController) GetMembership(c *fiber.Ctx) error {
	followers, err := um.clubMemberService.GetMembership(c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(followers)
}
