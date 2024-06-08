package followers

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type UserFollowerServiceInterface interface {
	GetFollowing(userID string) ([]models.Club, error)
}

type UserFollowerService struct {
	types.ServiceParams
}

func NewUserFollowerService(serviceParams types.ServiceParams) UserFollowerServiceInterface {
	return &UserFollowerService{serviceParams}
}

func (u *UserFollowerService) GetFollowing(userID string) ([]models.Club, error) {
	userIDAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return nil, err
	}

	return GetClubFollowing(u.DB, *userIDAsUUID)
}
