package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToValidateAtLeastOneField = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate at least one field",
	}
	FailedToParseRequestBody = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to parse request body",
	}
	FailedtoParseQueryParams = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to parse query params",
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
	FailedToCreateRefreshToken = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create refresh token",
	}
	FailedToParseRefreshToken = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to parse refresh token",
	}
	FailedToParseAccessToken = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to parse access token",
	}
	FailedToValidateRefreshToken = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate refresh token",
	}
	FailedToValidateAccessToken = Error{
		StatusCode: fiber.StatusUnauthorized,
		Message:    "failed to validate access token",
	}
)
