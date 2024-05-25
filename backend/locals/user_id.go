package locals

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserID(c *fiber.Ctx) (*uuid.UUID, error) {
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
