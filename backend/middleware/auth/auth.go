package auth

import (
	"log/slog"
	"slices"

	auth "github.com/GenerateNU/sac/backend/entities/auth/base"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"
	"github.com/GenerateNU/sac/backend/permission"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"

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
	sothUser, err := sothic.CompleteUserAuth(c)
	if err == nil {
		user, err := auth.FindOrCreateUser(c.Context(), m.db, sothUser)
		if err != nil {
			return err
		}

		locals.SetUser(c, user)

		return c.Next()
	} else {
		return utilities.Unauthorized()
	}
}

func (m *AuthMiddlewareHandler) Authorize(requiredPermissions ...permission.Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !locals.UserExists(c) {
			if err := m.Authenticate(c); err != nil {
				slog.Error("authenticating user in authorize middleware", "error", err)
				return utilities.Unauthorized()
			}
		}

		user, err := locals.UserFrom(c)
		if err != nil {
			return err
		}

		if user.Role == models.Super {
			return c.Next()
		}

		userPermissions := permission.GetPermissions(user.Role)

		for _, requiredPermission := range requiredPermissions {
			if !slices.Contains(userPermissions, requiredPermission) {
				return utilities.Forbidden()
			}
		}

		return c.Next()
	}
}
