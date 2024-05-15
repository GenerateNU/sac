package event

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/gofiber/fiber/v2"
)

type ClubEventController struct {
	clubEventService ClubEventServiceInterface
}

func NewClubEventController(clubEventService ClubEventServiceInterface) *ClubEventController {
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
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/clubs/{clubID}/events/  [get]
func (cl *ClubEventController) GetClubEvents(c *fiber.Ctx) error {
	pagination, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPagination
	}

	if events, err := cl.clubEventService.GetClubEvents(c.Params("clubID"), *pagination); err != nil {
		return err
	} else {
		return c.Status(http.StatusOK).JSON(events)
	}
}
