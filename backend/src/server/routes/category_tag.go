package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func CategoryTag(categoryRouter fiber.Router, categoryTagService services.CategoryTagServiceInterface) {
	categoryTagController := controllers.NewCategoryTagController(categoryTagService)

	categoryTags := categoryRouter.Group("/:categoryID/tags")

	categoryTags.Get("/", categoryTagController.GetTagsByCategory)

	categoryID := categoryTags.Group("/:tagID")

	categoryID.Get("/", categoryTagController.GetTagByCategory)
}
