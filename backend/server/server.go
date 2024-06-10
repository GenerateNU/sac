package server

import (
	"fmt"

	"github.com/garrettladley/fiberpaginate"
	go_json "github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/database/store"

	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/goog"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/msft"

	auth "github.com/GenerateNU/sac/backend/entities/auth/base"
	categories "github.com/GenerateNU/sac/backend/entities/categories/base"
	clubs "github.com/GenerateNU/sac/backend/entities/clubs/base"
	events "github.com/GenerateNU/sac/backend/entities/events/base"
	files "github.com/GenerateNU/sac/backend/entities/files/base"
	leader "github.com/GenerateNU/sac/backend/entities/leadership/base"
	socials "github.com/GenerateNU/sac/backend/entities/socials/base"
	tags "github.com/GenerateNU/sac/backend/entities/tags/base"
	users "github.com/GenerateNU/sac/backend/entities/users/base"
	"github.com/GenerateNU/sac/backend/integrations"
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	utilityMiddleware "github.com/GenerateNU/sac/backend/middleware/utility"
	search "github.com/GenerateNU/sac/backend/search/base"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"gorm.io/gorm"
)

// @title SAC API
// @version 1.0
// @description Backend Server for SAC App

// @contact.name	David Oduneye and Garrett Ladley
// @contact.email	oduneye.d@northeastern.edu and ladley.g@northeastern.edu
// @host 127.0.0.1:8080
// @BasePath /
// @schemes http
func Init(db *gorm.DB, stores *store.Stores, integrations integrations.Integrations, settings config.Settings) *fiber.App {
	app := newFiberApp(settings.Application)

	validate, err := utilities.RegisterCustomValidators()
	if err != nil {
		panic(fmt.Sprintf("Error registering custom validators: %s", err))
	}

	applicationURL := settings.Application.ApplicationURL()

	msftProvider := msft.New(settings.Microsft.Key, settings.Microsft.Secret, fmt.Sprintf("%s/api/v1/auth/microsoftonline/callback", applicationURL), settings.Microsft.Tenant)
	googProvider := goog.New(settings.Google.Key, settings.Google.Secret, fmt.Sprintf("%s/api/v1/auth/google/callback", applicationURL))

	authMiddleware := authMiddleware.New(
		db,
		validate,
		msftProvider,
	)

	utilityMiddleware := utilityMiddleware.New(fiberpaginate.New(), stores.Limiter)

	apiv1 := app.Group("/api/v1")

	routeParams := types.NewRouteParams(
		apiv1,
		authMiddleware,
		utilityMiddleware,
		types.NewServiceParams(
			db,
			validate,
			&settings.Calendar,
			stores,
			integrations,
		),
	)

	allRoutes(app, routeParams,
		auth.NewParams(
			msftProvider,
			applicationURL,
			apiv1,
			db,
			integrations.Email,
			validate,
			googProvider,
		))

	return app
}

func allRoutes(app *fiber.App, routeParams types.RouteParams, authParams auth.Params) {
	Utility(app)
	auth.Auth(authParams)
	users.UserRoutes(routeParams)
	clubs.ClubRoutes(routeParams)
	socials.Social(routeParams)
	leader.Leader(routeParams)
	tags.Tag(routeParams)
	categories.CategoryRoutes(routeParams)
	events.EventRoutes(routeParams)
	files.File(routeParams)
	search.SearchRoutes(routeParams)
}

func newFiberApp(appSettings config.ApplicationSettings) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  go_json.Marshal,
		JSONDecoder:  go_json.Unmarshal,
		ErrorHandler: utilities.ErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     appSettings.ApplicationURL(),
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		ExposeHeaders:    "Authorization, Content-Length, Content-Type",
	}))

	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(otelfiber.Middleware())

	return app
}
