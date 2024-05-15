package base

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/entities/models"
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

// GetSeriesByEventID godoc
//
// @Summary		Retrieve all series by event
// @Description	Retrieves all series associated with an event
// @ID			get-series-by-event
// @Tags      	event
// @Produce		json
// @Param		eventID	path	string	true	"Event ID"
// @Success		200	  {object}	    []models.Series
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/events/{eventID}/series/  [get]
func (e *EventController) GetSeriesByEventID(c *fiber.Ctx) error {
	events, err := e.eventService.GetSeriesByEventID(c.Params("eventID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(events)
}

// GetSeriesByID godoc
//
// @Summary		Retrieve a series by ID
// @Description	Retrieves a series by ID
// @ID			get-series-by-id
// @Tags      	event
// @Produce		json
// @Param		eventID  	path	string	true	"Event ID"
// @Param		seriesID	path	string	true	"Series ID"
// @Success		200	  {object}	    models.Series
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/events/{eventID}/series/{seriesID}/  [get]
func (e *EventController) GetSeriesByID(c *fiber.Ctx) error {
	events, err := e.eventService.GetSeriesByID(c.Params("seriesID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(events)
}

// CreateEvent godoc
//
// @Summary		Create an event
// @Description	Creates an event
// @ID			create-event
// @Tags      	event
// @Accept		json
// @Produce		json
// @Param		event	body	models.CreateEventRequestBody	true	"Event Body"
// @Success		201	  {object}	  models.Event
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/events/  [post]
func (e *EventController) CreateEvent(c *fiber.Ctx) error {
	var eventBody models.CreateEventRequestBody
	if err := c.BodyParser(&eventBody); err != nil {
		return utilities.InvalidJSON()
	}

	event, err := e.eventService.CreateEvent(eventBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(event)
}

// UpdateEvent godoc
//
// @Summary		Create a series
// @Description	Creates a series
// @ID			create-series
// @Tags      	event
// @Accept		json
// @Produce		json
// @Param		eventID	path	string	true	"Event ID"
// @Param		seriesBody	body	models.UpdateEventRequestBody	true	"Series Body"
// @Success		201	  {object}	  models.Series
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/events/{eventID}/series/  [patch]
func (e *EventController) UpdateEvent(c *fiber.Ctx) error {
	var eventBody models.UpdateEventRequestBody
	if err := c.BodyParser(&eventBody); err != nil {
		return utilities.InvalidJSON()
	}

	updatedEvent, err := e.eventService.UpdateEvent(c.Params("eventID"), eventBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(updatedEvent)
}

// UpdateSeriesByID godoc
//
// @Summary		Update a series by ID
// @Description	Updates a series by ID. If individual events have been edited prior,
//
//	this update will override the previous changes
//
// @ID			update-series-by-id
// @Tags      	event
// @Accept		json
// @Produce		json
// @Param		seriesID	path	string	true	"Series ID"
// @Param		seriesBody	body	models.UpdateSeriesRequestBody	true	"Series Body"
// @Success		200	  {object}	  models.Series
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/events/{eventID}/series/{seriesID}/  [patch]
func (e *EventController) UpdateSeriesByID(c *fiber.Ctx) error {
	var seriesBody models.UpdateSeriesRequestBody
	if err := c.BodyParser(&seriesBody); err != nil {
		return utilities.InvalidJSON()
	}

	updatedSeries, err := e.eventService.UpdateSeries(c.Params("seriesID"), seriesBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(updatedSeries)
}

// UpdateSeriesByEventID godoc
//
// @Summary		Update a series by event ID
// @Description	Updates a series by event ID
// @ID			update-series-by-event-id
// @Tags      	event
// @Accept		json
// @Produce		json
// @Param		eventID	    path	string	true	"Event ID"
// @Param		seriesBody	body	models.UpdateSeriesRequestBody	true	"Series Body"
// @Success		200	  {object}	  models.Series
// @Failure     400   {object}    error
// @Failure     401   {object}    error
// @Failure     404   {object}    error
// @Failure     500   {object}    error
// @Router		/events/{eventID}/series/{seriesID}/  [patch]
func (e *EventController) UpdateSeriesByEventID(c *fiber.Ctx) error {
	var seriesBody models.UpdateSeriesRequestBody
	if err := c.BodyParser(&seriesBody); err != nil {
		return utilities.InvalidJSON()
	}

	updatedSeries, err := e.eventService.UpdateSeriesByEventID(c.Params("eventID"), seriesBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(updatedSeries)
}

// DeleteSeriesByID godoc
//
// @Summary		Delete a series by ID
// @Description	Deletes a series by ID
// @ID			delete-series-by-id
// @Tags      	event
// @Produce		json
// @Param		eventID  	path	string	true	"Event ID"
// @Param		seriesID	path	string	true	"Series ID"
// @Success		204	  {string}	    utilities.SuccessResponse
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/events/{eventID}/series/{seriesID}/  [delete]
func (e *EventController) DeleteSeriesByID(c *fiber.Ctx) error {
	if err := e.eventService.DeleteSeriesByID(c.Params("seriesID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// DeleteSeriesByEventID godoc
//
// @Summary		Delete all series by event
// @Description	Deletes all series associated with an event
// @ID			delete-series-by-event
// @Tags      	event
// @Produce		json
// @Param		eventID	path	string	true	"Event ID"
// @Success		204	  {string}	    utilities.SuccessResponse
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/events/{eventID}/series/  [delete]
func (e *EventController) DeleteSeriesByEventID(c *fiber.Ctx) error {
	if err := e.eventService.DeleteSeriesByEventID(c.Params("eventID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// DeleteEvent godoc
//
// @Summary		Delete an event
// @Description	Deletes an event
// @ID			delete-event
// @Tags      	event
// @Produce		json
// @Param		eventID	path	string	true	"Event ID"
// @Success		204	  {string}	    utilities.SuccessResponse
// @Failure     400   {object}      error
// @Failure     401   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/events/{eventID}/  [delete]
func (e *EventController) DeleteEvent(c *fiber.Ctx) error {
	if err := e.eventService.DeleteEvent(c.Params("eventID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// GetHostsByEventID godoc
//
// @Summary		Retrieve all hosts by event
// @Description	Retrieves all hosts associated with an event
// @ID			get-hosts-by-event
// @Tags      	event
// @Produce		json
// @Param		eventID	path	string	true	"Event ID"
// @Success		200	  {object}	    []models.Club
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/events/{eventID}/hosts  [get]
func (e *EventController) GetHostsByEventID(c *fiber.Ctx) error {
	hosts, err := e.eventService.GetHostsByEventID(c.Params("eventID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(hosts)
}

// GetTagsByEventID godoc
//
// @Summary		Retrieve all tags by event
// @Description	Retrieves all tags associated with an event
// @ID			get-tags-by-event
// @Tags      	event
// @Produce		json
// @Param		eventID	path	string	true	"Event ID"
// @Success		200	  {object}	    []models.Tag
// @Failure     400   {object}      error
// @Failure     404   {object}      error
// @Failure     500   {object}      error
// @Router		/events/{eventID}/tags  [get]
func (e *EventController) GetTagsByEventID(c *fiber.Ctx) error {
	tags, err := e.eventService.GetTagsByEventID(c.Params("eventID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(tags)
}
