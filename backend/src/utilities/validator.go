package utilities

import (
	"regexp"
	"strconv"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/mcnijman/go-emailaddress"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterValidation("neu_email", ValidateEmail)
	validate.RegisterValidation("password", ValidatePassword)

	return validate
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

// Validates that an id follows postgres uint format, returns a uint otherwise returns an error
func ValidateID(id string) (*uint, *errors.Error) {
	idAsInt, err := strconv.Atoi(id)

	if err != nil || idAsInt < 1 { // postgres ids start at 1
		return nil, &errors.FailedToValidateID
	}

	idAsUint := uint(idAsInt)

	return &idAsUint, nil
}

func ValidateNonNegative(value string) (*int, *errors.Error) {
	valueAsInt, err := strconv.Atoi(value)

	if err != nil || valueAsInt < 0 {
		return nil, &errors.FailedToValidateNonNegativeValue
	}

	return &valueAsInt, nil
}
