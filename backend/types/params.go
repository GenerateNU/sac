package types

import (
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/GenerateNU/sac/backend/integrations"
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	utilityMiddleware "github.com/GenerateNU/sac/backend/middleware/utility"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RouteParams struct {
	Router            fiber.Router
	AuthMiddleware    authMiddleware.AuthMiddlewareService
	UtilityMiddleware utilityMiddleware.Service
	ServiceParams     ServiceParams
}

func NewRouteParams(router fiber.Router, authMiddleware authMiddleware.AuthMiddlewareService, utilityMiddleware utilityMiddleware.Service, serviceParams ServiceParams) RouteParams {
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
	Calendar     *config.CalendarSettings
	Search       *config.SearchSettings
	Stores       *store.Stores
	Integrations integrations.Integrations
}

func NewServiceParams(db *gorm.DB, validate *validator.Validate, calendar *config.CalendarSettings, stores *store.Stores, integrations integrations.Integrations) ServiceParams {
	return ServiceParams{
		DB:           db,
		Validate:     validate,
		Calendar:     calendar,
		Stores:       stores,
		Integrations: integrations,
	}
}
