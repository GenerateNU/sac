package server

import (
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/server/routes"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/goccy/go-json"

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
func Init(db *gorm.DB, settings config.Settings) *fiber.App {
	app := newFiberApp()

	validate, err := utilities.RegisterCustomValidators()
	if err != nil {
		panic(err)
	}

	middlewareService := middleware.NewMiddlewareService(db, validate, settings.Auth)

	apiv1 := app.Group("/api/v1")
	apiv1.Use(middlewareService.Authenticate)

	routes.Utility(app)

	routes.Auth(apiv1, services.NewAuthService(db, validate), settings.Auth)

	userRouter := routes.User(apiv1, services.NewUserService(db, validate), middlewareService)
	routes.UserTag(userRouter, services.NewUserTagService(db, validate))

	routes.Contact(apiv1, services.NewContactService(db, validate))

	clubsRouter := routes.Club(apiv1, services.NewClubService(db, validate), middlewareService)
	routes.ClubContact(clubsRouter, services.NewClubContactService(db, validate))

	routes.Tag(apiv1, services.NewTagService(db, validate))

	categoryRouter := routes.Category(apiv1, services.NewCategoryService(db, validate))
	routes.CategoryTag(categoryRouter, services.NewCategoryTagService(db, validate))

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

	return app
}

func userFollowerRoutes(router fiber.Router, userFollowerService services.UserFollowerServiceInterface) {
	userFollowerController := controllers.NewUserFollowerController(userFollowerService)

	userFollower := router.Group("/:userID/follower")

	userFollower.Put("/:clubID", userFollowerController.CreateFollowing)
	userFollower.Delete("/:clubID", userFollowerController.DeleteFollowing)
	userFollower.Get("/", userFollowerController.GetAllFollowing)
}

