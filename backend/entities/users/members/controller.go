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

func (um *UserMemberController) GetMembership(c *fiber.Ctx) error {
	followers, err := um.clubMemberService.GetMembership(c.Params("userID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(followers)
}
