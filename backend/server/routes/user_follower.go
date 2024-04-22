package routes

import (
	"github.com/GenerateNU/sac/backend/controllers"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/types"
)

func UserFollower(userParams types.RouteParams) {
	userFollowerController := controllers.NewUserFollowerController(services.NewUserFollowerService(userParams.ServiceParams))

	// api/v1/users/:userID/follower/*
	userFollower := userParams.Router.Group("/follower")

	userFollower.Get("/", userFollowerController.GetFollowing)
	userFollower.Post("/:clubID", userParams.AuthMiddleware.UserAuthorizeById, userFollowerController.CreateFollowing)
	userFollower.Delete("/:clubID", userParams.AuthMiddleware.UserAuthorizeById, userFollowerController.DeleteFollowing)
}
