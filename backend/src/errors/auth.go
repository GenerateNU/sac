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
	FailedToCreatePasswordReset = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to create password reset",
	}
	FailedToDeletePasswordReset = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete password reset",
	}
	TokenExpired = Error{
		StatusCode: fiber.StatusUnauthorized,
		Message:    "token expired",
	}
	FailedToGetPasswordResetToken = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get password reset token",
	}
	PasswordResetTokenNotFound = Error{
		StatusCode: fiber.StatusNotFound,
		Message:    "password reset token not found",
	}
	EmailAlreadyVerified = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "email already verified",
	}
	FailedToGenerateOTP = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to generate otp",
	}
	FailedToSaveOTP = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to save otp",
	}
	FailedToGetOTP = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get otp",
	}
	InvalidOTP = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "invalid otp",
	}
	OTPExpired = Error{
		StatusCode: fiber.StatusBadRequest,
		Message:    "otp expired",
	}
	FailedToUpdateEmailVerification = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to update email verification",
	}
	FailedToDeleteOTP = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to delete otp",
	}
	FailedToSendCode = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to send code",
	}
	
)
