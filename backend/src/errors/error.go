package errors

import (
	"github.com/gofiber/fiber/v2"
)

type Error struct {
	StatusCode int
	Message    string
}

func (e *Error) FiberError(c *fiber.Ctx) error {
	return c.Status(e.StatusCode).JSON(fiber.Map{"error": e.Message})
}

func (e *Error) Error() string {
	return e.Message
}
