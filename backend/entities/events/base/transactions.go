package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"

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

func CreateEvent(db *gorm.DB, event models.Event) (*models.Event, error) {
	if err := db.Create(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func UpdateEvent(db *gorm.DB, id uuid.UUID, event models.Event) (*models.Event, error) {

	if result := db.Model(&models.Event{}).Where("id = ?", id).Updates(event); result.RowsAffected == 0 {
		if result.Error == nil {
			return nil, utilities.ErrNotFound
		}
		return nil, result.Error
	}

	return &event, nil
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
