package follower

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubFollower(clubParams types.RouteParams) {
	clubFollowerController := NewClubFollowerController(NewClubFollowerService(clubParams.ServiceParams))

	clubFollowers := clubParams.Router.Group("/followers")

	// api/clubs/:clubID/followers/*
	clubFollowers.Get("/", middleware.Paginate, clubFollowerController.GetClubFollowers)
}
