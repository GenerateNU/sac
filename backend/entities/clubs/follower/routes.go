package follower

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func ClubFollower(clubParams types.RouteParams) {
	clubFollowerController := NewClubFollowerController(NewClubFollowerService(clubParams.ServiceParams))

	clubFollowers := clubParams.Router.Group("/followers")

	// api/clubs/:clubID/followers/*
	clubFollowers.Get("/", utilities.Make(clubFollowerController.GetClubFollowers))
}
