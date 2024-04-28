package middleware

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareService) UserAuthorizeById(c *fiber.Ctx) error {
	authErr := m.Authenticate(c)
	if authErr != nil {
		return errors.Unauthorized.FiberError(c)
	}

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
