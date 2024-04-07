package helpers

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/gofiber/fiber/v2"
)

type RedisMockClient struct{}

func NewRedisMockClient() *RedisMockClient {
	return &RedisMockClient{}
}

func (RedisMockClient) BlacklistToken(c *fiber.Ctx, userId string, token string, expiration time.Duration) *errors.Error {
	return nil
}
func (RedisMockClient) IsBlackListed(c *fiber.Ctx, userId string, token string) (bool, *errors.Error) {
	return false, nil
}
