package utilities

import (
	"regexp"
	"strconv"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/google/uuid"

	"github.com/go-playground/validator/v10"
	"github.com/mcnijman/go-emailaddress"
)

func RegisterCustomValidators(validate *validator.Validate) {
	validate.RegisterValidation("neu_email", validateEmail)
	validate.RegisterValidation("password", validatePassword)
	validate.RegisterValidation("mongo_url", validateMongoURL)
	validate.RegisterValidation("s3_url", validateS3URL)
	validate.RegisterValidation("contact_pointer", func(fl validator.FieldLevel) bool {
		return validateContactPointer(validate, fl)
	})
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

func validateContactPointer(validate *validator.Validate, fl validator.FieldLevel) bool {
	contact, ok := fl.Parent().Interface().(models.Contact)

	if !ok {
		return false
	}

	switch contact.Type {
	case models.Email:
		return validate.Var(contact.Content, "email") == nil
	default:
		return validate.Var(contact.Content, "http_url") == nil
	}
}

func ValidateID(id string) (*uuid.UUID, *errors.Error) {
	idAsUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return &idAsUUID, nil
}

func ValidateNonNegative(value string) (*int, *errors.Error) {
	valueAsInt, err := strconv.Atoi(value)

	if err != nil || valueAsInt < 0 {
		return nil, &errors.FailedToValidateNonNegativeValue
	}

	return &valueAsInt, nil
}
