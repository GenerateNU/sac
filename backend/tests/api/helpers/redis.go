package helpers

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/gofiber/fiber/v2"
)

type RedisMockClient struct{}

func NewRedisMockClient() *RedisMockClient {
	return &RedisMockClient{}
}

func (RedisMockClient) BlacklistToken(c *fiber.Ctx, tokenString string, tokenType auth.JWTType) *errors.Error {
	return nil
}

func (RedisMockClient) IsBlackListed(c *fiber.Ctx, tokenString string, tokenType auth.JWTType) (bool, *errors.Error) {
	return false, nil
}
