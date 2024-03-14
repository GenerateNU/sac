package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/middleware"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func ClubTag(clubIDRouter fiber.Router, clubTagService services.ClubTagServiceInterface, authMiddleware *middleware.AuthMiddlewareService) {
	clubTagController := controllers.NewClubTagController(clubTagService)

	clubTag := clubIDRouter.Group("/tags")

	clubTag.Get("/", clubTagController.GetClubTags)
	clubTag.Post("/", authMiddleware.ClubAuthorizeById, clubTagController.CreateClubTags)
	clubTag.Delete("/:tagID", authMiddleware.ClubAuthorizeById, clubTagController.DeleteClubTag)
}
