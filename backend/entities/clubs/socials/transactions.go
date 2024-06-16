package socials

import (
	"errors"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PutClubSocial(db *gorm.DB, social models.Social) (*models.Social, error) {
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "club_id"}, {Name: "type"}},
		DoUpdates: clause.AssignmentColumns([]string{"content"}),
	}).Create(&social).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrForeignKeyViolated) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}
	return &social, nil
}

func GetClubSocials(db *gorm.DB, clubID uuid.UUID) ([]models.Social, error) {
	db = cache.SetUseCache(db, true)

	var club models.Club
	if err := db.Preload("Social").First(&club, clubID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return club.Social, nil
}
