package middleware

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func getExcludedPaths() []map[string]string {
	return []map[string]string{
		{"/api/v1/users/": "POST"},
		{"/api/v1/auth/login": "POST"},
		{"/api/v1/auth/logout": "POST"},
		{"/api/v1/auth/refresh": "POST"},
		{"/api/v1/auth/send-code/": "POST"},
		{"/api/v1/auth/verify-email": "POST"},
		{"/api/v1/auth/verify-reset": "POST"},
		{"/api/v1/auth/forgot-password": "POST"},
	}
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


func splitPath(path string) []string {
	// Split the path by "/" separator
	parts := strings.Split(path, "/")
	
	// Remove empty elements caused by leading/trailing slashes
	nonEmptyParts := []string{}
	for _, part := range parts {
		if part != "" {
		nonEmptyParts = append(nonEmptyParts, part)
		}
	}
	return nonEmptyParts
}

func isPathParameterized(path string) bool {
	uuidRegex := regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")

	parts := splitPath(path)

	// if any part of the path is a uuid, return true
	for _, part := range parts {
		if uuidRegex.MatchString(part) {
			return true
		}
	}

	return false
}
func (m *AuthMiddlewareService) Authenticate(c *fiber.Ctx) error {
	for _, excludedPath := range getExcludedPaths() {
		for path, method := range excludedPath {
			if c.Path() == path && c.Method() == method {
				return c.Next()
			}
		}
	}

	if c.Method() == "OPTIONS" || (c.Method() == "GET" && !isPathParameterized(c.Path())) {
		return c.Next()
	}

	accessToken := GetAuthroizationToken(c)
	if accessToken == nil {
		return errors.Unauthorized.FiberError(c)
	}

	token, err := func() (*jwt.Token, error) {
		return jwt.ParseWithClaims(*accessToken, &auth.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.AuthSettings.AccessKey.Expose()), nil
		})
	}()
	if err != nil {
		return errors.Unauthorized.FiberError(c)
	}

	claims, ok := token.Claims.(*auth.CustomClaims)
	if !ok || !token.Valid {
		return errors.Unauthorized.FiberError(c)
	}

	// if auth.IsBlacklisted(*accessToken) {
	// 	return errors.Unauthorized.FiberError(c)
	// }

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

		userPermissions := auth.GetPermissions(models.UserRole(claims.Role))

		for _, requiredPermission := range requiredPermissions {
			if !slices.Contains(userPermissions, requiredPermission) {
				return errors.Forbidden.FiberError(c)
			}
		}

		return c.Next()
	}
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
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	})
}
