package middleware

import (
	"time"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthMiddlewareInterface interface {
	ClubAuthorizeById(c *fiber.Ctx) error
	UserAuthorizeById(c *fiber.Ctx) error
	Authenticate(c *fiber.Ctx) error
	Authorize(requiredPermissions ...auth.Permission) func(c *fiber.Ctx) error
	Skip(h fiber.Handler) fiber.Handler
	IsSuper(c *fiber.Ctx) bool
	Limiter(rate int, duration time.Duration) func(c *fiber.Ctx) error
}

type AuthMiddlewareService struct {
	DB           *gorm.DB
	Validate     *validator.Validate
	AuthSettings config.AuthSettings
}

func NewAuthAuthMiddlewareService(db *gorm.DB, validate *validator.Validate, authSettings config.AuthSettings) *AuthMiddlewareService {
	return &AuthMiddlewareService{
		DB:           db,
		Validate:     validate,
		AuthSettings: authSettings,
	}
}
