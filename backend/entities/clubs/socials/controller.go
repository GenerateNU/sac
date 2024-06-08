package socials

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type ClubSocialController struct {
	clubSocialService ClubSocialServiceInterface
}

func NewClubSocialController(clubSocialService ClubSocialServiceInterface) *ClubSocialController {
	return &ClubSocialController{clubSocialService: clubSocialService}
}

// GetClubSocials godoc
//
// @Summary		Retrieve all socials for a club
// @Description	Retrieves all socials associated with a club
// @ID			get-socials-by-club
// @Tags      	club-social
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Success		200	  {object}	    []models.Social
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/socials/  [get]
func (cc *ClubSocialController) GetClubSocials(c *fiber.Ctx) error {
	socials, err := cc.clubSocialService.GetClubSocials(c.Params("clubID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(socials)
}

// PutSocial godoc
//
// @Summary		Creates a social
// @Description	Creates a social
// @ID			put-social
// @Tags      	club-social
// @Accept		json
// @Produce		json
// @Param		clubID		path	string	true	"Club ID"
// @Param		socialBody	body	PutSocialRequestBody	true	"Social Body"
// @Success		201	  {object}	  models.Social
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/clubs/{clubID}/socials/  [put]
func (cc *ClubSocialController) PutSocial(c *fiber.Ctx) error {
	var socialBody PutSocialRequestBody

	if err := c.BodyParser(&socialBody); err != nil {
		return utilities.InvalidJSON()
	}

	social, err := cc.clubSocialService.PutClubSocial(c.Params("clubID"), socialBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(social)
}
