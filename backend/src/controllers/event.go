package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type EventController struct {
	eventService services.EventServiceInterface
}

func NewEventController(eventService services.EventServiceInterface) *EventController {
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
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/event/  [get]
func (e *EventController) GetAllEvents(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	events, err := e.eventService.GetEvents(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(events)
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
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/event/{eventID}  [get]
func (e *EventController) GetEvent(c *fiber.Ctx) error {
	event, err := e.eventService.GetEvent(c.Params("eventID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(event)
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
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/event/{eventID}/series  [get]
func (e *EventController) GetSeriesByEventID(c *fiber.Ctx) error {
	events, err := e.eventService.GetSeriesByEventID(c.Params("eventID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

// GetSeriesByID godoc
//
// @Summary		Retrieve a series by ID
// @Description	Retrieves a series by ID
// @ID			get-series-by-id
// @Tags      	event
// @Produce		json
// @Param		seriesID	path	string	true	"Series ID"
// @Success		200	  {object}	    models.Series
// @Failure     400   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/series/{seriesID}  [get]
func (e *EventController) GetSeriesByID(c *fiber.Ctx) error {
	events, err := e.eventService.GetSeriesByID(c.Params("seriesID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(events)
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
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/event/  [post]
func (e *EventController) CreateEvent(c *fiber.Ctx) error {
	var eventBody models.CreateEventRequestBody
	if err := c.BodyParser(&eventBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	event, err := e.eventService.CreateEvent(eventBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}

// CreateSeries godoc
//
// @Summary		Create a series
// @Description	Creates a series
// @ID			create-series
// @Tags      	event
// @Accept		json
// @Produce		json
// @Param		eventID	path	string	true	"Event ID"
// @Param		seriesBody	body	models.CreateSeriesRequestBody	true	"Series Body"
// @Success		201	  {object}	  models.Series
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/event/{eventID}/series  [post]
func (e *EventController) UpdateEvent(c *fiber.Ctx) error {
	var eventBody models.UpdateEventRequestBody
	if err := c.BodyParser(&eventBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedEvent, err := e.eventService.UpdateEvent(c.Params("eventID"), eventBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedEvent)
}

// UpdateSeriesByID godoc
//
// @Summary		Update a series by ID
// @Description	Updates a series by ID
// @ID			update-series-by-id
// @Tags      	event
// @Accept		json
// @Produce		json
// @Param		eventID	path	string	true	"Event ID"
// @Param		seriesID	path	string	true	"Series ID"
// @Param		seriesBody	body	models.UpdateSeriesRequestBody	true	"Series Body"
// @Success		200	  {object}	  models.Series
// @Failure     400   {object}    errors.Error
// @Failure     401   {object}    errors.Error
// @Failure     404   {object}    errors.Error
// @Failure     500   {object}    errors.Error
// @Router		/event/{eventID}/series/{seriesID}  [put]
func (e *EventController) UpdateSeriesByID(c *fiber.Ctx) error {
	var seriesBody models.UpdateSeriesRequestBody
	if err := c.BodyParser(&seriesBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedSeries, err := e.eventService.UpdateSeries(c.Params("seriesID"), seriesBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedSeries)
}

func (e *EventController) UpdateSeriesByEventID(c *fiber.Ctx) error {
	var seriesBody models.UpdateSeriesRequestBody
	if err := c.BodyParser(&seriesBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedSeries, err := e.eventService.UpdateSeriesByEventID(c.Params("eventID"), seriesBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedSeries)
}

// DeleteSeriesByID godoc
//
// @Summary		Delete a series by ID
// @Description	Deletes a series by ID
// @ID			delete-series-by-id
// @Tags      	event
// @Produce		json
// @Param		seriesID	path	string	true	"Series ID"
// @Success		204	  {string}	    utilities.SuccessResponse
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/series/{seriesID}  [delete]
func (e *EventController) DeleteSeriesByID(c *fiber.Ctx) error {
	if err := e.eventService.DeleteSeriesByID(c.Params("seriesID")); err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
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
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/event/{eventID}/series  [delete]
func (e *EventController) DeleteSeriesByEventID(c *fiber.Ctx) error {
	if err := e.eventService.DeleteSeriesByEventID(c.Params("eventID")); err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
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
// @Failure     400   {object}      errors.Error
// @Failure     401   {object}      errors.Error
// @Failure     404   {object}      errors.Error
// @Failure     500   {object}      errors.Error
// @Router		/event/{eventID}  [delete]
func (l *EventController) DeleteEvent(c *fiber.Ctx) error {
	if err := l.eventService.DeleteEvent(c.Params("eventID")); err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
