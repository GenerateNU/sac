package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func ClubTag(clubParams types.RouteParams) {
	clubTagController := controllers.NewClubTagController(services.NewClubTagService(clubParams.ServiceParams))

	clubTags := clubParams.Router.Group("/tags")

	clubTags.Get("/", clubTagController.GetClubTags)
	clubTags.Post("/", clubParams.AuthMiddleware.ClubAuthorizeById, clubTagController.CreateClubTags)
	clubTags.Delete("/:tagID", clubParams.AuthMiddleware.ClubAuthorizeById, clubTagController.DeleteClubTag)
}
