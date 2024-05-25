package auth

import (
	"fmt"
	"slices"
	"strings"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/locals"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
)

func (m *AuthMiddlewareService) IsSuper(c *fiber.Ctx) bool {
	claims, err := locals.CustomClaimsFrom(c)
	if err != nil {
		_ = err
		return false
	}
	if claims == nil {
		return false
	}
	return claims.Role == string(models.Super)
}

func GetAuthorizationToken(c *fiber.Ctx) *string {
	accessToken := c.Get("Authorization")
	if accessToken == "" {
		return nil
	}

	token := strings.Split(accessToken, "Bearer ")
	if len(token) != 2 {
		return nil
	}

	return &token[1]
}

func (m *AuthMiddlewareService) Authenticate(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken := GetAuthorizationToken(c)
		if accessToken == nil {
			return utilities.Unauthorized()
		}

		token, err := func() (*jwt.Token, error) {
			return jwt.ParseWithClaims(*accessToken, &auth.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(m.Auth.AccessKey.Expose()), nil
			})
		}()
		if err != nil {
			return utilities.Unauthorized()
		}

		claims, ok := token.Claims.(*auth.CustomClaims)
		if !ok || !token.Valid {
			return utilities.Unauthorized()
		}

		// if auth.IsBlacklisted(*accessToken) {
		// 	return errors.Unauthorized.FiberError(c)
		// }

		locals.SetCustomClaims(c, claims)

		rawUserID := claims.Issuer
		userID, err := uuid.Parse(rawUserID)
		if err != nil {
			return fmt.Errorf("invalid user id: %s", rawUserID)
		}

		locals.SetUserID(c, &userID)

		return nil
	}(c)
}

func (m *AuthMiddlewareService) Authorize(requiredPermissions ...auth.Permission) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authErr := m.Authenticate(c)
		if authErr != nil {
			return utilities.Unauthorized()
		}

		claims, err := locals.CustomClaimsFrom(c)
		if err != nil {
			return err
		}

		if claims != nil && claims.Role == string(models.Super) {
			return c.Next()
		}

		userPermissions := auth.GetPermissions(models.UserRole(claims.Role))

		for _, requiredPermission := range requiredPermissions {
			if !slices.Contains(userPermissions, requiredPermission) {
				return utilities.Forbidden()
			}
		}

		return c.Next()
	}
}
