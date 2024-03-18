package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func UserTag(userParams types.RouteParams) {
	userTagController := controllers.NewUserTagController(services.NewUserTagService(userParams.ServiceParams))

	// api/v1/user/:userID/tags/*
	userTags := userParams.Router.Group("/tags")

	userTags.Post("/", userTagController.CreateUserTags)
	userTags.Get("/", userTagController.GetUserTags)

	tagID := userTags.Group("/:tagID")
	tagID.Delete("/", userTagController.DeleteUserTag)
}
