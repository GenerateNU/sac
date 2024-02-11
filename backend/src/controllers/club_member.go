package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubMemberController struct {
	clubMemberService services.ClubMemberServiceInterface
}

func NewClubMemberController(clubMemberService services.ClubMemberServiceInterface) *ClubMemberController {
	return &ClubMemberController{clubMemberService: clubMemberService}
}

func (cm *ClubMemberController) GetClubMembers(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	followers, err := cm.clubMemberService.GetClubMembers(c.Params("clubID"), c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(followers)
}
