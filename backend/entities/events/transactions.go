package events

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetEvent(db *gorm.DB, eventID uuid.UUID, preloads ...transactions.OptionalQuery) (*models.Event, error) {
	var event models.Event

	query := db

	for _, preload := range preloads {
		query = preload(query)
	}
	if err := query.First(&event, eventID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &event, nil
}

func GetEventHostAdminIDs(db *gorm.DB, eventID uuid.UUID) ([]uuid.UUID, error) {
	event, err := GetEvent(db, eventID)
	if err != nil {
		return nil, err
	}

	return clubs.GetAdminIDs(db, *event.Host)
}
