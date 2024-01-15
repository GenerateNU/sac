package server

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"

	"github.com/goccy/go-json"

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

	utilityRoutes(app)

	apiv1 := app.Group("/api/v1")

	userRoutes(apiv1, &services.UserService{DB: db})
	categoryRoutes(apiv1, &services.CategoryService{DB: db})
	tagRoutes(apiv1, &services.TagService{DB: db})

	return app
}

func newFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New())
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

func userRoutes(router fiber.Router, userService services.UserServiceInterface) {
	userController := controllers.NewUserController(userService)

	users := router.Group("/users")

	users.Get("/", userController.GetAllUsers)
}

func categoryRoutes(router fiber.Router, categoryService services.CategoryServiceInterface) {
	categoryController := controllers.NewCategoryController(categoryService)

	categories := router.Group("/categories")

	categories.Post("/", categoryController.CreateCategory)
}

func tagRoutes(router fiber.Router, tagService services.TagServiceInterface) {
	tagController := controllers.NewTagController(tagService)

	tags := router.Group("/tags")

	tags.Post("/", tagController.CreateTag)
	tags.Get("/:id", tagController.GetTag)
	tags.Patch("/:id", tagController.UpdateTag)
	tags.Delete("/:id", tagController.DeleteTag)
}
