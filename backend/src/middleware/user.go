package middleware

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

func (m *MiddlewareService) UserAuthorizeById(c *fiber.Ctx) error {
	idAsUUID, err := utilities.ValidateID(c.Params("id"))
	if err != nil {
		return errors.FailedToValidateID.FiberError(c)
	}

	token, tokenErr := auth.ParseAccessToken(c.Cookies("access_token"), m.AuthSettings.AccessToken)
	if tokenErr != nil {
		return err
	}

	claims, ok := token.Claims.(*types.CustomClaims)
	if !ok || !token.Valid {
		return errors.FailedToValidateAccessToken.FiberError(c)
	}

	issuerIDAsUUID, err := utilities.ValidateID(claims.Issuer)
	if err != nil {
		return errors.FailedToValidateID.FiberError(c)
	}

	if issuerIDAsUUID.String() == idAsUUID.String() {
		return c.Next()
	}

	return errors.Unauthorized.FiberError(c)
}
