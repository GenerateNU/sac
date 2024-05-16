package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/integrations/calendar"
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
	calendar, err := cal.parseCalendarType(calendarType)
	if err != nil {
		return err
	}

	// Call the respective authorize method:
	authUrl := calendar.Authorize()

	return c.Redirect(authUrl, fiber.StatusFound)
}

func (cal *CalendarController) Token(c *fiber.Ctx) error {
	// Extract the calendar type from the query params:
	calendarType := c.Query("type")
	calendar, err := cal.parseCalendarType(calendarType)
	if err != nil {
		return err
	}
	
	// Extract the grant_type from the query params:
	grantType := c.Query("grant_type")

	// Extract the calendar type from the query params:
	calendarToken := c.Get("x-calendar-auth-code")

	// Call the respective token method:
	calendar.Token(grantType, &calendarToken)

	return nil
}

func (cal *CalendarController) parseCalendarType(calendarType string) (calendar.CalendarClientInterface, error) {
	switch calendarType {
		case "google":
			return cal.serviceSettings.Integrations.Calendar.Google, nil
		case "outlook":
			return cal.serviceSettings.Integrations.Calendar.Outlook, nil
		default:
			return nil, errors.New("invalid calendar type")
	}
}




