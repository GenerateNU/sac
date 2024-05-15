package follower

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/garrettladley/fiberpaginate"
)

func ClubFollower(clubParams types.RouteParams) {
	clubFollowerController := NewClubFollowerController(NewClubFollowerService(clubParams.ServiceParams))

	clubFollowers := clubParams.Router.Group("/followers")

	// api/clubs/:clubID/followers/*
	clubFollowers.Get("/", fiberpaginate.New(), clubFollowerController.GetClubFollowers)
}
