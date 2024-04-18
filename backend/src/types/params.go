package types

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/email"
	"github.com/GenerateNU/sac/backend/src/file"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/redis"
	"github.com/GenerateNU/sac/backend/src/search"
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
	Deps         Dependencies
}

type Dependencies struct {
	Redis  redis.RedisClientInterface
	JWT    auth.JWTClientInterface
	Search search.SearchClientInterface
	AI     search.AIClientInterface
	Email  email.EmailClientInterface
	File   file.FileClientInterface
}
