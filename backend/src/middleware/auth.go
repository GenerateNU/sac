package middleware

import (
	"slices"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/types"

	"github.com/gofiber/fiber/v2"
)

var paths = []string{
	"/api/v1/auth/login",
	"/api/v1/auth/refresh",
	"/api/v1/users",
	"/api/v1/auth/logout",
}

func (m *MiddlewareService) Authenticate(c *fiber.Ctx) error {
	if slices.Contains(paths, c.Path()) {
		return c.Next()
	}

	token, err := auth.ParseAccessToken(c.Cookies("access_token"))
	if err != nil {
		return errors.FailedToParseAccessToken.FiberError(c)
	}

	_, ok := token.Claims.(*types.CustomClaims)
	if !ok || !token.Valid {
		return errors.FailedToValidateAccessToken.FiberError(c)
	}

	// check if token is blacklisted
	if auth.IsBlacklisted(c.Cookies("access_token")) {
		return errors.Unauthorized.FiberError(c)
	}

	return c.Next()
}

func (m *MiddlewareService) Authorize(requiredPermissions ...types.Permission) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		role, err := auth.GetRoleFromToken(c.Cookies("access_token"))
		if err != nil {
			return errors.FailedToParseAccessToken.FiberError(c)
		}

		userPermissions := types.GetPermissions(models.UserRole(*role))

		for _, requiredPermission := range requiredPermissions {
			if !containsPermission(userPermissions, requiredPermission) {
				return errors.Unauthorized.FiberError(c)
			}
		}

		return c.Next()
	}
}

func containsPermission(permissions []types.Permission, targetPermission types.Permission) bool {
	for _, permission := range permissions {
		if permission == targetPermission {
			return true
		}
	}
	return false
}
