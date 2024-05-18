package contacts

import (
	"errors"
	"fmt"

	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PutClubContact(db *gorm.DB, contact models.Contact) (*models.Contact, error) {
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "club_id"}, {Name: "type"}},
		DoUpdates: clause.AssignmentColumns([]string{"content"}),
	}).Create(&contact).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrForeignKeyViolated) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}
	return &contact, nil
}

func GetClubContacts(db *gorm.DB, clubID uuid.UUID) ([]models.Contact, error) {
	var club models.Club

	if err := db.Preload("Contact").First(&club, clubID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	if club.Contact == nil {
		return nil, fmt.Errorf("club with ID %s has no contacts", clubID)
	}

	return club.Contact, nil
}
