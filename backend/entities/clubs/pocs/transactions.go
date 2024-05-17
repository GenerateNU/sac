package pocs

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetClubPointOfContacts(db *gorm.DB, clubID uuid.UUID) ([]models.PointOfContact, error) {
	var pointOfContacts []models.PointOfContact

	result := db.Preload("PhotoFile").Where("club_id = ?", clubID).Find(&pointOfContacts)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, result.Error
	}

	return pointOfContacts, nil
}

func GetClubPointOfContact(db *gorm.DB, clubID uuid.UUID, pocID uuid.UUID) (*models.PointOfContact, error) {
	var pointOfContact models.PointOfContact

	if err := db.Preload("PhotoFile").First(&pointOfContact, "id = ? AND club_id = ?", pocID, clubID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &pointOfContact, nil
}

func GetClubPointOfContactByClubIDAndEmail(db *gorm.DB, clubID uuid.UUID, email string) (*models.PointOfContact, error) {
	var pointOfContact models.PointOfContact

	if err := db.First(&pointOfContact, "email = ? AND club_id = ?", email, clubID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &pointOfContact, nil
}

func CreateClubPointOfContact(db *gorm.DB, clubID uuid.UUID, pointOfContactBody models.CreatePointOfContactBody) (*models.PointOfContact, error) {
	pointOfContact := models.PointOfContact{
		Name:     pointOfContactBody.Name,
		Email:    pointOfContactBody.Email,
		Position: pointOfContactBody.Position,
		ClubID:   clubID,
	}

	if err := db.Create(&pointOfContact).Error; err != nil {
		return nil, err
	}
	return &pointOfContact, nil
}

func UpdateClubPointOfContact(db *gorm.DB, clubID uuid.UUID, pocID uuid.UUID, pointOfContactBody models.UpdatePointOfContactBody) (*models.PointOfContact, error) {
	pointOfContact, err := GetClubPointOfContact(db, clubID, pocID)
	if err != nil {
		return nil, err
	}
	if err := db.Model(&pointOfContact).Clauses(clause.Returning{}).Updates(models.PointOfContact{
		Name:     pointOfContactBody.Name,
		Email:    pointOfContactBody.Email,
		Position: pointOfContactBody.Position,
	}).Error; err != nil {
		return nil, err
	}

	return pointOfContact, nil
}

func DeleteClubPointOfContact(db *gorm.DB, clubID uuid.UUID, pocID uuid.UUID) error {
	if result := db.Delete(&models.PointOfContact{}, "id = ? AND club_id = ?", pocID, clubID); result.RowsAffected == 0 {
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}
	return nil
}
