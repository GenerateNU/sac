package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type PointOfContactServiceInterface interface {
	GetPointOfContacts(limit string, page string) ([]models.PointOfContact, error)
	GetPointOfContact(pocID string) (*models.PointOfContact, error)
}

type PointOfContactService struct {
	types.ServiceParams
}

func NewPointOfContactService(serviceParams types.ServiceParams) PointOfContactServiceInterface {
	return &PointOfContactService{serviceParams}
}

func (poc *PointOfContactService) GetPointOfContacts(limit string, page string) ([]models.PointOfContact, error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, err
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, err
	}

	return GetPointOfContacts(poc.DB, *limitAsInt, *pageAsInt)
}

func (poc *PointOfContactService) GetPointOfContact(pocID string) (*models.PointOfContact, error) {
	idAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, err
	}

	return GetPointOfContact(poc.DB, *idAsUUID)
}
