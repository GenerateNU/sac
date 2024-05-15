package event

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type ClubEventServiceInterface interface {
	GetClubEvents(clubID string, pageInfo fiberpaginate.PageInfo) ([]models.Event, error)
}

type ClubEventService struct {
	types.ServiceParams
}

func NewClubEventService(params types.ServiceParams) ClubEventServiceInterface {
	return &ClubEventService{params}
}

func (c *ClubEventService) GetClubEvents(clubID string, pageInfo fiberpaginate.PageInfo) ([]models.Event, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	return GetClubEvents(c.DB, *idAsUUID, pageInfo)
}
