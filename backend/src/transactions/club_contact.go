package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PutClubContact(db *gorm.DB, contact models.Contact) (*models.Contact, *errors.Error) {
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "club_id"}, {Name: "type"}},
		DoUpdates: clause.AssignmentColumns([]string{"content"}),
	}).Create(&contact).Error
	if err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) || stdliberrors.Is(err, gorm.ErrForeignKeyViolated) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToPutContact
		}
	}
	return &contact, nil
}

func GetClubContacts(db *gorm.DB, clubID uuid.UUID) ([]models.Contact, *errors.Error) {
	var club models.Club
	if err := db.Preload("Contact").First(&club, clubID).Error; err != nil {
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
