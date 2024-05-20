package extractors

import (
	"github.com/GenerateNU/sac/backend/entities/events"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetSeriesHost(db *gorm.DB, seriesID uuid.UUID) (*uuid.UUID, error) {
	event, err := events.GetFirstEventInSeries(db, seriesID)
	if err != nil {
		return nil, err
	}

	return event.Host, nil
}
