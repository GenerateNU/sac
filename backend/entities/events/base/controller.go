package base

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/entities/events"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/gofiber/fiber/v2"
)

type EventController struct {
	eventService EventServiceInterface
}

func NewEventController(eventService EventServiceInterface) *EventController {
	return &EventController{eventService: eventService}
}

// GetAllEvents godoc
//
// @Summary		Retrieve all events
// @Description	Retrieves all events
// @ID			get-all-events
// @Tags      	event
// @Produce		json
// @Param		limit		query	int	    false	"Limit"
// @Param		page		query	int	    false	"Page"
// @Success		200	  {object}	    []models.Event
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/events/  [get]
func (e *EventController) GetAllEvents(c *fiber.Ctx) error {
	pagination, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPagination
	}

	events, err := e.eventService.GetEvents(*pagination)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(events)
}

// GetEvent godoc
//
// @Summary		Retrieve an event
// @Description	Retrieves an event
// @ID			get-event
// @Tags      	event
// @Produce		json
// @Param		eventID	path	string	true	"Event ID"
// @Success		200	  {object}	    models.Event
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/events/{eventID}/  [get]
func (e *EventController) GetEvent(c *fiber.Ctx) error {
	event, err := e.eventService.GetEvent(c.Params("eventID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(event)
}

func (e *EventController) CreateEvent(c *fiber.Ctx) error {
	var body events.CreateEventRequestBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	event, err := e.eventService.CreateEvent(body)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(event)
}

func (e *EventController) UpdateEvent(c *fiber.Ctx) error {
	var body events.UpdateEventRequestBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	event, err := e.eventService.UpdateEvent(c.Params("eventID"), body)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(event)
}

func (e *EventController) DeleteEvent(c *fiber.Ctx) error {
	err := e.eventService.DeleteEvent(c.Params("eventID"))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
