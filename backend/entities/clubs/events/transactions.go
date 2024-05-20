package events

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubEvents(db *gorm.DB, clubID uuid.UUID, pageInfo fiberpaginate.PageInfo) ([]models.Event, error) {
	var events []models.Event

	if err := db.Where("club_id = ?", clubID).Scopes(utilities.IntoScope(pageInfo, db)).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}
