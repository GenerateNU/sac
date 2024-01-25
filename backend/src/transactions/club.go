package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// func UpdatePointOfContact(db *gorm.DB, id uint, email string, pointOfContact *models.PointOfContact) (*models.PointOfContact, *errors.Error) {
// 	var existingPointOfContact models.PointOfContact

// 	err := db.First(&existingPointOfContact, id).Error
// 	if err != nil {
// 		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetPointOfContact}
// 	}
// 	if err := db.Model(&existingPointOfContact).Updates(&existingPointOfContact).Error; err != nil {
// 		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToUpdatePointOfContact}
// 	}
// 	return &existingPointOfContact, nil
// }

// Create or Update Point of Contact
func CreateorUpdatePointOfContact(db *gorm.DB, pointOfContact *models.PointOfContact) (*models.PointOfContact, *errors.Error) {
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}, {Name: "club_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "email", "photo", "position"}),
	}).Save(&pointOfContact).Error; err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToUpsertPointOfContact}
	}
	return pointOfContact, nil
}

// Get Point of Contact
func GetPointOfContact(db *gorm.DB, email string, id uint) (*models.PointOfContact, *errors.Error) {
	var pointOfContact models.PointOfContact

	if err := db.First(&pointOfContact, email, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.PointOfContactNotFound}
		} else {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetPointOfContact}
		}
	}
	return &pointOfContact, nil
}

// Delete Point of Contact
func DeletePointOfContact(db *gorm.DB, email string, id uint) *errors.Error {
	var deletedPointOfContact models.PointOfContact

	result := db.Where("email = ? AND id = ?", email, id).Delete(&deletedPointOfContact)
	if result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.PointOfContactNotFound}
		} else {
			return &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToDeletePointOfContact}
		}
	}
	return nil
}
