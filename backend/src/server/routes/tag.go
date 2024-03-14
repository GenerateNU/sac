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

	tag := router.Group("/tags")

	tag.Get("/", tagController.GetTags)
	tag.Post("/", authMiddleware.Authorize(p.CreateAll), tagController.CreateTag)

	tagID := tag.Group("/:tagID")

	tagID.Get("/", tagController.GetTag)
	tagID.Patch("/", authMiddleware.Authorize(p.WriteAll), tagController.UpdateTag)
	tagID.Delete("/", authMiddleware.Authorize(p.DeleteAll), tagController.DeleteTag)
}
