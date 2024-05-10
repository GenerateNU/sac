package event

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubEventServiceInterface interface {
	GetClubEvents(clubID string, limit string, page string) ([]models.Event, error)
}

type ClubEventService struct {
	types.ServiceParams
}

func NewClubEventService(params types.ServiceParams) ClubEventServiceInterface {
	return &ClubEventService{params}
}

func (c *ClubEventService) GetClubEvents(clubID string, limit string, page string) ([]models.Event, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, err
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, err
	}

	return GetClubEvents(c.DB, *idAsUUID, *limitAsInt, *pageAsInt)
}
