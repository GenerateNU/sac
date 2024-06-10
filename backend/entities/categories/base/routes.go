package base

import (
	"github.com/GenerateNU/sac/backend/entities/categories/tags"
	"github.com/GenerateNU/sac/backend/permission"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(categoryParams types.RouteParams) {
	categoryIDRoute := Category(categoryParams)

	// update the router in params
	categoryParams.Router = categoryIDRoute

	tags.CategoryTag(categoryParams)
}

func Category(categoryParams types.RouteParams) fiber.Router {
	categoryController := NewCategoryController(NewCategoryService(categoryParams.ServiceParams))

	// api/v1/categories/*
	categories := categoryParams.Router.Group("/categories")

	categories.Post("/", categoryParams.AuthMiddleware.Authorize(permission.CreateAll), categoryController.CreateCategory)
	categories.Get("/", categoryParams.UtilityMiddleware.Paginator, categoryController.GetCategories)

	// api/v1/categories/:categoryID/*
	categoryID := categories.Group("/:categoryID")

	categoryID.Get("/", categoryController.GetCategory)
	categoryID.Delete("/", categoryParams.AuthMiddleware.Authorize(permission.DeleteAll), categoryController.DeleteCategory)
	categoryID.Patch("/", categoryParams.AuthMiddleware.Authorize(permission.WriteAll), categoryController.UpdateCategory)

	return categoryID
}
