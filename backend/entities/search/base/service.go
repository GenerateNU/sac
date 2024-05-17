package base

import (
	"github.com/GenerateNU/sac/backend/search"

	"github.com/GenerateNU/sac/backend/types"
)

type SearchServiceInterface interface {
	SearchClubs(query search.ClubSearchQuery) (*search.ClubSearchResult, error)
}

type SearchService struct {
	types.ServiceParams
}

func NewSearchService(serviceParams types.ServiceParams) SearchServiceInterface {
	return &SearchService{serviceParams}
}

func (s *SearchService) SearchClubs(query search.ClubSearchQuery) (*search.ClubSearchResult, error) {
	return SearchClubs(s.DB, query)
}
