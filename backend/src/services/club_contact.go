package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ClubContactServiceInterface interface {
	GetClubContacts(clubID string) ([]models.Contact, *errors.Error)
	PutClubContact(clubID string, contactBody models.PutContactRequestBody) (*models.Contact, *errors.Error)
}

type ClubContactService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewClubContactService(db *gorm.DB, validate *validator.Validate) *ClubContactService {
	return &ClubContactService{DB: db, Validate: validate}
}

func (c *ClubContactService) GetClubContacts(clubID string) ([]models.Contact, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetClubContacts(c.DB, *idAsUUID)
}

func (c *ClubContactService) PutClubContact(clubID string, contactBody models.PutContactRequestBody) (*models.Contact, *errors.Error) {
	idAsUUID, idErr := utilities.ValidateID(clubID)
	if idErr != nil {
		return nil, idErr
	}

	if err := c.Validate.Struct(contactBody); err != nil {
		return nil, &errors.FailedToValidateContact
	}

	contact, err := utilities.MapRequestToModel(contactBody, &models.Contact{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	contact.ClubID = *idAsUUID

	return transactions.PutClubContact(c.DB, *contact)
}
