package base

import (
	"github.com/GenerateNU/sac/backend/search"

	"github.com/GenerateNU/sac/backend/types"
)

type SearchServiceInterface interface {
	SearchClubs(query search.ClubSearchRequest) (*search.ClubSearchResult, error)
	SearchEvents(query search.EventSearchRequest) (*search.EventSearchResult, error)
}

type SearchService struct {
	types.ServiceParams
}

func NewSearchService(serviceParams types.ServiceParams) SearchServiceInterface {
	return &SearchService{serviceParams}
}

func (s *SearchService) SearchClubs(query search.ClubSearchRequest) (*search.ClubSearchResult, error) {
	return SearchClubs(s.DB, query)
}

func (s *SearchService) SearchEvents(query search.EventSearchRequest) (*search.EventSearchResult, error) {
	return SearchEvents(s.DB, query)
}
