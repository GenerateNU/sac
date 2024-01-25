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
	CreateOrUpdatePointOfContact(pointOfContactBody models.PointOfContact) (*models.PointOfContact, *errors.Error)
	GetPointOfContact(email string, club_id string) (*models.PointOfContact, *errors.Error)
	DeletePointOfContact(email string, club_id string) *errors.Error
}

type ClubService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

// Point of Contact
// Create or Update Point of Contact 
func (u *ClubService) CreateOrUpdatePointOfContact(pointOfContact models.PointOfContact) (*models.PointOfContact, *errors.Error) {
	if err := u.Validate.Struct(pointOfContact); err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.FailedToValidatePointOfContact}
	}
	return transactions.CreateorUpdatePointOfContact(u.DB, &pointOfContact)
}

// Get Point of Contact 
func (u *ClubService) GetPointOfContact(email string, club_id string) (*models.PointOfContact, *errors.Error) {
	club_idAsInt, err := utilities.ValidateID(club_id)
	if err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.FailedToValidatePointOfContact}
	}
	return transactions.GetPointOfContact(u.DB, email, *club_idAsInt)
}

// Delete Point of Contact 
func (u *ClubService) DeletePointOfContact(email string, club_id string) *errors.Error {
	club_idAsInt, err := utilities.ValidateID(club_id)
	if err != nil {
		return &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.FailedToValidatePointOfContact}
	}
	return transactions.DeletePointOfContact(u.DB, email, *club_idAsInt)
}