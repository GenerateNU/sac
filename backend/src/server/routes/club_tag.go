package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func ClubTag(router fiber.Router, clubTagService services.ClubTagServiceInterface, authMiddleware *middleware.AuthMiddlewareService) {
	clubTagController := controllers.NewClubTagController(clubTagService)

	clubTags := router.Group("/tags")

	clubTags.Get("/", clubTagController.GetClubTags)
	clubTags.Post("/", authMiddleware.ClubAuthorizeById, clubTagController.CreateClubTags)
	clubTags.Delete("/:tagID", authMiddleware.ClubAuthorizeById, clubTagController.DeleteClubTag)
}
