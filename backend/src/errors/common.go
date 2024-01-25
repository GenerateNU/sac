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
	FailedToValidateNonNegativeValue = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate non-negative value",
	}
	FailedToMapRequestToModel = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to map request to model",
	}
	InternalServerError = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "internal server error",
	}
	FailedToValidateLimit = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate limit",
	}
	FailedToValidatePage = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate page",
	}
	Unauthorized = Error{
		StatusCode: fiber.StatusUnauthorized,
		Message:    "unauthorized",
	}
	FailedToSignToken = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to sign token",
	}
	FailedToCreateAccessToken = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create access token",
	}
	FailedToParseRefreshToken = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to parse refresh token",
	}
	FailedToParseAccessToken = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to parse access token",
	}
)
