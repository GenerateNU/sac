package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToUpsertPointOfContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update or insert point of contact",
	}
	FailedToGetAllPointOfContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get all point of contact",
	}
	PointOfContactNotFound = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "point of contact not found",
	}
	FailedToDeletePointOfContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete point of contact",
	}
	FailedToValidatePointOfContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to validate point of contact",
	}
	FailedToValidateEmail = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate email",
	}
	FailedToMapResponseToModel = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to map response to model",
	}
	FailedToGetAPointOfContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get a point of contact",
	}
	FailedToGetPointOfContacts = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get point of contacts",
	}
	FailedToValidatePointOfContactId = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate point of contact id",
	}
	FailedToGetPointOfContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get point of contact",
	}
)
