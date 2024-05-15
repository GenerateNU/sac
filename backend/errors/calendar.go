package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToExchangeAuthCode = Error{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "failed to exchange auth code for tokens",
	}
)
