package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type PointOfContactServiceInterface interface {
	GetPointOfContacts(limit string, page string) ([]models.PointOfContact, *errors.Error)
	GetPointOfContact(pocID string) (*models.PointOfContact, *errors.Error)
}

type PointOfContactService struct {
	types.ServiceParams
}

func NewPointOfContactService(serviceParams types.ServiceParams) PointOfContactServiceInterface {
	return &PointOfContactService{serviceParams}
}

func (poc *PointOfContactService) GetPointOfContacts(limit string, page string) ([]models.PointOfContact, *errors.Error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	return transactions.GetPointOfContacts(poc.DB, *limitAsInt, *pageAsInt)
}

func (poc *PointOfContactService) GetPointOfContact(pocID string) (*models.PointOfContact, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetPointOfContact(poc.DB, *idAsUUID)
}
