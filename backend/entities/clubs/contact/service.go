package contact

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubContactServiceInterface interface {
	GetClubContacts(clubID string) ([]models.Contact, *errors.Error)
	PutClubContact(clubID string, contactBody models.PutContactRequestBody) (*models.Contact, *errors.Error)
}

type ClubContactService struct {
	types.ServiceParams
}

func NewClubContactService(params types.ServiceParams) ClubContactServiceInterface {
	return &ClubContactService{params}
}

func (c *ClubContactService) GetClubContacts(clubID string) ([]models.Contact, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return GetClubContacts(c.DB, *idAsUUID)
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

	return PutClubContact(c.DB, *contact)
}
