package tag

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubTag(clubParams types.RouteParams) {
	clubTagController := NewClubTagController(NewClubTagService(clubParams.ServiceParams))

	clubTags := clubParams.Router.Group("/tags")

	// api/v1/clubs/:clubID/tags/*
	clubTags.Get("/", clubTagController.GetClubTags)
	clubTags.Post(
		"/",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		clubTagController.CreateClubTags,
	)
	clubTags.Delete(
		"/:tagID",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		clubTagController.DeleteClubTag,
	)
}
