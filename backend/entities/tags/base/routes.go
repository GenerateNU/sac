package base

import (
	"github.com/GenerateNU/sac/backend/permission"
	"github.com/GenerateNU/sac/backend/types"
)

func Tag(tagParams types.RouteParams) {
	tagController := NewTagController(NewTagService(tagParams.ServiceParams))

	tags := tagParams.Router.Group("/tags")

	tags.Get("/", tagController.GetTags)
	tags.Post("/", tagParams.AuthMiddleware.Authorize(permission.CreateAll), tagController.CreateTag)

	tagID := tags.Group("/:tagID")

	tagID.Get("/", tagController.GetTag)
	tagID.Patch("/", tagParams.AuthMiddleware.Authorize(permission.WriteAll), tagController.UpdateTag)
	tagID.Delete("/", tagParams.AuthMiddleware.Authorize(permission.DeleteAll), tagController.DeleteTag)
}
