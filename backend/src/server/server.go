package server

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
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

	apiv1 := app.Group("/api/v1")

	userRoutes(apiv1, &services.UserService{DB: db, Validate: validate})
	clubRoutes(apiv1, &services.ClubService{DB: db, Validate: validate})
	categoryRouter := categoryRoutes(apiv1, &services.CategoryService{DB: db, Validate: validate})
	tagRoutes(categoryRouter, &services.TagService{DB: db, Validate: validate})

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

	users.Post("/", userController.CreateUser)
	users.Get("/", userController.GetUsers)
	users.Get("/:id", userController.GetUser)
	users.Patch("/:id", userController.UpdateUser)
	users.Delete("/:id", userController.DeleteUser)

	users.Get("/:userID/membership", userController.GetUserMemberships)

	userTags := users.Group("/:uid/tags")

	userTags.Post("/", userController.CreateUserTags)
	userTags.Get("/", userController.GetUserTags)
}

func clubRoutes(router fiber.Router, clubService services.ClubServiceInterface) {
	clubController := controllers.NewClubController(clubService)

	clubs := router.Group("/clubs")

	clubs.Get("/", clubController.GetAllClubs)
	clubs.Post("/", clubController.CreateClub)
	clubs.Get("/:id", clubController.GetClub)
	clubs.Patch("/:id", clubController.UpdateClub)
	clubs.Delete("/:id", clubController.DeleteClub)

	memberships := clubs.Group("/:clubID/membership")

	memberships.Get("/", clubController.GetClubMembers)           // good
	memberships.Post("/:userID", clubController.CreateMembership) // good
	memberships.Post("/", clubController.CreateMembershipsByEmail)
	memberships.Delete("/:userID", clubController.DeleteMembership) // good
	memberships.Delete("/", clubController.DeleteMemberships)
}

func categoryRoutes(router fiber.Router, categoryService services.CategoryServiceInterface) fiber.Router {
	categoryController := controllers.NewCategoryController(categoryService)

	categories := router.Group("/categories")

	categories.Post("/", categoryController.CreateCategory)
	categories.Get("/", categoryController.GetCategories)
	categories.Get("/:id", categoryController.GetCategory)
	categories.Delete("/:id", categoryController.DeleteCategory)
	categories.Patch("/:id", categoryController.UpdateCategory)

	return categories
}

func tagRoutes(router fiber.Router, tagService services.TagServiceInterface) {
	tagController := controllers.NewTagController(tagService)

	tags := router.Group("/:categoryID/tags")

	tags.Get("/:tagID", tagController.GetTag)
	tags.Post("/", tagController.CreateTag)
	tags.Patch("/:tagID", tagController.UpdateTag)
	tags.Delete("/:tagID", tagController.DeleteTag)
}
