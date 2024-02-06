package middleware

import (
	"slices"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
)

var paths = []string{
	"/api/v1/auth/login",
	"/api/v1/auth/refresh",
	"/api/v1/users/",
	"/api/v1/auth/logout",
}

func SuperSkipper(h fiber.Handler) fiber.Handler {
	return skip.New(h, func(c *fiber.Ctx) bool {
		claims, err := types.From(c)
		if err != nil {
			_ = err.FiberError(c)
			return false
		}
		if claims == nil {
			return false
		}
		return claims.Role == string(models.Super)
	})
}

func (m *MiddlewareService) Authenticate(c *fiber.Ctx) error {
	if slices.Contains(paths, c.Path()) {
		return c.Next()
	}

	token, err := auth.ParseAccessToken(c.Cookies("access_token"), m.AuthSettings.AccessToken)
	if err != nil {
		return errors.FailedToParseAccessToken.FiberError(c)
	}

	claims, ok := token.Claims.(*types.CustomClaims)
	if !ok || !token.Valid {
		return errors.FailedToValidateAccessToken.FiberError(c)
	}

	if auth.IsBlacklisted(c.Cookies("access_token")) {
		return errors.Unauthorized.FiberError(c)
	}

	c.Locals("claims", claims)

	return c.Next()
}

func (m *MiddlewareService) Authorize(requiredPermissions ...types.Permission) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		claims, fromErr := types.From(c)
		if fromErr != nil {
			return fromErr.FiberError(c)
		}

		if claims != nil && claims.Role == string(models.Super) {
			return c.Next()
		}

		role, err := auth.GetRoleFromToken(c.Cookies("access_token"), m.AuthSettings.AccessToken)
		if err != nil {
			return errors.FailedToParseAccessToken.FiberError(c)
		}

		userPermissions := types.GetPermissions(models.UserRole(*role))

		for _, requiredPermission := range requiredPermissions {
			if !slices.Contains(userPermissions, requiredPermission) {
				return errors.Unauthorized.FiberError(c)
			}
		}

		return c.Next()
	}
}
