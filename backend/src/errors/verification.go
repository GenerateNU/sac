package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToGetTemplate = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to get template",
	}
	FailedToSendEmail = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to send email",
	}
)