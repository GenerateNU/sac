package series

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateEventSeries(db *gorm.DB, events []models.Event) ([]models.Event, error) {
	if err := db.Create(&events).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func GetEventSeries(db *gorm.DB, seriesID uuid.UUID) ([]models.Event, error) {
	var events []models.Event

	if err := db.Where("series_id = ?", seriesID).Find(&events).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return events, nil
}

func DeleteEventSeries(db *gorm.DB, seriesID uuid.UUID) error {
	if err := db.Where("series_id = ?", seriesID).Delete(&models.Event{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utilities.ErrNotFound
		}
		return err
	}

	return nil
}
