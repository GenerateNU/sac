package utilities

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Validate the data sent to the server if the data is invalid, return an error otherwise, return nil
func ValidateData(model interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(model); err != nil {
		return err
	}

	return nil
}

// Validates that an id follows postgres uint format, returns a uint otherwise returns an error
func ValidateID(id string) (*uint, error) {
	idAsInt, err := strconv.Atoi(id)

	errMsg := "failed to validate id"

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, errMsg)
	}

	if idAsInt < 1 { // postgres ids start at 1
		return nil, fiber.NewError(fiber.StatusBadRequest, errMsg)
	}

	idAsUint := uint(idAsInt)

	return &idAsUint, nil
}
