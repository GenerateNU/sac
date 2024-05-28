package base

import (
	"github.com/GenerateNU/sac/backend/types"
)

func SearchRoutes(params types.RouteParams) {
	searchService := NewSearchService(params.ServiceParams)
	searchController := NewSearchController(searchService)

	// api/v1/search/*
	search := params.Router.Group("/search")

	search.Get("/clubs", searchController.SearchClubs)
	search.Get("/events", searchController.SearchEvents)
}
