package members

import (
	"net/http"

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
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/users/{userID}/member/  [get]
func (um *UserMemberController) GetMembership(c *fiber.Ctx) error {
	followers, err := um.clubMemberService.GetMembership(c.Params("userID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(followers)
}
