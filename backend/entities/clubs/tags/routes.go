package tags

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubTag(clubParams types.RouteParams) {
	clubTagController := NewClubTagController(NewClubTagService(clubParams.ServiceParams))

	clubTags := clubParams.Router.Group("/tags")

	// api/v1/clubs/:clubID/tags/*
	clubTags.Get("/", clubTagController.GetClubTags)
	clubTags.Post(
		"/",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubTagController.CreateClubTags,
	)
	clubTags.Delete(
		"/:tagID",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubTagController.DeleteClubTag,
	)
}
