package tag

import (
	"github.com/GenerateNU/sac/backend/types"
)

func ClubTag(clubParams types.RouteParams) {
	clubTagController := NewClubTagController(NewClubTagService(clubParams.ServiceParams))

	clubTags := clubParams.Router.Group("/tags")

	clubTags.Get("/", clubTagController.GetClubTags)
	clubTags.Post("/", clubParams.AuthMiddleware.ClubAuthorizeById, clubTagController.CreateClubTags)
	clubTags.Delete("/:tagID", clubParams.AuthMiddleware.ClubAuthorizeById, clubTagController.DeleteClubTag)
}
