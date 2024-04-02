package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func UserFollower(userParams types.RouteParams) {
	userFollowerController := controllers.NewUserFollowerController(services.NewUserFollowerService(userParams.ServiceParams))

	// api/v1/users/:userID/follower/*
	userFollower := userParams.Router.Group("/follower")

	userFollower.Get("/", userFollowerController.GetFollowing)
	userFollower.Post("/:clubID", userFollowerController.CreateFollowing)
	userFollower.Delete("/:clubID", userFollowerController.DeleteFollowing)
}
