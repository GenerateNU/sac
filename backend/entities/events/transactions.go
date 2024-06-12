package events

import (
	"errors"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetEvents(db *gorm.DB, pageInfo fiberpaginate.PageInfo) ([]models.Event, error) {
	var events []models.Event

	db = cache.SetUseCache(db, true)

	if err := db.Scopes(utilities.IntoScope(pageInfo, db)).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func GetEvent(db *gorm.DB, eventID uuid.UUID, preloads ...transactions.OptionalQuery) (*models.Event, error) {
	query := cache.SetUseCache(db, true)

	for _, preload := range preloads {
		query = preload(query)
	}

	var event models.Event
	if err := query.First(&event, eventID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &event, nil
}

func GetEventHostAdminIDs(db *gorm.DB, eventID uuid.UUID) ([]uuid.UUID, error) {
	db = cache.SetUseCache(db, true)

	event, err := GetEvent(db, eventID)
	if err != nil {
		return nil, err
	}

	return clubs.GetAdminIDs(db, *event.Host)
}

func GetFirstEventInSeries(db *gorm.DB, seriesID uuid.UUID) (*models.Event, error) {
	db = cache.SetUseCache(db, true)

	var event models.Event
	if err := db.Where("series_id = ?", seriesID).First(&event).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &event, nil
}
