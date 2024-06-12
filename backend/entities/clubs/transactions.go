package clubs

import (
	"errors"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClub(db *gorm.DB, id uuid.UUID, preloads ...transactions.OptionalQuery) (*models.Club, error) {
	query := cache.SetUseCache(db, true)

	for _, preload := range preloads {
		query = preload(query)
	}

	var club models.Club
	if err := query.First(&club, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &club, nil
}

func GetAdminIDs(db *gorm.DB, clubID uuid.UUID) ([]uuid.UUID, error) {
	db = cache.SetUseCache(db, true)

	var adminIDs []models.Membership
	if err := db.Where("club_id = ? AND membership_type = ?", clubID, models.MembershipTypeAdmin).Find(&adminIDs).Error; err != nil {
		return nil, err
	}

	adminUUIDs := make([]uuid.UUID, len(adminIDs))
	for _, adminID := range adminIDs {
		adminUUIDs = append(adminUUIDs, adminID.ClubID)
	}

	return adminUUIDs, nil
}
