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

// GetClubFollowers godoc
//
// @Summary		Retrieve all followers for a club
// @Description	Retrieves all followers associated with a club
// @ID			get-followers-by-club
// @Tags      	club-follower
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	    []models.User
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/{clubID}/followers/  [get]
func (cf *ClubFollowerController) GetClubFollowers(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	followers, err := cf.clubFollowerService.GetClubFollowers(c.Params("clubID"), c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(followers)
}
