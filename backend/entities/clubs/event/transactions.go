package event

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubEvents(db *gorm.DB, clubID uuid.UUID, pagination middleware.Pagination) ([]models.Event, error) {
	var events []models.Event

	if err := db.Where("club_id = ?", clubID).Scopes(pagination.IntoScope(db)).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}
