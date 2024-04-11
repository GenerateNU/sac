package controllers

import (
	"github.com/GenerateNU/sac/backend/src/constants"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubMemberController struct {
	clubMemberService services.ClubMemberServiceInterface
}

func NewClubMemberController(clubMemberService services.ClubMemberServiceInterface) *ClubMemberController {
	return &ClubMemberController{clubMemberService: clubMemberService}
}

// GetClubMembers godoc
//
// @Summary		Retrieve all members for a club
// @Description	Retrieves all members associated with a club
// @ID			get-members-by-club
// @Tags      	club-member
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	    []models.User
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/{clubID}/members/  [get]
func (cm *ClubMemberController) GetClubMembers(c *fiber.Ctx) error {
	followers, err := cm.clubMemberService.GetClubMembers(c.Params("clubID"), c.Query("limit", constants.DEFAULT_LIMIT_STRING), c.Query("page", constants.DEFAULT_PAGE_STRING))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(followers)
}
