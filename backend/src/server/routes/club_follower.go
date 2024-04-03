package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func ClubFollower(clubParams types.RouteParams) {
	clubFollowerController := controllers.NewClubFollowerController(services.NewClubFollowerService(clubParams.ServiceParams))

	clubFollower := clubParams.Router.Group("/followers")

	// api/clubs/:clubID/followers/*
	clubFollower.Get("/", clubParams.AuthMiddleware.ClubAuthorizeById, clubFollowerController.GetClubFollowers)
}
