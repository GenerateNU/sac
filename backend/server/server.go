package server

import (
	"encoding/json"
	"fmt"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/server/routes"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
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
func Init(db *gorm.DB, dependencies types.Dependencies, settings config.Settings) *fiber.App {
	app := newFiberApp(settings.Application)

	validate, err := utilities.RegisterCustomValidators()
	if err != nil {
		panic(fmt.Sprintf("Error registering custom validators: %s", err))
	}

	jwt := auth.NewJWTClient(settings.Auth, jwt.SigningMethodHS256)
	authMiddleware := middleware.NewAuthAuthMiddlewareService(db, validate, settings.Auth)

	apiv1 := app.Group("/api/v1")
	apiv1.Use(authMiddleware.Authenticate)

	routeParams := types.RouteParams{
		Router:         apiv1,
		AuthMiddleware: authMiddleware,
		ServiceParams: types.ServiceParams{
			DB:           db,
			Validate:     validate,
			AuthSettings: &settings.Auth,
			JWT:          jwt,
			Deps:         dependencies,
		},
	}

	allRoutes(app, routeParams)

	return app
}

func allRoutes(app *fiber.App, routeParams types.RouteParams) {
	routes.Utility(app)
	routes.Auth(routeParams)
	routes.UserRoutes(routeParams)
	routes.ClubRoutes(routeParams)
	routes.Contact(routeParams)
	routes.PointOfContact(routeParams)
	routes.Tag(routeParams)
	routes.CategoryRoutes(routeParams)
	routes.Event(routeParams)
	routes.File(routeParams)
}

func newFiberApp(appSettings config.ApplicationSettings) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("http://%s:%d", appSettings.Host, appSettings.Port),
		AllowCredentials: true,
	}))
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))

	return app
}
