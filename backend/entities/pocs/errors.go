package errors

import (
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/gofiber/fiber/v2"
)

var (
	FailedToUpsertPointOfContact = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update or insert point of contact",
	}
	FailedToGetAllPointOfContact = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get all point of contact",
	}
	FailedToGetClubPointOfContacts = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get club point of contacts",
	}
	FailedToUpdateClubPointOfContact = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update club point of contact",
	}
	FailedToGetClubPointOfContact = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get club point of contact",
	}
	FailedToDeleteClubPointOfContact = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete club point of contact",
	}
	PointOfContactNotFound = errors.Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "point of contact not found",
	}
	PointOfContactsNotFound = errors.Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "point of contact not found",
	}
	FailedToDeletePointOfContact = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete point of contact",
	}
	PointOfContactAlreadyExists = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "point of contact already exists",
	}
	FailedToValidatePointOfContact = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to validate point of contact",
	}
	FailedToCreatePointOfContact = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create point of contact",
	}
	FailedToValidateEmail = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate email",
	}
	FailedToMapResponseToModel = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to map response to model",
	}
	FailedToGetAPointOfContact = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get a point of contact",
	}
	FailedToGetPointOfContacts = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get point of contacts",
	}
	FailedToValidatePointOfContactId = errors.Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate point of contact id",
	}
	FailedToGetPointOfContact = errors.Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get point of contact",
	}
)
