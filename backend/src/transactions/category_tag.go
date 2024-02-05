package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/src/models"

	"gorm.io/gorm"
)

func GetTagsByCategory(db *gorm.DB, categoryID uuid.UUID, limit int, offset int) ([]models.Tag, *errors.Error) {
	var tags []models.Tag
	if err := db.Where("category_id = ?", categoryID).Limit(limit).Offset(offset).Find(&tags).Error; err != nil {
		return nil, &errors.FailedToGetTags
	}

	return tags, nil
}

func GetTagByCategory(db *gorm.DB, categoryID uuid.UUID, tagID uuid.UUID) (*models.Tag, *errors.Error) {
	var tag models.Tag
	if err := db.Where("category_id = ? AND id = ?", categoryID, tagID).First(&tag).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.TagNotFound
		} else {
			return nil, &errors.FailedToGetTag
		}
	}

	return &tag, nil
}
