package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateClubPointOfContact(db *gorm.DB, clubID uuid.UUID, pointOfContactBody models.CreatePointOfContactBody, fileID uuid.UUID) (*models.PointOfContact, *errors.Error) {
	pointOfContact := models.PointOfContact{
		Name:        pointOfContactBody.Name,
		Email:       pointOfContactBody.Email,
		Position:    pointOfContactBody.Position,
		ClubID:      clubID,
		PhotoFileID: fileID,
	}

	if err := db.Create(&pointOfContact).Error; err != nil {
		return nil, &errors.FailedToCreatePointOfContact
	}
	return &pointOfContact, nil
}

func GetClubPointOfContacts(db *gorm.DB, clubID uuid.UUID) ([]models.PointOfContact, *errors.Error) {
	var pointOfContacts []models.PointOfContact
	var club models.Club

	if err := db.First(&club, clubID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.FailedToGetAllPointOfContact
		} else {
			return nil, &errors.FailedToGetClub
		}
	} else {
		if err = db.Find(&pointOfContacts).Error; err != nil {
			return nil, &errors.FailedToGetAllPointOfContact
		}
		return pointOfContacts, nil
	}
}

// also get the file associated with the point of contact
func GetClubPointOfContact(db *gorm.DB, pocID uuid.UUID, clubID uuid.UUID) (*models.PointOfContact, *errors.Error) {
	var pointOfContact models.PointOfContact
	var club models.Club

	if err := db.First(&club, clubID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.FailedToGetClub
		} else {
			return nil, &errors.FailedToGetPointOfContact
		}
	} else {
		if err = db.Where("id = ?", pocID).First(&pointOfContact).Error; err != nil {
			if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
				return nil, &errors.PointOfContactNotFound
			} else {
				return nil, &errors.FailedToGetPointOfContact
			}
		}
		return &pointOfContact, nil
	}
}

func DeleteClubPointOfContact(db *gorm.DB, pocID uuid.UUID, clubID uuid.UUID) *errors.Error {
	var pointOfContact models.PointOfContact
	var club models.Club

	if err := db.First(&club, clubID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return &errors.FailedToGetClub
		} else {
			return &errors.FailedToDeletePointOfContact
		}
	} else {
		if err = db.Where("id = ?", pocID).First(&pointOfContact).Error; err != nil {
			if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
				return &errors.PointOfContactNotFound
			} else {
				return &errors.FailedToDeletePointOfContact
			}
		}
		if err = db.Delete(&pointOfContact).Error; err != nil {
			return &errors.FailedToDeletePointOfContact
		}
		return nil
	}
}
