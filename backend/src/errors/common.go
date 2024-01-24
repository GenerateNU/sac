package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToParseRequestBody = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to parse request body",
	}
	FailedToValidateID = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate id",
	}
	FailedToMapResposeToModel = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to map response to model",
	}
	InternalServerError = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "internal server error",
	}
)
