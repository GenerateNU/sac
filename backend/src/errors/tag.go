package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToValidateTag = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate tag",
	}
	FailedToCreateTag = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create tag",
	}
	FailedToGetTags = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get tags",
	}
	FailedToGetTag = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get tag",
	}
	FailedToUpdateTag = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update tag",
	}
	FailedToDeleteTag = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete tag",
	}
	TagNotFound = Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "tag not found",
	}
)
