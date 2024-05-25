package auth

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type localsKey byte

const (
	claimsKey localsKey = 0
	userIDKey localsKey = 1
)

func CustomClaimsFrom(c *fiber.Ctx) (*CustomClaims, error) {
	rawClaims := c.Locals(claimsKey)
	if rawClaims == nil {
		return nil, utilities.Forbidden()
	}

	claims, ok := rawClaims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("claims are not of type CustomClaims. got: %T", rawClaims)
	}

	return claims, nil
}

func SetClaims(c *fiber.Ctx, claims *CustomClaims) {
	c.Locals(claimsKey, claims)
}

func UserIDFrom(c *fiber.Ctx) (*uuid.UUID, error) {
	userID := c.Locals(userIDKey)
	if userID == nil {
		return nil, utilities.Forbidden()
	}

	id, ok := userID.(*uuid.UUID)
	if !ok {
		return nil, fmt.Errorf("userID is not of type uuid.UUID. got: %T", userID)
	}

	return id, nil
}

func SetUserID(c *fiber.Ctx, id *uuid.UUID) {
	c.Locals(userIDKey, id)
}
