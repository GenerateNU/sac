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
