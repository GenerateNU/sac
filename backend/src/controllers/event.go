package controllers

import (
	"strconv"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

type EventController struct {
	clubService services.EventServiceInterface
}

func NewEventController(clubService services.EventServiceInterface) *EventController {
	return &EventController{clubService: clubService}
}

func (l *EventController) GetAllEvents(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	clubs, err := l.clubService.GetEvents(c.Query("limit", strconv.Itoa(defaultLimit)), c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(clubs)
}

func (l *EventController) CreateEvent(c *fiber.Ctx) error {
	var clubBody models.CreateEventRequestBody
	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	club, err := l.clubService.CreateEvent(clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusCreated).JSON(club)
}

func (l *EventController) GetEvent(c *fiber.Ctx) error {
	club, err := l.clubService.GetEvent(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(club)
}

func (l *EventController) UpdateEvent(c *fiber.Ctx) error {
	var clubBody models.UpdateEventRequestBody

	if err := c.BodyParser(&clubBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	updatedEvent, err := l.clubService.UpdateEvent(c.Params("id"), clubBody)
	if err != nil {
		return err.FiberError(c)
	}

	return c.Status(fiber.StatusOK).JSON(updatedEvent)
}

func (l *EventController) DeleteEvent(c *fiber.Ctx) error {
	err := l.clubService.DeleteEvent(c.Params("id"))
	if err != nil {
		return err.FiberError(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
