package followers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserFollowerController struct {
	userFollowerService UserFollowerServiceInterface
}

func NewUserFollowerController(userFollowerService UserFollowerServiceInterface) *UserFollowerController {
	return &UserFollowerController{userFollowerService: userFollowerService}
}

func (ufc *UserFollowerController) GetFollowing(c *fiber.Ctx) error {
	followers, err := ufc.userFollowerService.GetFollowing(c.Params("userID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(followers)
}
