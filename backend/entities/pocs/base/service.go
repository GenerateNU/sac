package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type PointOfContactServiceInterface interface {
	GetPointOfContacts(pageInfo fiberpaginate.PageInfo) ([]models.PointOfContact, error)
	GetPointOfContact(pocID string) (*models.PointOfContact, error)
}

type PointOfContactService struct {
	types.ServiceParams
}

func NewPointOfContactService(serviceParams types.ServiceParams) PointOfContactServiceInterface {
	return &PointOfContactService{serviceParams}
}

func (poc *PointOfContactService) GetPointOfContacts(pageInfo fiberpaginate.PageInfo) ([]models.PointOfContact, error) {
	return GetPointOfContacts(poc.DB, pageInfo)
}

func (poc *PointOfContactService) GetPointOfContact(pocID string) (*models.PointOfContact, error) {
	idAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, err
	}

	return GetPointOfContact(poc.DB, *idAsUUID)
}
