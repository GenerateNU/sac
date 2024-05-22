package base

import (
	"errors"
	"log/slog"

	"github.com/GenerateNU/sac/backend/entities/events"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/garrettladley/fiberpaginate"

	search "github.com/GenerateNU/sac/backend/search/base"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetEvents(db *gorm.DB, pageInfo fiberpaginate.PageInfo) ([]models.Event, error) {
	var events []models.Event
	if err := db.Scopes(utilities.IntoScope(pageInfo, db)).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func CreateEvent(db *gorm.DB, event models.Event) (*models.Event, error) {
	if err := db.Create(&event).Error; err != nil {
		return nil, err
	}

	err := search.Upsert[models.Event](db, "events", event.ID.String(), &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func UpdateEvent(db *gorm.DB, id uuid.UUID, event models.Event, updateAllInSeries bool) ([]models.Event, error) {
	if updateAllInSeries {
		return updateEventSeries(db, id, event)
	} else {
		return updateEvent(db, id, event)
	}
}

func updateEvent(db *gorm.DB, id uuid.UUID, event models.Event) ([]models.Event, error) {
	var resultingEvent models.Event
	if result := db.Model(&resultingEvent).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&event); result.RowsAffected == 0 {
		if result.Error == nil {
			return nil, utilities.ErrNotFound
		}
		return nil, result.Error
	}

	err := search.Upsert[models.Event](db, "events", id.String(), &event)
	if err != nil {
		return nil, err
	}

	return []models.Event{resultingEvent}, nil
}

func updateEventSeries(db *gorm.DB, id uuid.UUID, event models.Event) ([]models.Event, error) {
	targetEvent, err := events.GetEvent(db, id)
	if err != nil {
		return nil, err
	}

	if targetEvent.SeriesID == nil {
		slog.Error("event is not part of a series", "eventID", id)
		return nil, utilities.BadRequest(errors.New("event is not part of a series"))
	}

	var resultingEvents []models.Event
	if result := db.Model(&resultingEvents).Clauses(clause.Returning{}).Where("series_id = ?", targetEvent.SeriesID).Updates(&event); result.RowsAffected == 0 {
		if result.Error == nil {
			return nil, utilities.ErrNotFound
		}
		return nil, result.Error
	}

	return resultingEvents, nil
}

func DeleteEvent(db *gorm.DB, id uuid.UUID) error {
	if result := db.Delete(&models.Event{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}

	err := search.Delete(db, "events", id.String())
	if err != nil {
		return err
	}

	return nil
}
