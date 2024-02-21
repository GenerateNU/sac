package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Category(router fiber.Router, categoryService services.CategoryServiceInterface) fiber.Router {
	categoryController := controllers.NewCategoryController(categoryService)

	categories := router.Group("/categories")

	categories.Post("/", categoryController.CreateCategory)
	categories.Get("/", categoryController.GetCategories)

	categoryID := categories.Group("/:categoryID")

	categoryID.Get("/", categoryController.GetCategory)
	categoryID.Delete("/", categoryController.DeleteCategory)
	categoryID.Patch("/", categoryController.UpdateCategory)

	return categories
}
