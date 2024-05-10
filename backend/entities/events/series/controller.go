package series

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type EventSeriesController struct {
	eventSeriesService EventSeriesServiceInterface
}

func NewEventSeriesController(eventSeriesService EventSeriesServiceInterface) *EventSeriesController {
	return &EventSeriesController{eventSeriesService: eventSeriesService}
}

func (e *EventSeriesController) CreateEventSeries(c *fiber.Ctx) error {
	var body CreateEventSeriesRequestBody
	if err := c.BodyParser(&body); err != nil {
		return utilities.InvalidJSON()
	}

	events, err := e.eventSeriesService.CreateEventSeries(body)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(events)
}

func (e *EventSeriesController) GetEventSeries(c *fiber.Ctx) error {
	events, err := e.eventSeriesService.GetEventSeries(c.Params("seriesID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(events)
}

func (e *EventSeriesController) DeleteEventSeries(c *fiber.Ctx) error {
	if err := e.eventSeriesService.DeleteEventSeries(c.Params("seriesID")); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
