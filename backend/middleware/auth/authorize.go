package auth

import (
	"net/http"
	"slices"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/permission"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareHandler) Authorize(requiredPermissions ...permission.Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		strUser, err := sothic.GetFromSession("microsoftonline", c)
		if err != nil {
			c.Set("redirect", "/api/v1/auth/login")
			return c.SendStatus(http.StatusUnauthorized)
		}

		user := models.UnmarshalUser(strUser)

		if user.Role == models.Super {
			locals.SetUser(c, user)
			return c.Next()
		}

		userPermissions := permission.GetPermissions(user.Role)

		for _, requiredPermission := range requiredPermissions {
			if !slices.Contains(userPermissions, requiredPermission) {
				return utilities.Forbidden()
			}
		}

		locals.SetUser(c, user)
		return c.Next()
	}
}
