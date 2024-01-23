package tests

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHealthWorks(t *testing.T) {
	TestRequest{
		Method: fiber.MethodGet,
		Path:   "/health",
	}.TestOnStatus(t, nil,
		200,
	).Close()
}
