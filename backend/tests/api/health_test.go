package tests

import (
	"testing"

	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
)

func TestHealthWorks(t *testing.T) {
	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   "/health",
	}.TestOnStatus(t, nil,
		fiber.StatusOK,
	).Close()
}
