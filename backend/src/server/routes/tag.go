package routes

import (
	p "github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func Tag(tagParams types.RouteParams) {
	tagController := controllers.NewTagController(services.NewTagService(tagParams.ServiceParams))

	tags := tagParams.Router.Group("/tags")

	tags.Get("/", tagController.GetTags)
	tags.Post("/", tagParams.authMiddleware.Authorize(p.CreateAll), tagController.CreateTag)

	tagID := tags.Group("/:tagID")

	tagID.Get("/", tagController.GetTag)
	tagID.Patch("/", tagParams.authMiddleware.Authorize(p.WriteAll), tagController.UpdateTag)
	tagID.Delete("/", tagParams.authMiddleware.Authorize(p.DeleteAll), tagController.DeleteTag)
}
