package types

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/integrations"
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	utilityMiddleware "github.com/GenerateNU/sac/backend/middleware/utility"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RouteParams struct {
	Router            fiber.Router
	AuthMiddleware    *authMiddleware.AuthMiddlewareService
	UtilityMiddleware *utilityMiddleware.UtilityMiddlewareService
	ServiceParams     ServiceParams
}

func NewRouteParams(router fiber.Router, authMiddleware *authMiddleware.AuthMiddlewareService, utilityMiddleware *utilityMiddleware.UtilityMiddlewareService, serviceParams ServiceParams) RouteParams {
	return RouteParams{
		Router:            router,
		AuthMiddleware:    authMiddleware,
		UtilityMiddleware: utilityMiddleware,
		ServiceParams:     serviceParams,
	}
}

type ServiceParams struct {
	DB           *gorm.DB
	Validate     *validator.Validate
	Auth         *config.AuthSettings
	JWT          auth.JWTClientInterface
	Calendar     *config.CalendarSettings
	Integrations integrations.Integrations
}

func NewServiceParams(db *gorm.DB, validate *validator.Validate, auth *config.AuthSettings, jwt auth.JWTClientInterface, calendar *config.CalendarSettings, integrations integrations.Integrations) ServiceParams {
	return ServiceParams{
		DB:           db,
		Validate:     validate,
		Auth:         auth,
		JWT:          jwt,
		Calendar:     calendar,
		Integrations: integrations,
	}
}
