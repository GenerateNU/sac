package middleware

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/types"

	"github.com/gofiber/fiber/v2"
)

const (
	AuthLoginPath    = "/api/v1/users/auth/login"
	AuthRefreshPath  = "/api/v1/users/auth/refresh"
	AuthRegisterPath = "/api/v1/users/auth/register"
)

func (m *MiddlewareService) Authenticate(c *fiber.Ctx) error {
	// TODO: use a contains function instead of this
	if c.Path() == AuthLoginPath || c.Path() == AuthRefreshPath || c.Path() == AuthRegisterPath {
		return c.Next()
	}

	token, err := auth.ParseAccessToken(c.Cookies("access_token"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	_, ok := token.Claims.(*types.CustomClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	return c.Next()
}

// Authorize is a middleware that checks if the user has the required permissions
func (m *MiddlewareService) Authorize(requiredPermissions ...models.Permission) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get user role from the token
		role, err := auth.GetRoleFromToken(c.Cookies("access_token"))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}

		// Get permissions for the user's role
		userPermissions := models.GetPermissions(models.UserRole(*role))

		// Check if the user has the required permissions
		for _, requiredPermission := range requiredPermissions {
			if !containsPermission(userPermissions, requiredPermission) {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"message": "forbidden",
				})
			}
		}

		// User has the required permissions, continue to the next middleware/handler
		return c.Next()
	}
}

// containsPermission checks if a permission is present in a slice of permissions
func containsPermission(permissions []models.Permission, targetPermission models.Permission) bool {
	for _, permission := range permissions {
		if permission == targetPermission {
			return true
		}
	}
	return false
}
