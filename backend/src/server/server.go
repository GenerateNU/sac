package server

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

// @title SAC API
// @version 1.0
// @description Backend Server for SAC App

// @contact.name	David Oduneye and Garrett Ladley
// @contact.email	oduneye.d@northeastern.edu and ladley.g@northeastern.edu
// @host 127.0.0.1:8080
// @BasePath /
func Init(db *gorm.DB) *fiber.App {
	app := newFiberApp()

	validate := validator.New(validator.WithRequiredStructEnabled())
	// MARK: Custom validator tags can be registered here.
	utilities.RegisterCustomValidators(validate)

	utilityRoutes(app)

	middlewareService := middleware.NewMiddlewareService(db, validate)

	apiv1 := app.Group("/api/v1")
	apiv1.Use(middlewareService.Authenticate)

	userRoutes(apiv1, services.NewUserService(db, validate), middlewareService)
	categoryRoutes(apiv1, services.NewCategoryService(db, validate))
	tagRoutes(apiv1, services.NewTagService(db, validate))

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

func utilityRoutes(router fiber.Router) {
	router.Get("/swagger/*", swagger.HandlerDefault)
	router.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
}

func userRoutes(router fiber.Router, userService services.UserServiceInterface, middlewareService middleware.MiddlewareInterface) {
	userController := controllers.NewUserController(userService)

	// api/v1/users/*
	users := router.Group("/users")
	users.Post("/", userController.CreateUser)
	users.Get("/", userController.GetUsers)

	// api/v1/users/:id/*
	// usersID := users.Group("/:id")
	// users.Use(middlewareService.UserAuthorizeById)
	users.Get("/:id", middlewareService.UserAuthorizeById, middlewareService.Authorize(models.UserRead), userController.GetUser)
	users.Patch("/:id", middlewareService.UserAuthorizeById, middlewareService.Authorize(models.UserWrite), userController.UpdateUser)
	users.Delete("/:id", middlewareService.UserAuthorizeById, middlewareService.Authorize(models.UserDelete), userController.DeleteUser)

	// api/v1/auth/*
	users.Get("/auth/logout", userController.Logout)
	users.Get("/auth/refresh", userController.Refresh)
	users.Post("/auth/login", userController.Login)
}

func categoryRoutes(router fiber.Router, categoryService services.CategoryServiceInterface) {
	categoryController := controllers.NewCategoryController(categoryService)

	categories := router.Group("/categories")

	categories.Post("/", categoryController.CreateCategory)
	categories.Get("/", categoryController.GetCategories)
	categories.Get("/:id", categoryController.GetCategory)
	categories.Delete("/:id", categoryController.DeleteCategory)
	categories.Patch("/:id", categoryController.UpdateCategory)
}

func tagRoutes(router fiber.Router, tagService services.TagServiceInterface) {
	tagController := controllers.NewTagController(tagService)

	tags := router.Group("/tags")

	tags.Get("/:id", tagController.GetTag)
	tags.Post("/", tagController.CreateTag)
	tags.Patch("/:id", tagController.UpdateTag)
	tags.Delete("/:id", tagController.DeleteTag)
}
