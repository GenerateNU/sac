package controllers

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubController struct {
	clubService services.ClubServiceInterface
}

func NewClubController(clubService services.ClubServiceInterface) *ClubController {
	return &ClubController{clubService: clubService}
}

// GetClubs godoc
//
// @Summary		Retrieve all clubs
// @Description	Retrieves all clubs
// @ID			get-all-clubs
// @Tags      	club
// @Accept		json
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	    []models.Club
// @Failure     400   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/  [get]
func (cl *ClubController) GetClubs(c *fiber.Ctx) error {
	var queryParams models.ClubQueryParams

	queryParams.Limit = 10 // default limit
	queryParams.Page = 1   // default page

	if err := c.QueryParser(&queryParams); err != nil {
		return errors.FailedtoParseQueryParams.FiberError(c)
	}

	clubs, err := cl.clubService.GetClubs(&queryParams)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(clubs)
}

// CreateClub godoc
//
// @Summary		Create a club
// @Description	Creates a club
// @ID			create-club
// @Tags      	club
// @Accept		json
// @Produce		json
// @Param		club	body	models.CreateClubRequestBody	true	"Club"
// @Success		201	  {object}	  models.Club
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/clubs/  [post]
func (cl *ClubController) CreateClub(c *fiber.Ctx) error {
	var clubBody models.CreateClubRequestBody
	
	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	club, err := cl.clubService.CreateClub(clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(club)
}

// GetClub godoc
//
// @Summary		Retrieve a club
// @Description	Retrieves a club
// @ID			get-club
// @Tags      	club
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Success		200	  {object}	    models.Club
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/{clubID}/  [get]
func (cl *ClubController) GetClub(c *fiber.Ctx) error {
	club, err := cl.clubService.GetClub(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(club)
}

// UpdateClub godoc
//
// @Summary		Update a club
// @Description	Updates a club
// @ID			update-club
// @Tags      	club
// @Accept		json
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		club	body	models.UpdateClubRequestBody	true	"Club"
// @Success		200	  {object}	  models.Club
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/clubs/{clubID}/  [patch]
func (cl *ClubController) UpdateClub(c *fiber.Ctx) error {
	var clubBody models.UpdateClubRequestBody

	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedClub, err := cl.clubService.UpdateClub(c.Params("clubID"), clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedClub)
}

// DeleteClub godoc
//
// @Summary		Delete a club
// @Description	Deletes a club
// @ID			delete-club
// @Tags      	club
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Success		204	  {string}	    utilities.SuccessResponse
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/{clubID}/  [delete]
func (cl *ClubController) DeleteClub(c *fiber.Ctx) error {
	err := cl.clubService.DeleteClub(c.Params("clubID"))
	if err != nil {
		return err.FiberError(c)
	}
	return c.SendStatus(fiber.StatusNoContent)
}