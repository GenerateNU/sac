package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type ClubEventController struct {
	clubEventService services.ClubEventServiceInterface
}

func NewClubEventController(clubEventService services.ClubEventServiceInterface) *ClubEventController {
	return &ClubEventController{clubEventService: clubEventService}
}

// GetClubEvents godoc
//
// @Summary		Retrieve all events for a club
// @Description	Retrieves all events associated with a club
// @ID			get-events-by-club
// @Tags      	club-event
// @Produce		json
// @Param		clubID	path	string	true	"Club ID"
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	    []models.Event
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/clubs/{clubID}/events/  [get]
func (cl *ClubEventController) GetClubEvents(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	if events, err := cl.clubEventService.GetClubEvents(c.Params("clubID"), c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage))); err != nil {
		return err.FiberError(c)
	} else {
		return c.Status(fiber.StatusOK).JSON(events)
	}
}
