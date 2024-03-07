package middleware

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/gofiber/fiber/v2"
)

// // Authorizes admins of the specific club to make this request, skips check if super user
func (m *AuthMiddlewareService) UserAuthorizeById(c *fiber.Ctx) error {
	claims, err := auth.From(c)
	if err != nil {
		return errors.Unauthorized.FiberError(c)
	}

	if !m.IsSuper(c) {
		userID := c.Params("userID")
		if userID != claims.Issuer {
			return errors.Unauthorized.FiberError(c)
		}
	}

	return c.Next()
}