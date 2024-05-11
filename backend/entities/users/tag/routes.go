package tag

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func UserTag(userParams types.RouteParams) {
	userTagController := NewUserTagController(NewUserTagService(userParams.ServiceParams))

	// api/v1/user/:userID/tags/*
	userTags := userParams.Router.Group("/tags")

	userTags.Get("/", utilities.Make(userTagController.GetUserTags))
	userTags.Post("/", userParams.AuthMiddleware.UserAuthorizeById, utilities.Make(userTagController.CreateUserTags))

	tagID := userTags.Group("/:tagID")
	tagID.Delete("/", userParams.AuthMiddleware.UserAuthorizeById, utilities.Make(userTagController.DeleteUserTag))
}
