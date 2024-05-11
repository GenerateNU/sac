package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func (m *AuthMiddlewareService) IsSuper(c *fiber.Ctx) bool {
	claims, err := auth.From(c)
	if err != nil {
		_ = err
		return false
	}
	if claims == nil {
		return false
	}
	return claims.Role == string(models.Super)
}

func GetAuthroizationToken(c *fiber.Ctx) *string {
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
	accessToken := GetAuthroizationToken(c)
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

	c.Locals("claims", claims)

	return nil
}

func (m *AuthMiddlewareService) Authorize(requiredPermissions ...auth.Permission) func(c *fiber.Ctx) error {
	return utilities.Make(func(c *fiber.Ctx) error {
		authErr := m.Authenticate(c)
		if authErr != nil {
			return utilities.Unauthorized()
		}

		claims, err := auth.From(c)
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
	})
}

// TODO: implement rate limiting with redis
func (m *AuthMiddlewareService) Limiter(rate int, expiration time.Duration) func(c *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        rate,
		Expiration: expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s-%s", c.IP(), c.Path())
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	})
}
