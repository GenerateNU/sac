package follower

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubFollowerServiceInterface interface {
	GetClubFollowers(clubID string, limit string, page string) ([]models.User, error)
}

type ClubFollowerService struct {
	types.ServiceParams
}

func NewClubFollowerService(params types.ServiceParams) ClubFollowerServiceInterface {
	return &ClubFollowerService{params}
}

func (cf *ClubFollowerService) GetClubFollowers(clubID string, limit string, page string) ([]models.User, error) {
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

	return GetClubFollowers(cf.DB, *idAsUUID, *limitAsInt, *pageAsInt)
}
