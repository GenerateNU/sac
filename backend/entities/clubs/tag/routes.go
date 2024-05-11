package tag

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func ClubTag(clubParams types.RouteParams) {
	clubTagController := NewClubTagController(NewClubTagService(clubParams.ServiceParams))

	clubTags := clubParams.Router.Group("/tags")

	// api/v1/clubs/:clubID/tags/*
	clubTags.Get("/", utilities.Make(clubTagController.GetClubTags))
	clubTags.Post(
		"/",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		utilities.Make(clubTagController.CreateClubTags),
	)
	clubTags.Delete(
		"/:tagID",
		middleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, middleware.ExtractFromParams("clubID")),
		utilities.Make(clubTagController.DeleteClubTag),
	)
}
