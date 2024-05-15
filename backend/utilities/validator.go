package utilities

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/google/uuid"
	"github.com/mcnijman/go-emailaddress"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators() (*validator.Validate, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.RegisterValidation("neu_email", validateEmail); err != nil {
		return nil, err
	}

	if err := validate.RegisterValidation("s3_url", validateS3URL); err != nil {
		return nil, err
	}

	if err := validate.RegisterValidation("contact_pointer", func(fl validator.FieldLevel) bool {
		return validateContactPointer(validate, fl)
	}); err != nil {
		return nil, err
	}

	if err := validate.RegisterValidation("not_equal_if_not_empty", validateNotEqualIfNotEmpty); err != nil {
		return nil, err
	}

	return validate, nil
}

func AtLeastOne[Model any](body Model, model Model) bool {
	return reflect.DeepEqual(body, model)
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

func validateS3URL(fl validator.FieldLevel) bool {
	return strings.HasPrefix(fl.Field().String(), "https://s3.amazonaws.com/")
}

func validateContactPointer(validate *validator.Validate, fl validator.FieldLevel) bool {
	contact, ok := fl.Parent().Interface().(models.PutContactRequestBody)
	if !ok {
		return false
	}

	validationRules := map[models.ContactType]string{
		models.Facebook:   "http_url",
		models.Instagram:  "http_url",
		models.X:          "http_url",
		models.LinkedIn:   "http_url",
		models.YouTube:    "http_url",
		models.GitHub:     "http_url",
		models.Slack:      "http_url",
		models.Discord:    "http_url",
		models.Email:      "email",
		models.CustomSite: "http_url",
	}

	rule, ok := validationRules[contact.Type]
	if !ok {
		return false // invalid contact type
	}

	return validate.Var(contact.Content, rule) == nil && strings.HasPrefix(contact.Content, models.GetContentPrefix(contact.Type))
}

func validateNotEqualIfNotEmpty(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	otherField := fl.Parent().FieldByName(fl.Param()).String()

	return field == "" || field != otherField
}

func ValidateID(id string) (*uuid.UUID, error) {
	idAsUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &idAsUUID, nil
}

func ValidateNonNegative(value string) (*int, error) {
	valueAsInt, err := strconv.Atoi(value)

	if err != nil || valueAsInt < 0 {
		return nil, fmt.Errorf("value must be a non-negative integer. got %s", value)
	}

	return &valueAsInt, nil
}

func Validate(validate *validator.Validate, s any) error {
	if err := validate.Struct(s); err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if ok {
			errors := make(map[string]string, len(validationErrs))

			for _, validationErr := range validationErrs {
				errors[validationErr.Field()] = fmt.Sprintf("%s failed %s", validationErr.Field(), validationErr.Tag())
			}

			return InvalidRequestData(errors)
		} else {
			return fmt.Errorf("failed to validate struct: %w", err)
		}
	}

	return nil
}
