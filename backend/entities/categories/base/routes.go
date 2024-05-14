package base

import (
	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/categories/tag"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(categoryParams types.RouteParams) {
	categoryIDRoute := Category(categoryParams)

	// update the router in params
	categoryParams.Router = categoryIDRoute

	tag.CategoryTag(categoryParams)
}

func Category(categoryParams types.RouteParams) fiber.Router {
	categoryController := NewCategoryController(NewCategoryService(categoryParams.ServiceParams))

	// api/v1/categories/*
	categories := categoryParams.Router.Group("/categories")

	categories.Post("/", categoryParams.AuthMiddleware.Authorize(auth.CreateAll), categoryController.CreateCategory)
	categories.Get("/", middleware.Paginate, categoryController.GetCategories)

	// api/v1/categories/:categoryID/*
	categoryID := categories.Group("/:categoryID")

	categoryID.Get("/", categoryController.GetCategory)
	categoryID.Delete("/", categoryParams.AuthMiddleware.Authorize(auth.DeleteAll), categoryController.DeleteCategory)
	categoryID.Patch("/", categoryParams.AuthMiddleware.Authorize(auth.WriteAll), categoryController.UpdateCategory)

	return categoryID
}
