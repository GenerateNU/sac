package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func GetEvents(db *gorm.DB, limit int, offset int) ([]models.Event, *errors.Error) {
	var events []models.Event
	result := db.Limit(limit).Offset(offset).Find(&events)
	if result.Error != nil {
		return nil, &errors.FailedToGetEvents
	}

	return events, nil
}

// TODO get events by club id
func GetEvent(db *gorm.DB, id uuid.UUID) (*models.Event, *errors.Error) {
	var event models.Event
	if err := db.First(&event, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.EventNotFound
		} else {
			return nil, &errors.FailedToGetEvent
		}
	}

	return &event, nil
}

func GetSeriesByEventId(db *gorm.DB, id uuid.UUID) ([]models.Event, *errors.Error) {
	var sid string
	tx := db.Begin()
	if err := tx.Model(&models.Event_Series{}).Where("event_id = ?", id).Select("series_id").Find(&sid).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToGetEventSeries
	}

	if sid == "" {
		return nil, &errors.SeriesNotFound
	}
	sidAsUUID, err := utilities.ValidateID(sid)
	if err != nil {
		return nil, err
	}
	events, err := GetSeriesBySeriesId(tx, *sidAsUUID)
	if err != nil {
		tx.Rollback()
		return nil, &errors.FailedToGetEventSeries
	}
	tx.Commit()
	return events, nil
}

func GetSeriesBySeriesId(db *gorm.DB, id uuid.UUID) ([]models.Event, *errors.Error) {
	var series models.Series
	if err := db.Preload("Events").Find(&series, id).Error; err != nil {
		return nil, &errors.FailedToGetEventSeries
	}

	return series.Events, nil
}

func GetClubEvents(db *gorm.DB, id uuid.UUID) ([]models.Event, *errors.Error) {
	var events []models.Event

	if err := db.Where("club_id = ?", id).Find(&events).Error; err != nil {
		return nil, &errors.FailedToGetClubEvents
	}

	return events, nil
}

func CreateEvent(db *gorm.DB, event models.Event) (*models.Event, *errors.Error) {
	tx := db.Begin()

	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateEvent
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateEvent
	}

	return &event, nil
}

func CreateEventSeries(db *gorm.DB, series models.Series) ([]models.Event, *errors.Error) {
	tx := db.Begin()

	// if err := tx.Create(&events).Error; err != nil {
	// 	tx.Rollback()
	// 	return nil, &errors.FailedToCreateEventSeries
	// }

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

func UpdateEvent(db *gorm.DB, id uuid.UUID, event models.UpdateEventRequestBody) (*models.Event, *errors.Error) {
	result := db.Model(&models.Event{}).Where("id = ?", id).Updates(event)
	if result.Error != nil {
		if stdliberrors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToUpdateEvent
		}
	}
	var existingEvent models.Event

	err := db.First(&existingEvent, id).Error
	if err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.EventNotFound
		} else {
			return nil, &errors.FailedToCreateEvent
		}
	}

	if err := db.Model(&existingEvent).Updates(&event).Error; err != nil {
		return nil, &errors.FailedToUpdateUser
	}

	return &existingEvent, nil
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

// delete series of an event
// TODO factor out getting series by event id
func DeleteEventSeries(db *gorm.DB, id uuid.UUID) *errors.Error {

	var sid string
	tx := db.Begin()
	if err := tx.Model(&models.Event_Series{}).Where("event_id = ?", id).Select("series_id").Find(&sid).Error; err != nil {
		tx.Rollback()
		return &errors.FailedToGetEventSeries
	}

	if sid == "" {
		return &errors.SeriesNotFound
	}
	uid, err := utilities.ValidateID(sid)

	if err != nil {
		tx.Rollback()
		return err
	}

	var eventIDs uuid.UUIDs
	if err := tx.Model(&models.Event_Series{}).Select("event_id").Where("series_id = (?)", uid).Find(&eventIDs).Error; err != nil {
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
	tx.Commit()

	return nil
}
