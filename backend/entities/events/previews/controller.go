package previews

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/gofiber/fiber/v2"
)

type EventPreviewController struct {
	service Service
}

func NewEventController(service Service) *EventPreviewController {
	return &EventPreviewController{service: service}
}

func (e *EventPreviewController) GetEventPreviews(c *fiber.Ctx) error {
	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPageInfo
	}

	usePagination := c.QueryBool("pagination", true)
	if !usePagination {
		pageInfo = nil
	}

	events, err := e.service.GetEventPreviews(pageInfo, c.Query("start"), c.Query("end"))
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(events)
}
