package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubPointOfContacts(db *gorm.DB, clubID uuid.UUID) ([]models.PointOfContact, *errors.Error) {
	var pointOfContacts []models.PointOfContact

	result := db.Preload("PhotoFile").Where("club_id = ?", clubID).Find(&pointOfContacts)
	if result.Error != nil {
		return nil, &errors.FailedToGetClubPointOfContacts
	}

	return pointOfContacts, nil
}

func GetClubPointOfContact(db *gorm.DB, clubID uuid.UUID, pocID uuid.UUID) (*models.PointOfContact, *errors.Error) {
	var pointOfContact models.PointOfContact

	if err := db.Preload("PhotoFile").First(&pointOfContact, "id = ? AND club_id = ?", pocID, clubID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.PointOfContactNotFound
		} else {
			return nil, &errors.FailedToGetClubPointOfContact
		}
	}

	return &pointOfContact, nil
}

func GetClubPointOfContactByClubIDAndEmail(db *gorm.DB, clubID uuid.UUID, email string) (*models.PointOfContact, *errors.Error) {
	var pointOfContact models.PointOfContact

	if err := db.First(&pointOfContact, "email = ? AND club_id = ?", email, clubID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.PointOfContactNotFound
		} else {
			return nil, &errors.FailedToGetClubPointOfContact
		}
	}

	return &pointOfContact, nil
}

func CreateClubPointOfContact(db *gorm.DB, clubID uuid.UUID, pointOfContactBody models.CreatePointOfContactBody) (*models.PointOfContact, *errors.Error) {
	pointOfContact := models.PointOfContact{
		Name:     pointOfContactBody.Name,
		Email:    pointOfContactBody.Email,
		Position: pointOfContactBody.Position,
		ClubID:   clubID,
	}

	if err := db.Create(&pointOfContact).Error; err != nil {
		return nil, &errors.FailedToCreatePointOfContact
	}
	return &pointOfContact, nil
}

func UpdateClubPointOfContact(db *gorm.DB, clubID uuid.UUID, pocID uuid.UUID, pointOfContactBody models.UpdatePointOfContactBody) (*models.PointOfContact, *errors.Error) {
	var pointOfContact models.PointOfContact

	if err := db.Preload("PhotoFile").First(&pointOfContact, "id = ? AND club_id = ?", pocID, clubID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.PointOfContactNotFound
		} else {
			return nil, &errors.FailedToGetClubPointOfContact
		}
	}

	if err := db.Model(&pointOfContact).Updates(models.PointOfContact{
		Name:     pointOfContactBody.Name,
		Email:    pointOfContactBody.Email,
		Position: pointOfContactBody.Position,
	}).Error; err != nil {
		return nil, &errors.FailedToUpdateClubPointOfContact
	}

	return &pointOfContact, nil
}

func DeleteClubPointOfContact(db *gorm.DB, clubID uuid.UUID, pocID uuid.UUID) *errors.Error {
	if result := db.Delete(&models.PointOfContact{}, "id = ? AND club_id = ?", pocID, clubID); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.PointOfContactNotFound
		} else {
			return &errors.FailedToDeleteClubPointOfContact
		}
	}
	return nil
}
