package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type ClubFollowerServiceInterface interface {
	GetClubFollowers(clubID string, limit string, page string) ([]models.User, *errors.Error)
}

type ClubFollowerService struct {
	types.ServiceParams
}

func NewClubFollowerService(params types.ServiceParams) *ClubFollowerService {
	return &ClubFollowerService{params}
}

func (cf *ClubFollowerService) GetClubFollowers(clubID string, limit string, page string) ([]models.User, *errors.Error) {
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

	return transactions.GetClubFollowers(cf.DB, *idAsUUID, *limitAsInt, *pageAsInt)
}
