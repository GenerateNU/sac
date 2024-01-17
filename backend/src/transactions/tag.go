package transactions

import (
	"backend/src/models"

	"gorm.io/gorm"
)

func CreateTag(db *gorm.DB, tag models.Tag) (*models.Tag, error) {
	if err := db.Create(&tag).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}

func GetTag(db *gorm.DB, id uint) (models.Tag, error) {
	var tag models.Tag

	if err := db.First(&tag, id).Error; err != nil {
		return models.Tag{}, err
	}

	return tag, nil
}
