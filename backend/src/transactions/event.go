package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
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

// given an eventID, gets the event
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

// given an event ID, get the ID of the series that the event is a part of
func GetSeriesId(db *gorm.DB, eventId uuid.UUID) (*uuid.UUID, *errors.Error) {
	// TODO maybe check if event exists first

	var SeriesID uuid.UUID
	if err := db.Model(&models.Event_Series{}).Where("event_id = ?", eventId).Select("series_id").Find(&SeriesID).Error; err != nil {
		return nil, &errors.FailedToGetEventSeries
	}

	if SeriesID.String() == "" {
		return nil, &errors.SeriesNotFound
	}

	return &SeriesID, nil
}

// given an Event ID, finds all the events in the series that the event is a part of
func GetSeriesByEventId(db *gorm.DB, id uuid.UUID) ([]models.Event, *errors.Error) {

	SeriesID, err := GetSeriesId(db, id)
	if err != nil {
		return nil, err
	}
	events, err := GetSeriesById(db, *SeriesID)
	if err != nil {
		return nil, &errors.FailedToGetEventSeries
	}
	return events, nil
}

// given a seriesID, gets all events in that series
func GetSeriesById(db *gorm.DB, id uuid.UUID) ([]models.Event, *errors.Error) {
	var series models.Series
	if err := db.Preload("Events").Find(&series, id).Error; err != nil {
		return nil, &errors.FailedToGetEventSeries
	}

	return series.Events, nil
}

// given a clubId, finds all events of the club
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

// given an event id, delete all events in the event's series
func DeleteSeriesByEventId(db *gorm.DB, id uuid.UUID) *errors.Error {

	SeriesID, err := GetSeriesId(db, id)
	if err != nil {
		return err
	}

	if err:= DeleteSeriesById(db, *SeriesID); err !=nil {
		return err
	}


	return nil
}

func DeleteSeriesById(db *gorm.DB, SeriesID uuid.UUID) *errors.Error{
	tx:=db.Begin()
	var eventIDs uuid.UUIDs
	if err := tx.Model(&models.Event_Series{}).Select("event_id").Where("series_id = (?)", SeriesID).Find(&eventIDs).Error; err != nil {
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