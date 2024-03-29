package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func GetEvents(db *gorm.DB, limit int, page int) ([]models.Event, *errors.Error) {
	var events []models.Event

	offset := (page - 1) * limit

	if db.Limit(limit).Offset(offset).Find(&events).Error != nil {
		return nil, &errors.FailedToGetEvents
	}

	return events, nil
}

func GetEvent(db *gorm.DB, eventID uuid.UUID) ([]models.Event, *errors.Error) {
	var event models.Event

	if err := db.First(&event, eventID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.EventNotFound
		} else {
			return nil, &errors.FailedToGetEvent
		}
	}

	return []models.Event{event}, nil
}

func GetSeriesID(db *gorm.DB, eventID uuid.UUID) (*uuid.UUID, *errors.Error) {
	var event models.Event

	if err := db.First(&event, eventID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.EventNotFound
		} else {
			return nil, &errors.FailedToGetEvent
		}
	}

	var seriesID string

	if err := db.Model(&models.EventSeries{}).Where("event_id = ?", event.ID).Select("series_id").Find(&seriesID).Error; err != nil {
		return nil, &errors.FailedToGetEventSeries
	}

	parsedSeriesID, err := uuid.Parse(seriesID)
	if err != nil {
		return nil, &errors.FailedToGetEventSeries
	}

	return &parsedSeriesID, nil
}

func GetSeriesByEventID(db *gorm.DB, id uuid.UUID) ([]models.Event, *errors.Error) {
	seriesID, err := GetSeriesID(db, id)
	if err != nil {
		return nil, err
	}
	events, err := GetSeriesByID(db, *seriesID)
	if err != nil {
		return nil, &errors.FailedToGetEventSeries
	}
	return events, nil
}

func GetSeriesByID(db *gorm.DB, id uuid.UUID) ([]models.Event, *errors.Error) {
	var series models.Series

	if err := db.Preload("Events").Find(&series, id).Error; err != nil {
		return nil, &errors.FailedToGetEventSeries
	}

	return series.Events, nil
}

func CreateEvent(db *gorm.DB, event models.Event) ([]models.Event, *errors.Error) {
	tx := db.Begin()

	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateEvent
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateEvent
	}

	return []models.Event{event}, nil
}

func CreateEventSeries(db *gorm.DB, series models.Series) ([]models.Event, *errors.Error) {
	tx := db.Begin()

	if err := tx.Create(&series).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateEventSeries
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateEventSeries
	}

	return series.Events, nil
}

func UpdateEvent(db *gorm.DB, id uuid.UUID, event models.Event) ([]models.Event, *errors.Error) {
	if err := db.Model(&models.Event{}).Where("id = ?", id).Updates(event).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToUpdateEvent
		}
	}

	var existingEvent models.Event

	if err := db.First(&existingEvent, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.EventNotFound
		} else {
			return nil, &errors.FailedToCreateEvent
		}
	}

	if err := db.Model(&existingEvent).Updates(&event).Error; err != nil {
		return nil, &errors.FailedToUpdateUser
	}

	return []models.Event{existingEvent}, nil
}

func UpdateSeries(db *gorm.DB, seriesID uuid.UUID, series models.Series) ([]models.Event, *errors.Error) {
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

func UpdateSeriesByEventID(db *gorm.DB, eventID uuid.UUID, series models.Series) ([]models.Event, *errors.Error) {
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

func DeleteEvent(db *gorm.DB, id uuid.UUID) *errors.Error {
	if result := db.Delete(&models.Event{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.EventNotFound
		} else {
			return &errors.FailedToDeleteEvent
		}
	}

	return nil
}

func DeleteSeriesByEventID(db *gorm.DB, eventID uuid.UUID) *errors.Error {
	seriesID, err := GetSeriesID(db, eventID)
	if err != nil {
		return err
	}

	if err := DeleteSeriesByID(db, *seriesID); err != nil {
		return err
	}

	return nil
}

func DeleteSeriesByID(db *gorm.DB, seriesID uuid.UUID) *errors.Error {
	tx := db.Begin()

	var eventIDs uuid.UUIDs

	if err := tx.Model(&models.EventSeries{}).Select("event_id").Where("series_id = (?)", seriesID).Find(&eventIDs).Error; err != nil {
		tx.Rollback()
		return &errors.FailedToDeleteSeries
	} else if len(eventIDs) < 1 {
		tx.Rollback()
		return &errors.SeriesNotFound
	}

	if result := tx.Delete(&models.Event{}, eventIDs); result.RowsAffected == 0 {
		tx.Rollback()
		return &errors.FailedToDeleteSeries
	}

	if result := tx.Delete(&models.Series{}, seriesID); result.RowsAffected == 0 {
		tx.Rollback()
		return &errors.FailedToDeleteSeries
	}

	tx.Commit()

	return nil
}
