package tags

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

type EventTagController struct {
	eventTagService EventTagServiceInterface
}

func NewEventTagController(eventTagService EventTagServiceInterface) *EventTagController {
	return &EventTagController{eventTagService: eventTagService}
}

func (e *EventTagController) CreateEventTags(c *fiber.Ctx) error {
	var eventTagsBody CreateEventTagsRequestBody
	if err := c.BodyParser(&eventTagsBody); err != nil {
		return utilities.InvalidJSON()
	}

	eventTags, err := e.eventTagService.CreateEventTags(c.Params("eventID"), eventTagsBody)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(eventTags)
}

func (e *EventTagController) GetEventTags(c *fiber.Ctx) error {
	tags, err := e.eventTagService.GetEventTags(c.Params("eventID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(tags)
}

func (e *EventTagController) DeleteEventTag(c *fiber.Ctx) error {
	err := e.eventTagService.DeleteEventTag(c.Params("eventID"), c.Params("tagID"))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
