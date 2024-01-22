package utilities

import (
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
	"github.com/mcnijman/go-emailaddress"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

func InitValidators() {
	Validate.RegisterValidation("neu_email", ValidateEmail)
	Validate.RegisterValidation("password", ValidatePassword)
}

func ValidateEmail(fl validator.FieldLevel) bool {
	email, err := emailaddress.Parse(fl.Field().String())
	if err != nil {
		return false
	}

	if email.Domain != "northeastern.edu" {
		return false
	}

	return true
}

func ValidatePassword(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) < 8 {
		return false
	}
	specialCharactersMatch, _ := regexp.MatchString("[@#%&*+]", fl.Field().String())
	numbersMatch, _ := regexp.MatchString("[0-9]", fl.Field().String())
	return specialCharactersMatch && numbersMatch
}

// Validate the data sent to the server if the data is invalid, return an error otherwise, return nil
func ValidateData[T any](model T) error {
	if err := Validate.Struct(model); err != nil {
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
