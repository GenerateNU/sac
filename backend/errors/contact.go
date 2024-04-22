package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToGetContacts = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get contacts",
	}
	FailedToGetContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get contact",
	}
	ContactNotFound = Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "contact not found",
	}
	FailedToPutContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to put contact",
	}
	FailedToDeleteContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete contact",
	}
	FailedToValidateContact = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate contact",
	}
)
