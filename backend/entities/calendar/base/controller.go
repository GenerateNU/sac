package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/gofiber/fiber/v2"
)

type CalendarController struct {
	serviceSettings types.ServiceParams
}

func NewCalendarController(serviceSettings types.ServiceParams) *CalendarController {
	return &CalendarController{serviceSettings: serviceSettings}
}

func (cal *CalendarController) Authorize(c *fiber.Ctx) error {
	// Extract the calendar type from the query params:
	calendarType := c.Query("type")
	if calendarType == "" {
	}

	return nil
}

func (cal * CalendarController) Token(c *fiber.Ctx) error {
	// Extract the token from the body:
	var tokenBody models.CalendarTokenRequest

	if err := c.BodyParser(&tokenBody); err != nil {
		return errors.FailedToParseRequestBody.FiberError(c)
	}

	return nil
}




