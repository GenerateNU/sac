package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Tag(router fiber.Router, tagService services.TagServiceInterface) {
	tagController := controllers.NewTagController(tagService)

	tags := router.Group("/tags")

	tags.Post("/", tagController.CreateTag)

	tagID := tags.Group("/:tagID")

	tagID.Get("/", tagController.GetTag)
	tagID.Patch("/", tagController.UpdateTag)
	tagID.Delete("/", tagController.DeleteTag)
}
