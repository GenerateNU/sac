package tag

import (
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
)

func CategoryTag(categoryParams types.RouteParams) {
	categoryTagController := NewCategoryTagController(NewCategoryTagService(categoryParams.ServiceParams))

	// api/v1/categories/:categoryID/tags/*
	categoryTags := categoryParams.Router.Group("/tags")

	categoryTags.Get("/", middleware.Paginate, categoryTagController.GetTagsByCategory)
	categoryTags.Get("/:tagID", categoryTagController.GetTagByCategory)
}
