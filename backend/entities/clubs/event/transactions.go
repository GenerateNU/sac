package event

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubEvents(db *gorm.DB, clubID uuid.UUID, limit int, page int) ([]models.Event, error) {
	var events []models.Event

	offset := (page - 1) * limit

	if err := db.Where("club_id = ?", clubID).Limit(limit).Offset(offset).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}
