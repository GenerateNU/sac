package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func GetEvents(db *gorm.DB, limit int, page int) ([]models.Event, error) {
	var events []models.Event

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func GetEvent(db *gorm.DB, eventID uuid.UUID, preloads ...transactions.OptionalQuery) ([]models.Event, error) {
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

func GetSeriesID(db *gorm.DB, eventID uuid.UUID) (*uuid.UUID, error) {
	var event models.Event

	if err := db.First(&event, eventID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	var seriesID string

	if err := db.Model(&models.EventSeries{}).Where("event_id = ?", event.ID).Select("series_id").Find(&seriesID).Error; err != nil {
		return nil, err
	}

	parsedSeriesID, err := uuid.Parse(seriesID)
	if err != nil {
		return nil, err
	}

	return &parsedSeriesID, nil
}

func GetSeriesByEventID(db *gorm.DB, id uuid.UUID) ([]models.Event, error) {
	seriesID, err := GetSeriesID(db, id)
	if err != nil {
		return nil, err
	}
	events, err := GetSeriesByID(db, *seriesID)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func GetSeriesByID(db *gorm.DB, id uuid.UUID) ([]models.Event, error) {
	var series models.Series

	if err := db.Preload("Events").Find(&series, id).Error; err != nil {
		return nil, err
	}

	return series.Events, nil
}

func CreateEvent(db *gorm.DB, event models.Event) ([]models.Event, error) {
	tx := db.Begin()

	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return []models.Event{event}, nil
}

func CreateEventSeries(db *gorm.DB, series models.Series) ([]models.Event, error) {
	if err := db.Create(&series).Error; err != nil {
		return nil, err
	}

	return series.Events, nil
}

func UpdateEvent(db *gorm.DB, id uuid.UUID, event models.Event) ([]models.Event, error) {
	if err := db.Model(&models.Event{}).Where("id = ?", id).Updates(event).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	var existingEvent models.Event

	if err := db.First(&existingEvent, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	if err := db.Model(&existingEvent).Updates(&event).Error; err != nil {
		return nil, err
	}

	return []models.Event{existingEvent}, nil
}

func UpdateSeries(db *gorm.DB, seriesID uuid.UUID, series models.Series) ([]models.Event, error) {
	err := DeleteSeriesByID(db, seriesID)
	if err != nil {
		return nil, err
	}

	series.ID = seriesID

	events, err := CreateEventSeries(db, series)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func UpdateSeriesByEventID(db *gorm.DB, eventID uuid.UUID, series models.Series) ([]models.Event, error) {
	seriesID, err := GetSeriesID(db, eventID)
	if err != nil {
		return nil, err
	}

	events, err := UpdateSeries(db, *seriesID, series)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func DeleteEvent(db *gorm.DB, id uuid.UUID) error {
	if result := db.Delete(&models.Event{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}

	return nil
}

func DeleteSeriesByEventID(db *gorm.DB, eventID uuid.UUID) error {
	seriesID, err := GetSeriesID(db, eventID)
	if err != nil {
		return err
	}

	if err := DeleteSeriesByID(db, *seriesID); err != nil {
		return err
	}

	return nil
}

func DeleteSeriesByID(db *gorm.DB, seriesID uuid.UUID) error {
	tx := db.Begin()

	var eventIDs uuid.UUIDs

	if err := tx.Model(&models.EventSeries{}).Select("event_id").Where("series_id = (?)", seriesID).Find(&eventIDs).Error; err != nil {
		tx.Rollback()
		return err
	} else if len(eventIDs) < 1 {
		tx.Rollback()
		return utilities.ErrNotFound
	}

	if result := tx.Delete(&models.Event{}, eventIDs); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}

	if result := tx.Delete(&models.Series{}, seriesID); result.RowsAffected == 0 {
		tx.Rollback()
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func GetHostsByEventID(db *gorm.DB, eventID uuid.UUID) ([]models.Club, error) {
	// Get the club that hosts the event
	event, err := GetEvent(db, eventID)
	if err != nil {
		return nil, err
	}

	clubs, err := clubs.GetClub(db, *event[0].Host, transactions.PreloadEvent())
	if err != nil {
		return nil, err
	}

	return []models.Club{*clubs}, nil
}

func GetTagsByEventID(db *gorm.DB, eventID uuid.UUID) ([]models.Tag, error) {
	event, err := GetEvent(db, eventID, transactions.PreloadTag())
	if err != nil {
		return nil, err
	}

	return event[0].Tag, nil
}
