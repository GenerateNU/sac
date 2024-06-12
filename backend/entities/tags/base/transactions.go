package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/tags"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateTag(db *gorm.DB, tag models.Tag) (*models.Tag, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var category models.Category
	if err := tx.Where("id = ?", tag.CategoryID).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, utilities.ErrNotFound
		} else {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Create(&tag).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &tag, tx.Commit().Error
}

func GetTags(db *gorm.DB) ([]models.Tag, error) {
	db = cache.SetCache(db, true)

	var tags []models.Tag

	if err := db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func GetTagClubs(db *gorm.DB, id uuid.UUID) ([]models.Club, error) {
	db = cache.SetCache(db, true)

	tag, err := tags.GetTag(db, id)
	if err != nil {
		return nil, err
	}

	var clubs []models.Club
	if err := db.Model(&tag).Association("Club").Find(&clubs); err != nil {
		return nil, err
	}

	return clubs, nil
}

func UpdateTag(db *gorm.DB, id uuid.UUID, tag models.Tag) (*models.Tag, error) {
	var resultingTag models.Tag
	if err := db.Model(&resultingTag).Clauses(clause.Returning{}).Where("id = ?", id).Updates(tag).First(&tag, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &resultingTag, nil
}

func DeleteTag(db *gorm.DB, tagID uuid.UUID) error {
	if result := db.Where("id = ?", tagID).Delete(&models.Tag{}); result.RowsAffected == 0 {
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}

	return nil
}
