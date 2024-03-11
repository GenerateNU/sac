package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/src/models"

	"gorm.io/gorm"
)

func GetTagsByCategory(db *gorm.DB, categoryID uuid.UUID, limit int, page int) ([]models.Tag, *errors.Error) {
	var category models.Category

	if err := db.Where("id = ?", categoryID).First(&category).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.CategoryNotFound
		}
		return nil, &errors.FailedToGetCategory
	}

	var tags []models.Tag

	offset := (page - 1) * limit

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
