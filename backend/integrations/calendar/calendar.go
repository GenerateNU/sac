package calendar

import (
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/gofiber/fiber/v2"
)

type CalendarClientInterface interface {
	Authorize() *string
	Token(c *fiber.Ctx) *errors.Error
}
