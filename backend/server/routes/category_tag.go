package routes

import (
	"github.com/GenerateNU/sac/backend/controllers"
	"github.com/GenerateNU/sac/backend/services"
	"github.com/GenerateNU/sac/backend/types"
)

func CategoryTag(categoryParams types.RouteParams) {
	categoryTagController := controllers.NewCategoryTagController(services.NewCategoryTagService(categoryParams.ServiceParams))

	// api/v1/categories/:categoryID/tags/*
	categoryTags := categoryParams.Router.Group("/tags")

	categoryTags.Get("/", categoryTagController.GetTagsByCategory)
	categoryTags.Get("/:tagID", categoryTagController.GetTagByCategory)
}
