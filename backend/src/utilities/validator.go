package utilities

import (
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
	"github.com/mcnijman/go-emailaddress"
)

func RegisterCustomValidators(validate *validator.Validate) {
	validate.RegisterValidation("neu_email", validateEmail)
	validate.RegisterValidation("password", validatePassword)
	validate.RegisterValidation("mongo_url", validateMongoURL)
	validate.RegisterValidation("s3_url", validateS3URL)
	validate.RegisterValidation("id", validateID)
}

func validateEmail(fl validator.FieldLevel) bool {
	email, err := emailaddress.Parse(fl.Field().String())
	if err != nil {
		return false
	}

	if email.Domain != "northeastern.edu" {
		return false
	}

	return true
}

func validatePassword(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) < 8 {
		return false
	}
	specialCharactersMatch, _ := regexp.MatchString("[@#%&*+]", fl.Field().String())
	numbersMatch, _ := regexp.MatchString("[0-9]", fl.Field().String())
	return specialCharactersMatch && numbersMatch
}

func validateMongoURL(fl validator.FieldLevel) bool {
	return true
}

func validateS3URL(fl validator.FieldLevel) bool {
	return true
}

func validateID(fl validator.FieldLevel) bool {
	_, err := ValidateID(fl.Field().String())
	return err == nil
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

func ValidateNonNegative(value string) (*int, error) {
	valueAsInt, err := strconv.Atoi(value)

	errMsg := "failed to validate non-negative value"

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, errMsg)
	}

	if valueAsInt < 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, errMsg)
	}

	return &valueAsInt, nil
}
