package calendar

import (
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/gofiber/fiber/v2"
)

type GoogleClient struct {
	Settings config.GoogleClientSettings
}

func (client GoogleClient) Authorize() *string { return nil }

func (client GoogleClient) Token(c *fiber.Ctx) *errors.Error { return nil }