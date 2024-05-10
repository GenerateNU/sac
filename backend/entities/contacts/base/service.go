package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ContactServiceInterface interface {
	GetContacts(limit string, page string) ([]models.Contact, error)
	GetContact(contactID string) (*models.Contact, error)
	DeleteContact(contactID string) error
}

type ContactService struct {
	types.ServiceParams
}

func NewContactService(serviceParams types.ServiceParams) ContactServiceInterface {
	return &ContactService{serviceParams}
}

func (c *ContactService) GetContacts(limit string, page string) ([]models.Contact, error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, err
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, err
	}

	return GetContacts(c.DB, *limitAsInt, *pageAsInt)
}

func (c *ContactService) GetContact(contactID string) (*models.Contact, error) {
	idAsUUID, err := utilities.ValidateID(contactID)
	if err != nil {
		return nil, err
	}

	return GetContact(c.DB, *idAsUUID)
}

func (c *ContactService) DeleteContact(contactID string) error {
	idAsUUID, err := utilities.ValidateID(contactID)
	if err != nil {
		return err
	}

	return DeleteContact(c.DB, *idAsUUID)
}
