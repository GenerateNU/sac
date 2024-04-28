package base

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPointOfContacts(db *gorm.DB, limit int, page int) ([]models.PointOfContact, *errors.Error) {
	var pointOfContacts []models.PointOfContact

	offset := (page - 1) * limit

	result := db.Preload("PhotoFile").Limit(limit).Offset(offset).Find(&pointOfContacts)
	if result.Error != nil {
		if stdliberrors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &errors.PointOfContactsNotFound
		} else {
			return nil, &errors.FailedToGetPointOfContacts
		}
	}

	return pointOfContacts, nil
}

func GetPointOfContact(db *gorm.DB, id uuid.UUID) (*models.PointOfContact, *errors.Error) {
	var pointOfContact models.PointOfContact

	if err := db.Preload("PhotoFile").First(&pointOfContact, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.PointOfContactNotFound
		} else {
			return nil, &errors.FailedToGetPointOfContact
		}
	}

	return &pointOfContact, nil
}
