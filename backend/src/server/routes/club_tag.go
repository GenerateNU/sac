package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/gofiber/fiber/v2"
)

func ClubTag(router fiber.Router, clubTagService services.ClubTagServiceInterface) {
	clubTagController := controllers.NewClubTagController(clubTagService)

	clubTags := router.Group("/tags")

	clubTags.Post("/", clubTagController.CreateClubTags)
	clubTags.Get("/", clubTagController.GetClubTags)

	tagIDs := clubTags.Group("/:tagID")
	tagIDs.Delete("/", clubTagController.DeleteClubTag)
}
