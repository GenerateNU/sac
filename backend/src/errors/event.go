package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToValidateEvent = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate event",
	}
	FailedToCreateEvent = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create event",
	}
	FailedToGetEvents = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get events",
	}
	FailedToGetEvent = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get event",
	}
	FailedToDeleteEvent = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete event",
	}
	FailedToUpdateEvent = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update event",
	}
	EventNotFound = Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "event not found",
	}
)
