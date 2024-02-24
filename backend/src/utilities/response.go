package utilities

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// For swagger docs:
type SuccessResponse struct {
	Message string `json:"message"`
}

func FiberMessage(c *fiber.Ctx, statusCode int, response string) error {
	return c.Status(statusCode).JSON(fiber.Map{"message": response})
}

func FiberSuccess(c *fiber.Ctx, response string) error {
	return FiberMessage(c, fiber.StatusOK, response)
}

func FiberError(c *fiber.Ctx, statusCode int, response string) error {
	return FiberMessage(c, statusCode, response)
}

func Test(input string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Printf("Test: %s\n", input)
		fmt.Printf("Method: %s\n", c.Method())
		fmt.Printf("Path: %s\n", c.Path())
		fmt.Printf("Route Name: %s\n", c.Route().Name)
		fmt.Printf("Route Path: %s\n", c.Route().Path)
		fmt.Printf("Route Method: %s\n", c.Route().Method)
		return c.Next()
	}
}
