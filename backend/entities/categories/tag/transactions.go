package tag

import (
	"errors"

	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"

	"gorm.io/gorm"
)

func GetTagsByCategory(db *gorm.DB, categoryID uuid.UUID, limit int, page int) ([]models.Tag, error) {
	var category models.Category

	if err := db.Where("id = ?", categoryID).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	var tags []models.Tag

	offset := (page - 1) * limit

	if err := db.Where("category_id = ?", categoryID).Limit(limit).Offset(offset).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func GetTagByCategory(db *gorm.DB, categoryID uuid.UUID, tagID uuid.UUID) (*models.Tag, error) {
	var tag models.Tag
	if err := db.Where("category_id = ? AND id = ?", categoryID, tagID).First(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &tag, nil
}
