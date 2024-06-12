package series

import (
	"errors"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateEventSeries(db *gorm.DB, seriesID uuid.UUID, events []models.Event) (*Series, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	series, err := createSeries(tx, seriesID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(&events).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	series.Events = events

	return series, tx.Commit().Error
}

func GetEventSeries(db *gorm.DB, seriesID uuid.UUID) ([]models.Event, error) {
	db = cache.SetCache(db, true)

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
	if err := db.Where("id = ?", seriesID).Delete(&Series{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utilities.ErrNotFound
		}

		return err
	}

	return nil
}

func createSeries(db *gorm.DB, seriesID uuid.UUID) (*Series, error) {
	series := &Series{
		Model: models.Model{
			ID: seriesID,
		},
	}

	if err := db.Create(series).Error; err != nil {
		return nil, err
	}

	return series, nil
}
