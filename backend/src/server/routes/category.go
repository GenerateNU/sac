package routes

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(categoryParams types.RouteParams) {
	categoryIDRoute := Category(categoryParams)

	// update the router in params
	categoryParams.Router = categoryIDRoute

	CategoryTag(categoryParams)
}

func Category(categoryParams types.RouteParams) fiber.Router {
	categoryController := controllers.NewCategoryController(services.NewCategoryService(categoryParams.ServiceParams))

	// api/v1/categories/*
	categories := categoryParams.Router.Group("/categories")

	categories.Post("/", categoryParams.AuthMiddleware.Authorize(auth.CreateAll), categoryController.CreateCategory)
	categories.Get("/", categoryController.GetCategories)

	// api/v1/categories/:categoryID/*
	categoriesID := categories.Group("/:categoryID")

	categoriesID.Get("/", categoryController.GetCategory)
	categoriesID.Delete("/", categoryParams.AuthMiddleware.Authorize(auth.DeleteAll), categoryController.DeleteCategory)
	categoriesID.Patch("/", categoryParams.AuthMiddleware.Authorize(auth.WriteAll), categoryController.UpdateCategory)

	return categoriesID
}
