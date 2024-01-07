package utilities

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Validate the data sent to the server if the data is invalid, return an error otherwise, return nil
func ValidateData(c *fiber.Ctx, model interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(model); err != nil {
		return err
	}

	return nil
}
