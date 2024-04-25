package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ContactServiceInterface interface {
	GetContacts(limit string, page string) ([]models.Contact, *errors.Error)
	GetContact(contactID string) (*models.Contact, *errors.Error)
	DeleteContact(contactID string) *errors.Error
}

type ContactService struct {
	types.ServiceParams
}

func NewContactService(serviceParams types.ServiceParams) ContactServiceInterface {
	return &ContactService{serviceParams}
}

func (c *ContactService) GetContacts(limit string, page string) ([]models.Contact, *errors.Error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	return GetContacts(c.DB, *limitAsInt, *pageAsInt)
}

func (c *ContactService) GetContact(contactID string) (*models.Contact, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(contactID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return GetContact(c.DB, *idAsUUID)
}

func (c *ContactService) DeleteContact(contactID string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(contactID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return DeleteContact(c.DB, *idAsUUID)
}
