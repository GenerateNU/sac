package auth

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthMiddlewareService interface {
	ClubAuthorizeById(c *fiber.Ctx, extractor ExtractID) error
	UserAuthorizeById(c *fiber.Ctx) error
	EventAuthorizeById(c *fiber.Ctx, extractor ExtractID) error
	Authenticate(c *fiber.Ctx) error
	Authorize(requiredPermissions ...auth.Permission) fiber.Handler
	IsSuper(c *fiber.Ctx) bool
}

type AuthMiddlewareHandler struct {
	DB       *gorm.DB
	Validate *validator.Validate
	Auth     config.AuthSettings
	Stores   *store.Stores
	OAuth    oauth.Service
}

func New(db *gorm.DB, validate *validator.Validate, authSettings config.AuthSettings, stores *store.Stores, oauth oauth.Service) AuthMiddlewareService {
	return &AuthMiddlewareHandler{
		DB:       db,
		Validate: validate,
		Auth:     authSettings,
		Stores:   stores,
		OAuth:    oauth,
	}
}
