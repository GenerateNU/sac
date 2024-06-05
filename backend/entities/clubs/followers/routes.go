package followers

import (
	authMiddleware "github.com/GenerateNU/sac/backend/middleware/auth"

	"github.com/GenerateNU/sac/backend/types"
)

func ClubFollower(clubParams types.RouteParams) {
	clubFollowerController := NewClubFollowerController(NewClubFollowerService(clubParams.ServiceParams))

	clubFollowers := clubParams.Router.Group("/followers")

	// api/clubs/:clubID/followers/*
	clubFollowers.Get("/", clubParams.UtilityMiddleware.Paginator, clubFollowerController.GetClubFollowers)
	clubFollowers.Post(
		"/:userID",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubFollowerController.CreateClubFollowing,
	)
	clubFollowers.Delete(
		"/:userID",
		authMiddleware.AttachExtractor(clubParams.AuthMiddleware.ClubAuthorizeById, authMiddleware.ExtractFromParams("clubID")),
		clubFollowerController.DeleteClubFollowing,
	)
}
