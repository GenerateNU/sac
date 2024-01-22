package utilities

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse sends a standardized error response
func Error(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{"error": message})
}
