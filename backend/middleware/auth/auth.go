package auth

import (
	"slices"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareHandler) IsSuper(c *fiber.Ctx) bool {
	user, err := locals.UserFrom(c)
	if err != nil {
		return false
	}
	if user == nil {
		return false
	}
	return user.Role == models.Super
}

func (m *AuthMiddlewareHandler) Authenticate(c *fiber.Ctx) error {
	// TODO
	return c.Next()
}

func (m *AuthMiddlewareHandler) Authorize(requiredPermissions ...auth.Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authErr := m.Authenticate(c)
		if authErr != nil {
			return utilities.Unauthorized()
		}

		user, err := locals.UserFrom(c)
		if err != nil {
			return err
		}

		if user == nil {
			return utilities.Unauthorized()
		}

		if user.Role == models.Super {
			return c.Next()
		}

		userPermissions := auth.GetPermissions(user.Role)

		for _, requiredPermission := range requiredPermissions {
			if !slices.Contains(userPermissions, requiredPermission) {
				return utilities.Forbidden()
			}
		}

		return c.Next()
	}
}
