package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type LeaderServiceInterface interface {
	GetLeaders(pageInfo fiberpaginate.PageInfo) ([]models.Leader, error)
	GetLeader(leaderID string) (*models.Leader, error)
}

type LeaderService struct {
	types.ServiceParams
}

func NewLeaderService(serviceParams types.ServiceParams) LeaderServiceInterface {
	return &LeaderService{serviceParams}
}

func (ls *LeaderService) GetLeaders(pageInfo fiberpaginate.PageInfo) ([]models.Leader, error) {
	return GetLeadership(ls.DB, pageInfo)
}

func (ls *LeaderService) GetLeader(leaderID string) (*models.Leader, error) {
	idAsUUID, err := utilities.ValidateID(leaderID)
	if err != nil {
		return nil, err
	}

	return GetLeader(ls.DB, *idAsUUID)
}
