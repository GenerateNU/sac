package controllers

import (
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubFollowerController struct {
	clubFollowerService services.ClubFollowerServiceInterface
}

func NewClubFollowerController(clubFollowerService services.ClubFollowerServiceInterface) *ClubFollowerController {
	return &ClubFollowerController{clubFollowerService: clubFollowerService}
}

func (cf *ClubFollowerController) GetUserFollowingClubs(c *fiber.Ctx) error {
	clubs, err := cf.clubFollowerService.GetUserFollowingClubs(c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(&clubs)
}
