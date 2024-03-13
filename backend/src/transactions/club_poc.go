package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// // Upsert Point of Contact
// func UpsertPointOfContact(db *gorm.DB, pointOfContact *models.PointOfContact) (*models.PointOfContact, *errors.Error) {
// 	pocExist, errPOCExist := GetPointOfContact(db, pointOfContact.ID, pointOfContact.ClubID)
// 	if errPOCExist != nil {
// 		db.Model(&pointOfContact).Where("id = ?", pocExist.ID).Association("File").Replace(pocExist, pointOfContact)
// 	} else {
// 		db.Model(&pointOfContact).Association("File").Append(pointOfContact)
// 	}
// 	// err := db.Clauses(clause.OnConflict{
// 	// 	Columns:   []clause.Column{{Name: "email"}, {Name: "club_id"}},
// 	// 	DoUpdates: clause.AssignmentColumns([]string{"name", "email", "position"}),
// 	// }).Create(&pointOfContact).Error
// 	// if err != nil {
// 	// 	return nil, &errors.FailedToUpsertPointOfContact
// 	// }
// 	return pointOfContact, nil
// }

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

func GetClubPointOfContact(db *gorm.DB, pocID uuid.UUID, clubID uuid.UUID) (*models.PointOfContact, *errors.Error) {
	var club models.Club
	if err := db.First(&club, clubID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.ClubNotFound
		} else {
			return nil, &errors.FailedToGetClub
		}
	} else {
		var pointOfContact models.PointOfContact
		if err := db.First(&pointOfContact, pocID).Error; err != nil {
			if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
				return nil, &errors.PointOfContactNotFound
			} else {
				return nil, &errors.FailedToGetAPointOfContact
			}
		}
		return &pointOfContact, nil
	}
}

func DeleteClubPointOfContact(db *gorm.DB, pocID uuid.UUID, clubID uuid.UUID) *errors.Error {
	var deletedPointOfContact models.PointOfContact
	var club models.Club

	if err := db.First(&club, clubID).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return &errors.FailedToDeletePointOfContact
		} else {
			return &errors.FailedToGetClub
		}
	} else {
		result := db.Where("id = ?", pocID).Delete(&deletedPointOfContact)
		if result.RowsAffected == 0 {
			if result.Error == nil {
				return &errors.PointOfContactNotFound
			} else {
				return &errors.FailedToDeletePointOfContact
			}
		}
		return nil
	}
}
