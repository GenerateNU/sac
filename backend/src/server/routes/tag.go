package routes

import (
	p "github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func Tag(router fiber.Router, tagService services.TagServiceInterface, authMiddleware *middleware.AuthMiddlewareService) {
	tagController := controllers.NewTagController(tagService)

	tags := router.Group("/tags")

	tags.Get("/:tagID", tagController.GetTag)
	tags.Get("/", tagController.GetTags)
	tags.Post("/", authMiddleware.Authorize(p.CreateAll), tagController.CreateTag)
	tags.Patch("/:tagID", authMiddleware.Authorize(p.WriteAll), tagController.UpdateTag)
	tags.Delete("/:tagID", authMiddleware.Authorize(p.DeleteAll), tagController.DeleteTag)
}
