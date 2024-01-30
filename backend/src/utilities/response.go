package utilities

import "github.com/gofiber/fiber/v2"


func FiberMessage(c *fiber.Ctx, statusCode int, response string) error {
	return c.Status(statusCode).JSON(fiber.Map{"message": response})
}
