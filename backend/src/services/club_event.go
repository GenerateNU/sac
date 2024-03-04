package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type ClubEventServiceInterface interface {
	GetClubEvents(clubID string, limit string, page string) ([]models.Event, *errors.Error)
}

type ClubEventService struct {
	types.ServiceParams
}

func NewClubEventService(params types.ServiceParams) *ClubEventService {
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

	offset := (*pageAsInt - 1) * *limitAsInt

	return transactions.GetClubEvents(c.DB, *idAsUUID, *limitAsInt, offset)
}
