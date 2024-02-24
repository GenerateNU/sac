package routes

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CategoryRoutes(router fiber.Router, db *gorm.DB, validate *validator.Validate, authMiddleware *middleware.AuthMiddlewareService) {
	categoryIDRoute := Category(router, services.NewCategoryService(db, validate), authMiddleware)

	CategoryTag(categoryIDRoute, services.NewCategoryTagService(db, validate))
}

func Category(router fiber.Router, categoryService services.CategoryServiceInterface, authMiddleware *middleware.AuthMiddlewareService) fiber.Router {
	categoryController := controllers.NewCategoryController(categoryService)

	// api/v1/categories/*
	categories := router.Group("/categories")

	categories.Post("/", authMiddleware.Authorize(auth.CreateAll), categoryController.CreateCategory)
	categories.Get("/", categoryController.GetCategories)
	categories.Get("/:categoryID", categoryController.GetCategory)
	categories.Delete("/:categoryID", authMiddleware.Authorize(auth.DeleteAll), categoryController.DeleteCategory)
	categories.Patch("/:categoryID", authMiddleware.Authorize(auth.WriteAll), categoryController.UpdateCategory)

	return categories
}
