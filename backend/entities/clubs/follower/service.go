package follower

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubFollowerServiceInterface interface {
	GetClubFollowers(clubID string, pagination middleware.Pagination) ([]models.User, error)
}

type ClubFollowerService struct {
	types.ServiceParams
}

func NewClubFollowerService(params types.ServiceParams) ClubFollowerServiceInterface {
	return &ClubFollowerService{params}
}

func (cf *ClubFollowerService) GetClubFollowers(clubID string, pagination middleware.Pagination) ([]models.User, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	return GetClubFollowers(cf.DB, *idAsUUID, pagination)
}
