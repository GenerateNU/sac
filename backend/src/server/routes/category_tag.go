package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func CategoryTag(categoryIDRoute fiber.Router, categoryTagService services.CategoryTagServiceInterface) {
	categoryTagController := controllers.NewCategoryTagController(categoryTagService)

	// api/v1/categories/:categoryID/tags/*
	categoryTags := categoryIDRoute.Group("/tags")

	categoryTags.Get("/", categoryTagController.GetTagsByCategory)
	categoryTags.Get("/:tags", categoryTagController.GetTagByCategory)
}
