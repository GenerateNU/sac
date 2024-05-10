package contact

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubContactServiceInterface interface {
	GetClubContacts(clubID string) ([]models.Contact, error)
	PutClubContact(clubID string, contactBody models.PutContactRequestBody) (*models.Contact, error)
}

type ClubContactService struct {
	types.ServiceParams
}

func NewClubContactService(params types.ServiceParams) ClubContactServiceInterface {
	return &ClubContactService{params}
}

func (c *ClubContactService) GetClubContacts(clubID string) ([]models.Contact, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	return GetClubContacts(c.DB, *idAsUUID)
}

func (c *ClubContactService) PutClubContact(clubID string, contactBody models.PutContactRequestBody) (*models.Contact, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(c.Validate, contactBody); err != nil {
		return nil, err
	}

	contact, err := utilities.MapRequestToModel(contactBody, &models.Contact{})
	if err != nil {
		return nil, err
	}

	contact.ClubID = *idAsUUID

	return PutClubContact(c.DB, *contact)
}
