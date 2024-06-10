package members

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/gofiber/fiber/v2"
)

type ClubMemberController struct {
	clubMemberService ClubMemberServiceInterface
}

func NewClubMemberController(clubMemberService ClubMemberServiceInterface) *ClubMemberController {
	return &ClubMemberController{clubMemberService: clubMemberService}
}

func (cm *ClubMemberController) GetClubMembers(c *fiber.Ctx) error {
	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPageInfo
	}

	followers, err := cm.clubMemberService.GetClubMembers(c.Params("clubID"), *pageInfo)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(followers)
}

func (cm *ClubMemberController) CreateClubMember(c *fiber.Ctx) error {
	if err := cm.clubMemberService.CreateClubMember(c.Params("clubID"), c.Params("userID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusCreated)
}

func (cm *ClubMemberController) DeleteClubMember(c *fiber.Ctx) error {
	if err := cm.clubMemberService.DeleteClubMember(c.Params("clubID"), c.Params("userID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
