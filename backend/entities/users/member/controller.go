package member

import (
	"github.com/gofiber/fiber/v2"
)

type UserMemberController struct {
	clubMemberService UserMemberServiceInterface
}

func NewUserMemberController(clubMemberService UserMemberServiceInterface) *UserMemberController {
	return &UserMemberController{clubMemberService: clubMemberService}
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
