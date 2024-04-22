package types

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/email"
	"github.com/GenerateNU/sac/backend/file"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/search"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RouteParams struct {
	Router         fiber.Router
	AuthMiddleware *middleware.AuthMiddlewareService
	ServiceParams  ServiceParams
}

type ServiceParams struct {
	DB           *gorm.DB
	Validate     *validator.Validate
	AuthSettings *config.AuthSettings
	JWT          auth.JWTClientInterface
	Deps         Dependencies
}

type Dependencies struct {
	Search search.SearchClientInterface
	AI     search.AIClientInterface
	Email  email.EmailClientInterface
	File   file.FileClientInterface
}
