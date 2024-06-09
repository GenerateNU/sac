package auth

import (
	"slices"

	"github.com/GenerateNU/sac/backend/permission"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"

	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareHandler) Authorize(requiredPermissions ...permission.Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		strUser, err := sothic.GetFromSession("user", c)
		if err != nil {
			if err := sothic.StoreInSession("redirect", c.OriginalURL(), c); err != nil {
				return err
			}

			return c.Redirect("/api/v1/auth/microsoftonline")
		}

		user := models.UnmarshalUser(strUser)

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
