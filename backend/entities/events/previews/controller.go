package previews

import (
	"net/http"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (e *Controller) GetEventPreviews(c *fiber.Ctx) error {
	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		return utilities.ErrExpectedPageInfo
	}

	usePagination := c.QueryBool("pagination", true)
	if !usePagination {
		pageInfo = nil
	}

	previews, err := e.service.GetEventPreviews(pageInfo, c.Query("start"), c.Query("end"))
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(previews)
}
