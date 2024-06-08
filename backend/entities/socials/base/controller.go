package base

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/gofiber/fiber/v2"
)

type SocialController struct {
	socialService SocialServiceInterface
}

func NewSocialController(socialService SocialServiceInterface) *SocialController {
	return &SocialController{socialService: socialService}
}

// GetSocial godoc
//
// @Summary		Retrieves a social
// @Description	Retrieves a social by id
// @ID			get-social
// @Tags      	social
// @Accept		json
// @Produce		json
// @Param		socialID	path	string	true	"Social ID"
// @Success		201	  {object}	  models.Social
// @Failure     400   {string}    error
// @Failure     404   {string}    error
// @Failure     500   {string}    error
// @Router		/socials/{socialID}/  [get]
func (co *SocialController) GetSocial(c *fiber.Ctx) error {
	social, err := co.socialService.GetSocial(c.Params("socialID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(social)
}

// GetSocials godoc
//
// @Summary		Retrieve all socials
// @Description	Retrieves all socials
// @ID			get-socials
// @Tags      	social
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	  []models.Social
// @Failure     400   {string}    error
// @Failure     404   {string}    error
// @Failure     500   {string}    error
// @Router		/socials/  [get]
func (co *SocialController) GetSocials(c *fiber.Ctx) error {
	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPageInfo
	}

	socials, err := co.socialService.GetSocials(*pageInfo)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(socials)
}

// DeleteSocial godoc
//
// @Summary		Deletes a social
// @Description	Deletes a social
// @ID			delete-social
// @Tags      	social
// @Accept		json
// @Produce		json
// @Param		socialID	path	string	true	"Social ID"
// @Success		201	  {object}	  models.Social
// @Failure     400   {string}    error
// @Failure     404   {string}    error
// @Failure     500   {string}    error
// @Router		/socials/{socialID}/  [delete]
func (co *SocialController) DeleteSocial(c *fiber.Ctx) error {
	err := co.socialService.DeleteSocial(c.Params("socialID"))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
