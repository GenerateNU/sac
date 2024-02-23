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

func (e *EventController) GetAllEvents(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	events, err := e.eventService.GetEvents(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (e *EventController) GetEvent(c *fiber.Ctx) error {
	event, err := e.eventService.GetEvent(c.Params("eventID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

func (e *EventController) GetSeriesByEventID(c *fiber.Ctx) error {
	events, err := e.eventService.GetSeriesByEventID(c.Params("eventID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (e *EventController) GetSeriesByID(c *fiber.Ctx) error {
	events, err := e.eventService.GetSeriesByID(c.Params("seriesID"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

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

func (e *EventController) DeleteSeriesByID(c *fiber.Ctx) error {
	if err := e.eventService.DeleteSeriesByID(c.Params("seriesID")); err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (e *EventController) DeleteSeriesByEventID(c *fiber.Ctx) error {
	if err := e.eventService.DeleteSeriesByEventID(c.Params("eventID")); err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (l *EventController) DeleteEvent(c *fiber.Ctx) error {
	if err := l.eventService.DeleteEvent(c.Params("eventID")); err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
