package auth

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/database/store"
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
	db       *gorm.DB
	validate *validator.Validate
	session  store.Storer
}

func New(db *gorm.DB, validate *validator.Validate, session store.Storer) AuthMiddlewareService {
	return &AuthMiddlewareHandler{
		db:       db,
		validate: validate,
		session:  session,
	}
}
