package server

import (
	"encoding/json"
	"fmt"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/email"
	"github.com/GenerateNU/sac/backend/src/file"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/search"
	"github.com/GenerateNU/sac/backend/src/server/routes"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
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
func Init(db *gorm.DB, pinecone search.PineconeClientInterface, settings config.Settings) *fiber.App {
	app := newFiberApp(settings.Application)

	validate, err := utilities.RegisterCustomValidators()
	if err != nil {
		panic(fmt.Sprintf("Error registering custom validators: %s", err))
	}

	authMiddleware := middleware.NewAuthAuthMiddlewareService(db, validate, settings.Auth)
	awsProvider := file.NewAWSProvider(settings.AWS)
	resend := email.NewResendClient(settings.ResendSettings, true)
	jwt := auth.NewJWTClient(settings.Auth, jwt.SigningMethodHS256)

	apiv1 := app.Group("/api/v1")
	apiv1.Use(authMiddleware.Authenticate)

	routeParams := types.RouteParams{
		Router:         apiv1,
		AuthMiddleware: authMiddleware,
		ServiceParams: types.ServiceParams{
			DB:           db,
			Validate:     validate,
			Resend:       resend,
			Pinecone:     &pinecone,
			JWT:          jwt,
			AuthSettings: settings.Auth,
		},
	}

	routes.Utility(app)
	// routes.Auth(apiv1, services.NewAuthService(db, validate), settings.Auth, authMiddleware)
	// routes.UserRoutes(apiv1, db, validate, authMiddleware)
	// routes.ClubRoutes(apiv1, db, validate, authMiddleware, awsProvider)
	// routes.Contact(apiv1, services.NewContactService(db, validate), authMiddleware)
	// routes.PointOfContact(apiv1, services.NewPointOfContactService(db, validate), authMiddleware)
	// routes.Tag(apiv1, services.NewTagService(db, validate), authMiddleware)
	// routes.CategoryRoutes(apiv1, db, validate, authMiddleware)
	// routes.Event(apiv1, services.NewEventService(db, validate), authMiddleware)
	// // routes.File(apiv1, services.NewFileService(db, validate, awsClient))
	// routes.Auth(routeParams)
	// routes.UserRoutes(routeParams)
	// routes.Contact(routeParams)
	// routes.ClubRoutes(routeParams)
	// routes.Tag(routeParams)
	// routes.CategoryRoutes(routeParams)
	// routes.Event(routeParams)

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
