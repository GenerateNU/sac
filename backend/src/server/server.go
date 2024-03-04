package server

import (
	"encoding/json"
	"fmt"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/email"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/server/routes"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/gofiber/fiber/v2"
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
// @schemes http https
func Init(db *gorm.DB, settings config.Settings) *fiber.App {
	app := newFiberApp(settings.Application)

	validate, err := utilities.RegisterCustomValidators()
	if err != nil {
		panic(fmt.Sprintf("Error registering custom validators: %s", err))
	}

	authMiddleware := middleware.NewAuthAuthMiddlewareService(db, validate, settings.Auth)
	emailService := email.NewEmailClient(settings.ResendSettings)

	apiv1 := app.Group("/api/v1")
	apiv1.Use(authMiddleware.Authenticate)

	routeParams := types.RouteParams{
		Router:         apiv1,
		Settings:       settings.Auth,
		AuthMiddleware: authMiddleware,
		ServiceParams: types.ServiceParams{
			DB:       db,
			Validate: validate,
			Email:    emailService,
		},
	}

	routes.Utility(app)
	routes.Auth(routeParams)
	routes.UserRoutes(routeParams)
	routes.Contact(routeParams)
	routes.ClubRoutes(routeParams)
	routes.Tag(routeParams)
	routes.CategoryRoutes(routeParams)
	routes.Event(routeParams)

	return app
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
