package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Tag(router fiber.Router, tagService services.TagServiceInterface) {
	tagController := controllers.NewTagController(tagService)

	tags := router.Group("/tags")

	tags.Get("/:tagID", tagController.GetTag)
	tags.Post("/", tagController.CreateTag)
	tags.Patch("/:tagID", tagController.UpdateTag)
	tags.Delete("/:tagID", tagController.DeleteTag)
}
