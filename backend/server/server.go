package server

import (
	"fmt"

	"github.com/garrettladley/fiberpaginate"
	go_json "github.com/goccy/go-json"

	authenticator "github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/database/store"

	auth "github.com/GenerateNU/sac/backend/entities/auth/base"
	categories "github.com/GenerateNU/sac/backend/entities/categories/base"
	clubs "github.com/GenerateNU/sac/backend/entities/clubs/base"
	contacts "github.com/GenerateNU/sac/backend/entities/contacts/base"
	events "github.com/GenerateNU/sac/backend/entities/events/base"
	files "github.com/GenerateNU/sac/backend/entities/files/base"
	oauth "github.com/GenerateNU/sac/backend/entities/oauth/base"
	pocs "github.com/GenerateNU/sac/backend/entities/pocs/base"
	tags "github.com/GenerateNU/sac/backend/entities/tags/base"
	users "github.com/GenerateNU/sac/backend/entities/users/base"
	"github.com/GenerateNU/sac/backend/integrations"
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	utilityMiddleware "github.com/GenerateNU/sac/backend/middleware/utility"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/golang-jwt/jwt"

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

	jwt := authenticator.NewJWTClient(settings.Auth, jwt.SigningMethodHS256)
	authMiddleware := authMiddleware.New(db, validate, settings.Auth, stores)
	utilityMiddleware := utilityMiddleware.New(fiberpaginate.New(), stores)

	apiv1 := app.Group("/api/v1")

	routeParams := types.NewRouteParams(
		apiv1,
		authMiddleware,
		utilityMiddleware,
		types.NewServiceParams(
			db,
			validate,
			&settings.Auth,
			jwt,
			&settings.Calendar,
			stores,
			integrations,
		),
	)

	allRoutes(app, routeParams)

	return app
}

func allRoutes(app *fiber.App, routeParams types.RouteParams) {
	Utility(app)
	auth.Auth(routeParams)
	users.UserRoutes(routeParams)
	clubs.ClubRoutes(routeParams)
	contacts.Contact(routeParams)
	pocs.PointOfContact(routeParams)
	tags.Tag(routeParams)
	categories.CategoryRoutes(routeParams)
	events.EventRoutes(routeParams)
	files.File(routeParams)
	oauth.OAuth(routeParams)
}

func newFiberApp(appSettings config.ApplicationSettings) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  go_json.Marshal,
		JSONDecoder:  go_json.Unmarshal,
		ErrorHandler: utilities.ErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("http://%s:%d", appSettings.Host, appSettings.Port),
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
