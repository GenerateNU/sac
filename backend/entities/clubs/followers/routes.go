package followers

import (
	"github.com/GenerateNU/sac/backend/types"
)

func ClubFollower(clubParams types.RouteParams) {
	clubFollowerController := NewClubFollowerController(NewClubFollowerService(clubParams.ServiceParams))

	clubFollowers := clubParams.Router.Group("/followers")

	// api/clubs/:clubID/followers/*
	clubFollowers.Get("/", clubParams.UtilityMiddleware.Paginator, clubFollowerController.GetClubFollowers)
	clubFollowers.Post(
		"/:userID",
		clubParams.AuthMiddleware.UserAuthorizeById,
		clubFollowerController.CreateClubFollowing,
	)
	clubFollowers.Delete(
		"/:userID",
		clubParams.AuthMiddleware.UserAuthorizeById,
		clubFollowerController.DeleteClubFollowing,
	)
}
