package follower

import (
	"github.com/GenerateNU/sac/backend/src/types"
)

func ClubFollower(clubParams types.RouteParams) {
	clubFollowerController := NewClubFollowerController(NewClubFollowerService(clubParams.ServiceParams))

	clubFollowers := clubParams.Router.Group("/followers")

	// api/clubs/:clubID/followers/*
	clubFollowers.Get("/", clubParams.AuthMiddleware.ClubAuthorizeById, clubFollowerController.GetClubFollowers)
}
