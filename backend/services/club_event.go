package services

import (
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/models"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubEventServiceInterface interface {
	GetClubEvents(clubID string, limit string, page string) ([]models.Event, *errors.Error)
}

type ClubEventService struct {
	types.ServiceParams
}

func NewClubEventService(params types.ServiceParams) ClubEventServiceInterface {
	return &ClubEventService{params}
}

func (c *ClubEventService) GetClubEvents(clubID string, limit string, page string) ([]models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	return transactions.GetClubEvents(c.DB, *idAsUUID, *limitAsInt, *pageAsInt)
}
