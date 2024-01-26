package utilities

import (
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
	"github.com/mcnijman/go-emailaddress"
)

func RegisterCustomValidators(validate *validator.Validate) {
	validate.RegisterValidation("neu_email", ValidateEmail)
	validate.RegisterValidation("password", ValidatePassword)
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

func ValidateGenericEmail(email string) bool {
	_, err := emailaddress.Parse(email)
	return err == nil
}


func ValidatePassword(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) < 8 {
		return false
	}
	specialCharactersMatch, _ := regexp.MatchString("[@#%&*+]", fl.Field().String())
	numbersMatch, _ := regexp.MatchString("[0-9]", fl.Field().String())
	return specialCharactersMatch && numbersMatch
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
