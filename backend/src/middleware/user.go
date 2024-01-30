package middleware

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

func (m *MiddlewareService) UserAuthorizeById(c *fiber.Ctx) error {
	idAsUint, err := utilities.ValidateID(c.Params("id"))
	if err != nil {
		return err
	}

	token, tokenErr := auth.ParseAccessToken(c.Cookies("access_token"))
	if tokenErr != nil {
		return err
	}

	
	claims, ok := token.Claims.(*types.CustomClaims)
	if !ok || !token.Valid {
		return errors.FailedToValidateAccessToken.FiberError(c)
	}
	
	issuerIDAsUint, err := utilities.ValidateID(claims.Issuer)
	if err != nil {
		return err
	}

	if issuerIDAsUint == idAsUint {
		return c.Next()
	}

	return errors.Unauthorized.FiberError(c)
}