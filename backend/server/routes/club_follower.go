package routes

import (
	"github.com/GenerateNU/sac/backend/controllers"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/types"
)

func ClubFollower(clubParams types.RouteParams) {
	clubFollowerController := controllers.NewClubFollowerController(services.NewClubFollowerService(clubParams.ServiceParams))

	clubFollowers := clubParams.Router.Group("/followers")

	// api/clubs/:clubID/followers/*
	clubFollowers.Get("/", clubFollowerController.GetClubFollowers)
}
