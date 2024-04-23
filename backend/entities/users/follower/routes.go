package follower

import (
	"github.com/GenerateNU/sac/backend/types"
)

func UserFollower(userParams types.RouteParams) {
	userFollowerController := NewUserFollowerController(NewUserFollowerService(userParams.ServiceParams))

	// api/v1/users/:userID/follower/*
	userFollower := userParams.Router.Group("/follower")

	userFollower.Get("/", userFollowerController.GetFollowing)
	userFollower.Post("/:clubID", userParams.AuthMiddleware.UserAuthorizeById, userFollowerController.CreateFollowing)
	userFollower.Delete("/:clubID", userParams.AuthMiddleware.UserAuthorizeById, userFollowerController.DeleteFollowing)
}
