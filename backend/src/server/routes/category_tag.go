package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func CategoryTag(router fiber.Router, categoryTagService services.CategoryTagServiceInterface) {
	categoryTagController := controllers.NewCategoryTagController(categoryTagService)

	categoryTags := router.Group("/:categoryID/tags")

	categoryTags.Get("/", categoryTagController.GetTagsByCategory)
	categoryTags.Get("/:tagID", categoryTagController.GetTagByCategory)
}
