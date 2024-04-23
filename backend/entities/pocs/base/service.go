package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
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

	return GetPointOfContacts(poc.DB, *limitAsInt, *pageAsInt)
}

func (poc *PointOfContactService) GetPointOfContact(pocID string) (*models.PointOfContact, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return GetPointOfContact(poc.DB, *idAsUUID)
}
