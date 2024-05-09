package series

import (
	"github.com/gofiber/fiber/v2"
)

type EventSeriesController struct {
	eventSeriesService EventSeriesServiceInterface
}

func NewEventSeriesController(eventSeriesService EventSeriesServiceInterface) *EventSeriesController {
	return &EventSeriesController{eventSeriesService: eventSeriesService}
}

func (e *EventSeriesController) CreateEventSeries(c *fiber.Ctx) error {
	return nil
}

func (e *EventSeriesController) GetEventSeries(c *fiber.Ctx) error {
	return nil
}

func (e *EventSeriesController) UpdateEventSeries(c *fiber.Ctx) error {
	return nil
}

func (e *EventSeriesController) DeleteEventSeries(c *fiber.Ctx) error {
	return nil
}
