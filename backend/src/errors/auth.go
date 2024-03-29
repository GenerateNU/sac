package errors

import "github.com/gofiber/fiber/v2"

var (
	PassedAuthenticateMiddlewareButNilClaims = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "passed authenticate middleware but claims is nil",
	}
	FailedToCastToCustomClaims = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to cast to custom claims",
	}
	FailedToValidateUpdatePasswordBody = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "failed to validate update password body",
	}
)
