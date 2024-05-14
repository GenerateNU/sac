package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type PointOfContactServiceInterface interface {
	GetPointOfContacts(pagination middleware.Pagination) ([]models.PointOfContact, error)
	GetPointOfContact(pocID string) (*models.PointOfContact, error)
}

type PointOfContactService struct {
	types.ServiceParams
}

func NewPointOfContactService(serviceParams types.ServiceParams) PointOfContactServiceInterface {
	return &PointOfContactService{serviceParams}
}

func (poc *PointOfContactService) GetPointOfContacts(pagination middleware.Pagination) ([]models.PointOfContact, error) {
	return GetPointOfContacts(poc.DB, pagination)
}

func (poc *PointOfContactService) GetPointOfContact(pocID string) (*models.PointOfContact, error) {
	idAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, err
	}

	return GetPointOfContact(poc.DB, *idAsUUID)
}
