package middleware

import (
	"backend/src/auth"
	"backend/src/models"
	"backend/src/types"

	"github.com/gofiber/fiber/v2"
)

const (
	AuthLoginPath = "/api/v1/users/auth/login"
	AuthRefreshPath = "/api/v1/users/auth/refresh"
	AuthLogoutPath = "/api/v1/users/auth/logout"
)

func Authenticate(c *fiber.Ctx) error {
	if c.Path() == AuthLoginPath || c.Path() == AuthRefreshPath || c.Path() == AuthLogoutPath {
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
func Authorize(requiredPermissions []models.Permission) fiber.Handler {
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