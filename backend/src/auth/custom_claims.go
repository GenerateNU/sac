package auth

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
}

func From(c *fiber.Ctx) (*CustomClaims, *errors.Error) {
	rawClaims := c.Locals("claims")
	if rawClaims == nil {
		return nil, &errors.FailedToGetClaims
	}

	claims, ok := rawClaims.(*CustomClaims)
	if !ok {
		return nil, &errors.FailedToCastToCustomClaims
	}

	return claims, nil
}
