package auth

import (
	"net/url"
	"slices"
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/permission"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareHandler) Authorize(requiredPermissions ...permission.Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		strUser, err := sothic.GetFromSession("user", c)
		if err != nil {
			c.Cookie(&fiber.Cookie{
				Name:    "redirect",
				Value:   url.PathEscape(c.OriginalURL()),
				Expires: time.Now().Add(5 * time.Minute),
				// MARK: secure should be true in prod
				// use go build tags to do this
				HTTPOnly: true,
			})

			return c.Redirect("/api/v1/auth/login")
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
