package locals

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

func CustomClaims(c *fiber.Ctx) (*auth.CustomClaims, error) {
	rawClaims := c.Locals(claimsKey)
	if rawClaims == nil {
		return nil, utilities.Forbidden()
	}

	claims, ok := rawClaims.(*auth.CustomClaims)
	if !ok {
		return nil, fmt.Errorf("claims are not of type auth.CustomClaims. got: %T", rawClaims)
	}

	return claims, nil
}

func SetCustomClaims(c *fiber.Ctx, claims *auth.CustomClaims) {
	c.Locals(claimsKey, claims)
}
