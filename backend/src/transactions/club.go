package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Upsert Point of Contact
func UpsertPointOfContact(db *gorm.DB, pointOfContact *models.PointOfContact) (*models.PointOfContact, *errors.Error) {
	// if err := db.Clause(&pointOfContact).Error; err != nil {
	// 	return nil, &errors.Error{
	// 		StatusCode: fiber.StatusInternalServerError,
	// 		Message: errors.FailedToUpsertPointOfContact}
	// }
	// return pointOfContact, nil
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}, {Name: "club_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "photo", "email", "position"}),
	  }).Create(&pointOfContact).Error
	if 	err != nil { 
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToUpsertPointOfContact}
	}
	return pointOfContact, nil
}

// Get All Point of Contacts
// for users, find all points of contact
func GetAllPointOfContacts(db *gorm.DB, clubId uint) ([]models.PointOfContact, *errors.Error) {
	var pointOfContacts []models.PointOfContact
	var club models.Club

	// handles error when club id is not found
	if err := db.First(&club, clubId).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.ClubNotFound}
		} else {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetClub}
		}
	} else {
		// club id is found, handle error when failed to get point of contact
		if err = db.Find(&pointOfContacts).Error; err != nil {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetAllPointOfContact}
		}
		return pointOfContacts, nil
	}
}

// func GetPointOfContact(db *gorm.DB, email string, clubId uint) (*models.PointOfContact, *errors.Error) {
// 	var club models.Club

// 	if err := db.First(&club, clubId).Error; err != nil {
// 		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.ClubNotFound}
// 		} else {
// 			return nil,  &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetClub}
// 		}
// 	} else {
// 		var pointOfContact *models.PointOfContact
// 		result := db.First(&pointOfContact, email).Error 
// 		if result != nil {
// 			return nil,  &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetAPointOfContact}
// 		}
// 		return pointOfContact, nil
// 	}
// }

// Delete A Point of Contact with specific email
func DeletePointOfContact(db *gorm.DB, email string, clubId uint) *errors.Error {
	var deletedPointOfContact models.PointOfContact
	var club models.Club

	// handles when club id is not found
	if err := db.First(&club, clubId).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.ClubNotFound}
		} else {
			return &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetClub}
		}
	} else {
		// search for point of contact's email to delete
		result := db.Where("email = ?", email).Delete(&deletedPointOfContact)
		if result.RowsAffected == 0 {
			if result.Error == nil {
				return &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.PointOfContactNotFound}
			} else {
				return &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToDeletePointOfContact}
			}
		}
		return nil
	}
}
