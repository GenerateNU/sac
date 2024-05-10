package tag

import (
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

func CategoryTag(categoryParams types.RouteParams) {
	categoryTagController := NewCategoryTagController(NewCategoryTagService(categoryParams.ServiceParams))

	// api/v1/categories/:categoryID/tags/*
	categoryTags := categoryParams.Router.Group("/tags")

	categoryTags.Get("/", utilities.Make(categoryTagController.GetTagsByCategory))
	categoryTags.Get("/:tagID", utilities.Make(categoryTagController.GetTagByCategory))
}
