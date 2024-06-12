package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetLeadership(db *gorm.DB, pageInfo fiberpaginate.PageInfo) ([]models.Leader, error) {
	db = cache.SetCache(db, true)

	var leadership []models.Leader
	result := db.Preload("PhotoFile").Scopes(utilities.IntoScope(pageInfo, db)).Find(&leadership)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, result.Error
	}

	return leadership, nil
}

func GetLeader(db *gorm.DB, id uuid.UUID) (*models.Leader, error) {
	db = cache.SetCache(db, true)

	var leader models.Leader
	if err := db.Preload("PhotoFile").First(&leader, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &leader, nil
}
