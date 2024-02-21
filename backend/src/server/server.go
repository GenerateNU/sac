package server

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/server/routes"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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
func Init(db *gorm.DB, settings config.Settings) *fiber.App {
	app := newFiberApp()

	validate, err := utilities.RegisterCustomValidators()
	if err != nil {
		panic(fmt.Sprintf("Error registering custom validators: %s", err))
	}

	authMiddleware := middleware.NewAuthAuthMiddlewareService(db, validate, settings.Auth)

	apiv1 := app.Group("/api/v1")
	apiv1.Use(authMiddleware.Authenticate)

	routes.Utility(app)
	routes.Auth(apiv1, services.NewAuthService(db, validate), settings.Auth, authMiddleware)
	routes.UserRoutes(apiv1, db, validate, authMiddleware)
	routes.Contact(apiv1, services.NewContactService(db, validate))
	routes.ClubRoutes(apiv1, db, validate, authMiddleware)
	routes.Tag(apiv1, services.NewTagService(db, validate), authMiddleware)
	routes.CategoryRoutes(apiv1, db, validate, authMiddleware)
	routes.Event(apiv1, services.NewEventService(db, validate), authMiddleware)

	return app
}

func newFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(limiter.New()) // TODO: currently wrapping the whole app, makes more sense for specific endpoints like update password

	return app
}
