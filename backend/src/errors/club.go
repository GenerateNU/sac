package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToValidateUserID = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate user id",
	}
	FailedToValidateClub = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate club",
	}
	FailedToCreateClub = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create club",
	}
	FailedToGetClubs = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get clubs",
	}
	FailedToGetClub = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get club",
	}
	FailedToDeleteClub = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete club",
	}
	FailedToUpdateClub = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update club",
	}
	ClubNotFound = Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "club not found",
	}
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
	FailedToCreateContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create contact",
	}
	FailedToUpdateContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update contact",
	}
	FailedToDeleteContact = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete contact",
	}
	FailedToValidateContact = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate contact",
	FailedtoGetAdminIDs = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get admin ids",
	}
)
