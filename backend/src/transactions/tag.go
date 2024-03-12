package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/src/models"

	"gorm.io/gorm"
)

func CreateTag(db *gorm.DB, tag models.Tag) (*models.Tag, *errors.Error) {
	tx := db.Begin()

	var category models.Category
	if err := tx.Where("id = ?", tag.CategoryID).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			tx.Rollback()
			return nil, &errors.CategoryNotFound
		} else {
			tx.Rollback()
			return nil, &errors.InternalServerError
		}
	}

	if err := tx.Create(&tag).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateTag
	}

	tx.Commit()

	return &tag, nil
}

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

func GetTags(db *gorm.DB, limit int, page int) ([]models.Tag, *errors.Error) {
	var tags []models.Tag

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Find(&tags).Error; err != nil {
		return nil, &errors.FailedToGetTags
	}

	return tags, nil
}

func UpdateTag(db *gorm.DB, id uuid.UUID, tag models.Tag) (*models.Tag, *errors.Error) {
	if err := db.Model(&models.Tag{}).Where("id = ?", id).Updates(tag).First(&tag, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.TagNotFound
		} else {
			return nil, &errors.FailedToUpdateTag
		}
	}

	return &tag, nil
}

func DeleteTag(db *gorm.DB, tagID uuid.UUID) *errors.Error {
	if result := db.Where("id = ?", tagID).Delete(&models.Tag{}); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.TagNotFound
		} else {
			return &errors.FailedToDeleteTag
		}
	}

	return nil
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

// Get clubs for a tag
func GetTagClubs(db *gorm.DB, id uuid.UUID) ([]models.Club, *errors.Error) {
	var clubs []models.Club

	tag, err := GetTag(db, id)
	if err != nil {
		return nil, &errors.ClubNotFound
	}

	if err := db.Model(&tag).Association("Club").Find(&clubs); err != nil {
		return nil, &errors.FailedToGetTag
	}
	return clubs, nil
}
