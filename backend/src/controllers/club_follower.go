package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubFollowerController struct {
	clubFollowerService services.ClubFollowerServiceInterface
}

func NewClubFollowerController(clubFollowerService services.ClubFollowerServiceInterface) *ClubFollowerController {
	return &ClubFollowerController{clubFollowerService: clubFollowerService}
}

func (cf *ClubFollowerController) GetClubFollowers(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	followers, err := cf.clubFollowerService.GetClubFollowers(c.Params("clubID"), c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(followers)
}
