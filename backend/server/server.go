package server

import (
	"fmt"

	go_json "github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/config"
	authBase "github.com/GenerateNU/sac/backend/entities/auth/base"
	categories "github.com/GenerateNU/sac/backend/entities/categories/base"
	clubs "github.com/GenerateNU/sac/backend/entities/clubs/base"
	contacts "github.com/GenerateNU/sac/backend/entities/contacts/base"
	events "github.com/GenerateNU/sac/backend/entities/events/base"
	files "github.com/GenerateNU/sac/backend/entities/files/base"
	pocs "github.com/GenerateNU/sac/backend/entities/pocs/base"
	tags "github.com/GenerateNU/sac/backend/entities/tags/base"
	users "github.com/GenerateNU/sac/backend/entities/users/base"
	"github.com/GenerateNU/sac/backend/integrations"
	"github.com/GenerateNU/sac/backend/middleware"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
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
func Init(db *gorm.DB, integrations integrations.Integrations, settings config.Settings) *fiber.App {
	app := newFiberApp(settings.Application)

	validate, err := utilities.RegisterCustomValidators()
	if err != nil {
		panic(fmt.Sprintf("Error registering custom validators: %s", err))
	}

	jwt := auth.NewJWTClient(settings.Auth, jwt.SigningMethodHS256)
	authMiddleware := middleware.NewAuthMiddlewareService(db, validate, settings.Auth)

	apiv1 := app.Group("/api/v1")

	routeParams := types.RouteParams{
		Router:         apiv1,
		AuthMiddleware: authMiddleware,
		ServiceParams: types.ServiceParams{
			DB:           db,
			Validate:     validate,
			Auth:         &settings.Auth,
			JWT:          jwt,
			Integrations: integrations,
		},
	}

	allRoutes(app, routeParams)

	return app
}

func allRoutes(app *fiber.App, routeParams types.RouteParams) {
	Utility(app)
	authService := authBase.Auth(routeParams)
	users.UserRoutes(routeParams, authService)
	clubs.ClubRoutes(routeParams)
	contacts.Contact(routeParams)
	pocs.PointOfContact(routeParams)
	tags.Tag(routeParams)
	categories.CategoryRoutes(routeParams)
	events.Event(routeParams)
	files.File(routeParams)
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
	}))
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	return app
}
