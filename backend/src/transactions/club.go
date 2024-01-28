package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"

	"gorm.io/gorm"
)

func GetClubs(db *gorm.DB, limit int, offset int) ([]models.Club, *errors.Error) {
	var clubs []models.Club
	result := db.Limit(limit).Offset(offset).Find(&clubs)
	if result.Error != nil {
		return nil, &errors.FailedToGetClubs
	}

	return clubs, nil
}

func CreateClub(db *gorm.DB, userId uint, club models.Club) (*models.Club, *errors.Error) {
	user, err := GetUser(db, userId)
	if err != nil {
		return nil, &errors.UserNotFound
	}

	tx := db.Begin()

	if err := tx.Create(&club).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	if err := tx.Model(&club).Association("Admin").Append(user); err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToCreateClub
	}

	return &club, nil
}

func GetClub(db *gorm.DB, id uint) (*models.Club, *errors.Error) {
	var club models.Club
	if err := db.First(&club, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToGetClub
		}
	}

	return &club, nil
}

func UpdateClub(db *gorm.DB, id uint, club models.Club) (*models.Club, *errors.Error) {
	result := db.Model(&club).Where("id = ?", id).Updates(club)
	if result.Error != nil {
		return nil, &errors.FailedToUpdateClub
	}

	return &club, nil
}

func DeleteClub(db *gorm.DB, id uint) *errors.Error {
	result := db.Delete(&models.Club{}, id)
	if result.Error != nil {
		return &errors.FailedToDeleteClub
	}

	return nil
}

func GetContacts(db *gorm.DB, limit int, offset int) ([]models.Contact, *errors.Error) {
	var contacts []models.Contact
	result := db.Limit(limit).Offset(offset).Find(&contacts)
	if result.Error != nil {
		return nil, &errors.FailedToGetContacts
	}

	return contacts, nil
}

func GetClubContacts(db *gorm.DB, id uint) ([]models.Contact, *errors.Error) {
	var club models.Club
	if err := db.Preload("Contact").First(&club, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToGetContacts
		}
	}

	if club.Contact == nil {
		return nil, &errors.FailedToGetContacts
	}

	return club.Contact, nil
}

func PutContact(db *gorm.DB, clubID uint, contact models.Contact) (*models.Contact, *errors.Error) {
	if clubID != contact.ClubID {
		return nil, &errors.FailedToUpdateContact
	}

	var club models.Club
	if err := db.Preload("Contact").First(&club, clubID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToGetContacts
		}
	}

	// determines if the inputted contact type exists for the current club
	// if it exists, save the contactID so we can update it
	contacts := club.Contact
	alreadyExists := false
	var contactID uint = 0
	for _, c := range contacts {
		if c.Type == contact.Type {
			contactID = c.ID
			alreadyExists = true
		}
	}

	if alreadyExists {
		// update
		result := db.Model(&contact).Where("id = ?", contactID).Updates(contact)
		if result.Error != nil {
			return nil, &errors.FailedToUpdateContact
		}
		return &contact, nil
	} else {
		// create
		tx := db.Begin()
		if err := tx.Create(&contact).Error; err != nil {
			tx.Rollback()
			return nil, &errors.FailedToCreateContact
		}
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return nil, &errors.FailedToCreateContact
		}
		return &contact, nil
	}
}

func DeleteContact(db *gorm.DB, id uint) *errors.Error {
	result := db.Delete(&models.Contact{}, id)
	if result.Error != nil {
		return &errors.FailedToDeleteClub
	}

	return nil
}
