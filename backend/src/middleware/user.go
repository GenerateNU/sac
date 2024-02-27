package middleware

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
)

// Authorizes admins of the specific club to make this request, skips check if super user
func (m *AuthMiddlewareService) UserAuthorizeById(c *fiber.Ctx) error {
	if m.IsSuper(c) {
		return c.Next()
	}

	idAsUUID, err := utilities.ValidateID(c.Params("userID"))
	if err != nil {
		return errors.FailedToValidateID.FiberError(c)
	}

	token, tokenErr := auth.ParseAccessToken(c.Cookies("access_token"), m.AuthSettings.AccessKey)
	if tokenErr != nil {
		return err
	}

	claims, ok := token.Claims.(*auth.CustomClaims)
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
