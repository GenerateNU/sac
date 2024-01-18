package utilities

import (
	"github.com/go-playground/validator/v10"
)

// Validate the data sent to the server if the data is invalid, return an error otherwise, return nil
func ValidateData(model interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(model); err != nil {
		return err
	}

	return nil
}
