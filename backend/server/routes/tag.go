package routes

import (
	p "github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/controllers"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/types"
)

func Tag(tagParams types.RouteParams) {
	tagController := controllers.NewTagController(services.NewTagService(tagParams.ServiceParams))

	tags := tagParams.Router.Group("/tags")

	tags.Get("/", tagController.GetTags)
	tags.Post("/", tagParams.AuthMiddleware.Authorize(p.CreateAll), tagController.CreateTag)

	tagID := tags.Group("/:tagID")

	tagID.Get("/", tagController.GetTag)
	tagID.Patch("/", tagParams.AuthMiddleware.Authorize(p.WriteAll), tagController.UpdateTag)
	tagID.Delete("/", tagParams.AuthMiddleware.Authorize(p.DeleteAll), tagController.DeleteTag)
}
