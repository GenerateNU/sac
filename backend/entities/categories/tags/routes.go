package tags

import (
	"github.com/GenerateNU/sac/backend/types"
)

func CategoryTag(categoryParams types.RouteParams) {
	categoryTagController := NewCategoryTagController(NewCategoryTagService(categoryParams.ServiceParams))

	// api/v1/categories/:categoryID/tags/*
	categoryTags := categoryParams.Router.Group("/tags")

	categoryTags.Get("/", categoryParams.UtilityMiddleware.Paginator, categoryTagController.GetTagsByCategory)
	categoryTags.Get("/:tagID", categoryTagController.GetTagByCategory)
}
