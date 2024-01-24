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
	FailedToValidateLimit = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate limit",
	}
	FailedToValidatePage = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate page",
	}
)