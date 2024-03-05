package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubEvents(db *gorm.DB, clubID uuid.UUID, limit int, offset int) ([]models.Event, *errors.Error) {
	club, err := GetClub(db, clubID)
	if err != nil {
		return nil, err
	}

	var events []models.Event

	if err := db.Model(&club).Limit(limit).Offset(offset).Association("Event").Find(&events); err != nil {
		return nil, &errors.FailedToGetClubEvents
	}

	return events, nil
}
