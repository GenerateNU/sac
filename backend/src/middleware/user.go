package middleware

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareService) UserAuthorizeById(c *fiber.Ctx) error {
	if m.IsSuper(c) {
		return c.Next()
	}

	idAsUUID, err := utilities.ValidateID(c.Params("userID"))
	if err != nil {
		return errors.FailedToValidateID.FiberError(c)
	}

	claims, err := auth.From(c)
	if err != nil {
		return err.FiberError(c)
	}

	issuerIDAsUUID, err := utilities.ValidateID(claims.Issuer)
	if err != nil {
		return errors.FailedToValidateID.FiberError(c)
	}

	if issuerIDAsUUID.String() == idAsUUID.String() {
		return c.Next()
	}

	return errors.Forbidden.FiberError(c)
}
