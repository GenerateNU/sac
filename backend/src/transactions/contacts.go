package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetContacts(db *gorm.DB, limit int, offset int) ([]models.Contact, *errors.Error) {
	var contacts []models.Contact
	result := db.Limit(limit).Offset(offset).Find(&contacts)
	if result.Error != nil {
		return nil, &errors.FailedToGetContacts
	}

	return contacts, nil
}

func GetContact(db *gorm.DB, id uuid.UUID) (*models.Contact, *errors.Error) {
	var contact models.Contact
	if err := db.First(&contact, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ContactNotFound
		} else {
			return nil, &errors.FailedToGetContact
		}
	}

	return &contact, nil
}

func DeleteContact(db *gorm.DB, id uuid.UUID) *errors.Error {
	if result := db.Delete(&models.Contact{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.ContactNotFound
		} else {
			return &errors.FailedToDeleteContact
		}
	}
	return nil
}
