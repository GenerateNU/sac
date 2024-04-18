package errors

import "github.com/gofiber/fiber/v2"

var (
	InvalidToken = Error{
		fiber.StatusBadRequest,
		"invalid token",
	}
	FailedToGetExp = Error{
		fiber.StatusInternalServerError,
		"Failed to get expiry",
	}
)
