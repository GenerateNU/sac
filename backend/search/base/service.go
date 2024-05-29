package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	search_types "github.com/GenerateNU/sac/backend/search/types"

	"github.com/GenerateNU/sac/backend/types"
)

type SearchServiceInterface interface {
	SearchClubs(query search_types.ClubSearchRequest) (*search_types.SearchResult[models.Club], error)
	SearchEvents(query search_types.EventSearchRequest) (*search_types.SearchResult[models.Event], error)
}

type SearchService struct {
	types.ServiceParams
}

func NewSearchService(serviceParams types.ServiceParams) SearchServiceInterface {
	return &SearchService{serviceParams}
}

func (s *SearchService) SearchClubs(query search_types.ClubSearchRequest) (*search_types.SearchResult[models.Club], error) {
	return Search[models.Club](s.DB, &query)
}

func (s *SearchService) SearchEvents(query search_types.EventSearchRequest) (*search_types.SearchResult[models.Event], error) {
	return Search[models.Event](s.DB, &query)
}
