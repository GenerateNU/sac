package member

import (
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/gofiber/fiber/v2"
)

type ClubMemberController struct {
	clubMemberService ClubMemberServiceInterface
}

func NewClubMemberController(clubMemberService ClubMemberServiceInterface) *ClubMemberController {
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

// CreateClubMember godoc
//
// @Summary		Create a new member for a club
// @Description	Creates a new member associated with a club
// @ID			create-member-for-club
// @Tags      	club-member
// @Accept		json
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		userID	path	string	true	"User ID"
// @Success		201	  {object}	    models.User
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/{clubID}/members/  [post]
func (cm *ClubMemberController) CreateClubMember(c *fiber.Ctx) error {
	err := cm.clubMemberService.CreateClubMember(c.Params("clubID"), c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusCreated)
}

// DeleteClubMember godoc
//
// @Summary		Delete a member from a club
// @Description	Deletes a member associated with a club
// @ID			delete-member-from-club
// @Tags      	club-member
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		userID	path	string	true	"User ID"
// @Success		204	  {object}	    models.User
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/{clubID}/members/  [delete]
func (cm *ClubMemberController) DeleteClubMember(c *fiber.Ctx) error {
	err := cm.clubMemberService.DeleteClubMember(c.Params("clubID"), c.Params("userID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
