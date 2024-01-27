package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ClubServiceInterface interface {
	UpsertPointOfContact(clubId string, pointOfContactBody models.CreatePointOfContactBody) (*models.PointOfContact, *errors.Error)
	GetAllPointOfContacts(clubId string) ([]models.PointOfContact, *errors.Error)
	// GetPointOfContact(clubId string, email string) (models.PointOfContact, *errors.Error)
	DeletePointOfContact(email string, clubId string) *errors.Error
}

type ClubService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

// Upsert A Point of Contact
func (u *ClubService) UpsertPointOfContact(clubId string, pointOfContactBody models.CreatePointOfContactBody) (*models.PointOfContact, *errors.Error) {
	if err := u.Validate.Struct(pointOfContactBody); err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.FailedToValidatePointOfContact}
	}
	pointOfContact, err := utilities.MapResponseToModel(pointOfContactBody, &models.PointOfContact{})
	if err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToMapResposeToModel}
	}
	clubIdAsUInt, err := utilities.ValidateID(clubId)
	pointOfContact.ClubID = *clubIdAsUInt
	if err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.FailedToValidateID}
	}
    return transactions.UpsertPointOfContact(u.DB, pointOfContact)
}

// Get All Point of Contact
func (u *ClubService) GetAllPointOfContacts(clubId string) ([]models.PointOfContact, *errors.Error) {
	clubIdAsUint, err := utilities.ValidateID(clubId)
	if err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.FailedToValidateClubId}
	}
	return transactions.GetAllPointOfContacts(u.DB, *clubIdAsUint)
}

// // Get Point of Contact
// func (u *ClubService) GetPointOfContact(clubId string, email string) (*models.PointOfContact, *errors.Error) {
// 	clubIdAsUint, err := utilities.ValidateID(clubId)
// 	if err != nil {
// 		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.FailedToValidateClubId}
// 	}
// 	return transactions.GetPointOfContact(u.DB, email, *clubIdAsUint)
// }


// Delete A Point of Contact with Specific Email
func (u *ClubService) DeletePointOfContact(email string, clubId string) *errors.Error {
	clubIdAsUint, errID := utilities.ValidateID(clubId)
	if errID != nil {
		return &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.FailedToValidateClubId}
	}
	if !utilities.ValidateGenericEmail(email) {
		return &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.FailedToValidateEmail}
	}
	return transactions.DeletePointOfContact(u.DB, email, *clubIdAsUint)
}