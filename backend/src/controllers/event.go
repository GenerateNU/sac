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

func (l *EventController) GetAllEvents(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	events, err := l.eventService.GetEvents(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (l *EventController) GetEvent(c *fiber.Ctx) error {
	event, err := l.eventService.GetEvent(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

func (l *EventController) GetClubEvents(c *fiber.Ctx) error {
	events, err := l.eventService.GetClubEvents(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (l *EventController) GetEventSeries(c *fiber.Ctx) error {
	events, err := l.eventService.GetEventSeries(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (l *EventController) CreateEvent(c *fiber.Ctx) error {

	var eventBody models.CreateEventRequestBody
	if err := c.BodyParser(&eventBody); err != nil {
		return errors.FailedToCreateEvent.FiberError(c)
	}

	event, err := l.eventService.CreateEvent(eventBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}

func (l *EventController) UpdateEvent(c *fiber.Ctx) error {
	var eventBody models.UpdateEventRequestBody

	if err := c.BodyParser(&eventBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedEvent, err := l.eventService.UpdateEvent(c.Params("id"), eventBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedEvent)
}

func (l *EventController) DeleteEventSeries(c *fiber.Ctx) error {

	if err := l.eventService.DeleteEventSeries(c.Params("id")); err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (l *EventController) DeleteEvent(c *fiber.Ctx) error {

	if err := l.eventService.DeleteEvent(c.Params("id")); err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
