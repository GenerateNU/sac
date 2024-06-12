package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetSocials(db *gorm.DB, pageInfo fiberpaginate.PageInfo) ([]models.Social, error) {
	db = cache.SetUseCache(db, true)

	var socials []models.Social
	if err := db.Scopes(utilities.IntoScope(pageInfo, db)).Find(&socials).Error; err != nil {
		return nil, err
	}

	return socials, nil
}

func GetSocial(db *gorm.DB, id uuid.UUID) (*models.Social, error) {
	db = cache.SetUseCache(db, true)

	var social models.Social
	if err := db.First(&social, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &social, nil
}

func DeleteSocial(db *gorm.DB, id uuid.UUID) error {
	if result := db.Delete(&models.Social{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}
	return nil
}
