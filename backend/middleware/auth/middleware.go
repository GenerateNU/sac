package auth

import (
	"github.com/GenerateNU/sac/backend/permission"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthMiddlewareService interface {
	ClubAuthorizeById(c *fiber.Ctx, extractor ExtractID) error
	UserAuthorizeById(c *fiber.Ctx) error
	EventAuthorizeById(c *fiber.Ctx, extractor ExtractID) error
	Authorize(requiredPermissions ...permission.Permission) fiber.Handler
}

type AuthMiddlewareHandler struct {
	db       *gorm.DB
	validate *validator.Validate
}

func New(db *gorm.DB, validate *validator.Validate) AuthMiddlewareService {
	return &AuthMiddlewareHandler{
		db:       db,
		validate: validate,
	}
}
