package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPointOfContacts(db *gorm.DB, limit int, page int) ([]models.PointOfContact, error) {
	var pointOfContacts []models.PointOfContact

	offset := (page - 1) * limit

	result := db.Preload("PhotoFile").Limit(limit).Offset(offset).Find(&pointOfContacts)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, result.Error
	}

	return pointOfContacts, nil
}

func GetPointOfContact(db *gorm.DB, id uuid.UUID) (*models.PointOfContact, error) {
	var pointOfContact models.PointOfContact

	if err := db.Preload("PhotoFile").First(&pointOfContact, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &pointOfContact, nil
}
