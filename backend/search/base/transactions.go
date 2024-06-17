package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/search/types"
	"gorm.io/gorm"
)

func SearchClubs(db *gorm.DB, query *types.ClubSearchRequest) (*types.SearchResult[models.Club], error) {
	return nil, nil
}

func SearchEvents(db *gorm.DB, query *types.EventSearchRequest) (*types.SearchResult[models.Event], error) {
	return nil, nil
}
