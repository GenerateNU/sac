package tags

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetTag(db *gorm.DB, tagID uuid.UUID) (*models.Tag, *errors.Error) {
	var tag models.Tag
	if err := db.Where("id = ?", tagID).First(&tag).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.TagNotFound
		} else {
			return nil, &errors.FailedToGetTag
		}
	}

	return &tag, nil
}

func GetTagsByIDs(db *gorm.DB, selectedTagIDs []uuid.UUID) ([]models.Tag, *errors.Error) {
	if len(selectedTagIDs) != 0 {
		var tags []models.Tag
		if err := db.Model(models.Tag{}).Where("id IN ?", selectedTagIDs).Find(&tags).Error; err != nil {
			return nil, &errors.FailedToGetTag
		}

		return tags, nil
	}
	return []models.Tag{}, nil
}
