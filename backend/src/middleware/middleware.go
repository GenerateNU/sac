package middleware

import (
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MiddlewareInterface interface {
	ClubAuthorizeById(c *fiber.Ctx) error 
	UserAuthorizeById(c *fiber.Ctx) error
	Authenticate(c *fiber.Ctx) error
	Authorize(requiredPermissions ...models.Permission) func(c *fiber.Ctx) error
}

type MiddlewareService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewMiddlewareService(db *gorm.DB, validate *validator.Validate) *MiddlewareService {
	return &MiddlewareService{
		DB:       db,
		Validate: validate,
	}
}