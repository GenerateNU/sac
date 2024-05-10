package base

import (
	p "github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func Tag(tagParams types.RouteParams) {
	tagController := NewTagController(NewTagService(tagParams.ServiceParams))

	tags := tagParams.Router.Group("/tags")

	tags.Get("/", utilities.Make(tagController.GetTags))
	tags.Post("/", tagParams.AuthMiddleware.Authorize(p.CreateAll), utilities.Make(tagController.CreateTag))

	tagID := tags.Group("/:tagID")

	tagID.Get("/", utilities.Make(tagController.GetTag))
	tagID.Patch("/", tagParams.AuthMiddleware.Authorize(p.WriteAll), utilities.Make(tagController.UpdateTag))
	tagID.Delete("/", tagParams.AuthMiddleware.Authorize(p.DeleteAll), utilities.Make(tagController.DeleteTag))
}
