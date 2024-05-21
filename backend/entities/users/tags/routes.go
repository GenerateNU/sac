package tags

import (
	"github.com/GenerateNU/sac/backend/types"
)

func UserTag(userParams types.RouteParams) {
	userTagController := NewUserTagController(NewUserTagService(userParams.ServiceParams))

	// api/v1/user/:userID/tags/*
	userTags := userParams.Router.Group("/tags")

	userTags.Get("/", userTagController.GetUserTags)
	userTags.Post("/", userParams.AuthMiddleware.UserAuthorizeById, userTagController.CreateUserTags)

	tagID := userTags.Group("/:tagID")
	tagID.Delete("/", userParams.AuthMiddleware.UserAuthorizeById, userTagController.DeleteUserTag)
}
