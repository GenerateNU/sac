package utilities

import "github.com/gofiber/fiber/v2"

// For swagger docs:
type SuccessResponse struct {
	Message string `json:message`
}

func FiberMessage(c *fiber.Ctx, statusCode int, response string) error {
	return c.Status(statusCode).JSON(fiber.Map{"message": response})
}
