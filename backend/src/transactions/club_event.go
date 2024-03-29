package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubEvents(db *gorm.DB, clubID uuid.UUID, limit int, page int) ([]models.Event, *errors.Error) {
	var events []models.Event

	offset := (page - 1) * limit

	if err := db.Where("club_id = ?", clubID).Limit(limit).Offset(offset).Find(&events).Error; err != nil {
		return nil, &errors.FailedToGetEvents
	}

	return events, nil
}
