package auth

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthMiddlewareInterface interface {
	ClubAuthorizeById(c *fiber.Ctx) error
	UserAuthorizeById(c *fiber.Ctx) error
	Authenticate(c *fiber.Ctx) error
	Authorize(requiredPermissions ...auth.Permission) fiber.Handler
	IsSuper(c *fiber.Ctx) bool
}

type AuthMiddlewareService struct {
	DB       *gorm.DB
	Validate *validator.Validate
	Auth     config.AuthSettings
	Stores   *store.Stores
}

func New(db *gorm.DB, validate *validator.Validate, authSettings config.AuthSettings, stores *store.Stores) *AuthMiddlewareService {
	return &AuthMiddlewareService{
		DB:       db,
		Validate: validate,
		Auth:     authSettings,
		Stores:   stores,
	}
}
