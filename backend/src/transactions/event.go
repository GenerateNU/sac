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

func CreateEvent(db *gorm.DB, userId uuid.UUID, club models.Event) (*models.Event, *errors.Error) {
	tx := db.Begin()

	if err := tx.Create(&club).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	return &club, nil
}