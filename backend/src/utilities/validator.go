package utilities

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mcnijman/go-emailaddress"
)

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
	// TODO: we need to think of validation rules
	if len(fl.Field().String()) < 6 {
		return false
	}

	return true
}

// Validate the data sent to the server if the data is invalid, return an error otherwise, return nil
func ValidateData(model interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(model); err != nil {
		return err
	}

	return nil
}
