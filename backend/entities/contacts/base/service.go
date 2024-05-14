package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ContactServiceInterface interface {
	GetContacts(pagination middleware.Pagination) ([]models.Contact, error)
	GetContact(contactID string) (*models.Contact, error)
	DeleteContact(contactID string) error
}

type ContactService struct {
	types.ServiceParams
}

func NewContactService(serviceParams types.ServiceParams) ContactServiceInterface {
	return &ContactService{serviceParams}
}

func (c *ContactService) GetContacts(pagination middleware.Pagination) ([]models.Contact, error) {
	return GetContacts(c.DB, pagination)
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
