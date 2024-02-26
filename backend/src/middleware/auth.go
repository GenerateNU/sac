package middleware

import (
	"fmt"
	"slices"
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var paths = []string{
	"/api/v1/auth/login",
	"/api/v1/auth/refresh",
	"/api/v1/users/",
	"/api/v1/auth/logout",
	"/api/v1/auth/forgot-password",
	"/api/v1/auth/verify-reset",
	"/api/v1/auth/verify-email",
}

func (m *AuthMiddlewareService) IsSuper(c *fiber.Ctx) bool {
	claims, err := auth.From(c)
	if err != nil {
		_ = err.FiberError(c)
		return false
	}
	if claims == nil {
		return false
	}
	return claims.Role == string(models.Super)
}

func (m *AuthMiddlewareService) Authenticate(c *fiber.Ctx) error {
	if slices.Contains(paths, c.Path()) {
		return c.Next()
	}

	token, err := auth.ParseAccessToken(c.Cookies("access_token"), m.AuthSettings.AccessKey)
	if err != nil {
		return errors.Unauthorized.FiberError(c)
	}

	claims, ok := token.Claims.(*auth.CustomClaims)
	if !ok || !token.Valid {
		return errors.Unauthorized.FiberError(c)
	}

	if auth.IsBlacklisted(c.Cookies("access_token")) {
		return errors.Unauthorized.FiberError(c)
	}

	c.Locals("claims", claims)

	return c.Next()
}

func (m *AuthMiddlewareService) Authorize(requiredPermissions ...auth.Permission) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		claims, fromErr := auth.From(c)
		if fromErr != nil {
			return fromErr.FiberError(c)
		}

		if claims != nil && claims.Role == string(models.Super) {
			return c.Next()
		}

		role, err := auth.GetRoleFromToken(c.Cookies("access_token"), m.AuthSettings.AccessKey)
		if err != nil {
			return errors.Unauthorized.FiberError(c)
		}

		userPermissions := auth.GetPermissions(models.UserRole(*role))

		for _, requiredPermission := range requiredPermissions {
			if !slices.Contains(userPermissions, requiredPermission) {
				return errors.Unauthorized.FiberError(c)
			}
		}

		return c.Next()
	}
}

func (m *AuthMiddlewareService) Limiter(rate int, expiration time.Duration) func(c *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        rate,
		Expiration: expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s-%s", c.IP(), c.Path())
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	})
}
