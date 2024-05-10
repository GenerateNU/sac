package follower

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func UserFollower(userParams types.RouteParams) {
	userFollowerController := NewUserFollowerController(NewUserFollowerService(userParams.ServiceParams))

	// api/v1/users/:userID/follower/*
	userFollower := userParams.Router.Group("/follower")

	userFollower.Get("/", utilities.Make(userFollowerController.GetFollowing))
	userFollower.Post("/:clubID", userParams.AuthMiddleware.UserAuthorizeById, utilities.Make(userFollowerController.CreateFollowing))
	userFollower.Delete("/:clubID", userParams.AuthMiddleware.UserAuthorizeById, utilities.Make(userFollowerController.DeleteFollowing))
}
