package event

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubEventServiceInterface interface {
	GetClubEvents(clubID string, pagination middleware.Pagination) ([]models.Event, error)
}

type ClubEventService struct {
	types.ServiceParams
}

func NewClubEventService(params types.ServiceParams) ClubEventServiceInterface {
	return &ClubEventService{params}
}

func (c *ClubEventService) GetClubEvents(clubID string, pagination middleware.Pagination) ([]models.Event, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	return GetClubEvents(c.DB, *idAsUUID, pagination)
}
