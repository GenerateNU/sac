package follower

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/gofiber/fiber/v2"
)

type ClubFollowerController struct {
	clubFollowerService ClubFollowerServiceInterface
}

func NewClubFollowerController(clubFollowerService ClubFollowerServiceInterface) *ClubFollowerController {
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
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/followers/  [get]
func (cf *ClubFollowerController) GetClubFollowers(c *fiber.Ctx) error {
	followers, err := cf.clubFollowerService.GetClubFollowers(c.Params("clubID"), c.Query("limit", constants.DEFAULT_LIMIT_STRING), c.Query("page", constants.DEFAULT_PAGE_STRING))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(followers)
}
