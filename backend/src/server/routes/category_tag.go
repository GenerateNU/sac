package routes

import (
	"github.com/GenerateNU/sac/backend/src/controllers"
	"github.com/GenerateNU/sac/backend/src/services"
	"github.com/GenerateNU/sac/backend/src/types"
)

func CategoryTag(categoryParams types.RouteParams) {
	categoryTagController := controllers.NewCategoryTagController(services.NewCategoryTagService(categoryParams.ServiceParams))

	// api/v1/categories/:categoryID/tags/*
	categoryTags := categoryParams.Router.Group("/tags")

	categoryTags.Get("/", categoryTagController.GetTagsByCategory)
	categoryTags.Get("/:tagID", categoryTagController.GetTagByCategory)
}
