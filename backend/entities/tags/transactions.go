package tags

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetTag(db *gorm.DB, tagID uuid.UUID) (*models.Tag, error) {
	var tag models.Tag
	if err := db.Where("id = ?", tagID).First(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &tag, nil
}

func GetTagsByIDs(db *gorm.DB, selectedTagIDs []uuid.UUID) ([]models.Tag, error) {
	if len(selectedTagIDs) != 0 {
		var tags []models.Tag
		if err := db.Model(models.Tag{}).Where("id IN ?", selectedTagIDs).Find(&tags).Error; err != nil {
			return nil, err
		}

		return tags, nil
	}
	return []models.Tag{}, nil
}
