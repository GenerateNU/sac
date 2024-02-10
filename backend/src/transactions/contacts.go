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

func UpdateClub(db *gorm.DB, id uuid.UUID, club models.Club) (*models.Club, *errors.Error) {
	result := db.Model(&models.User{}).Where("id = ?", id).Updates(club)
	if result.Error != nil {
		if stdliberrors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToUpdateClub
		}
	}
	var existingClub models.Club

	err := db.First(&existingClub, id).Error
	if err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToCreateClub
		}
	}

	if err := db.Model(&existingClub).Updates(&club).Error; err != nil {
		return nil, &errors.FailedToUpdateUser
	}

	return &existingClub, nil
}

func DeleteClub(db *gorm.DB, id uuid.UUID) *errors.Error {
	if result := db.Delete(&models.Club{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.ClubNotFound
		} else {
			return &errors.FailedToDeleteClub
		}
	}

	return nil
}
