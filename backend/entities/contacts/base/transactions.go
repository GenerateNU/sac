package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetContacts(db *gorm.DB, limit int, page int) ([]models.Contact, error) {
	var contacts []models.Contact

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

func GetContact(db *gorm.DB, id uuid.UUID) (*models.Contact, error) {
	var contact models.Contact
	if err := db.First(&contact, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &contact, nil
}

func DeleteContact(db *gorm.DB, id uuid.UUID) error {
	if result := db.Delete(&models.Contact{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}
	return nil
}
