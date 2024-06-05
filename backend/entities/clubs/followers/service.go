package followers

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type ClubFollowerServiceInterface interface {
	GetClubFollowers(clubID string, pageInfo fiberpaginate.PageInfo) ([]models.User, error)
	CreateClubFollower(clubID string, userID string) error
	DeleteClubFollower(clubID string, userID string) error
}

type ClubFollowerService struct {
	types.ServiceParams
}

func NewClubFollowerService(params types.ServiceParams) ClubFollowerServiceInterface {
	return &ClubFollowerService{params}
}

func (cf *ClubFollowerService) GetClubFollowers(clubID string, pageInfo fiberpaginate.PageInfo) ([]models.User, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	return GetClubFollowers(cf.DB, *idAsUUID, pageInfo)
}

func (cf *ClubFollowerService) CreateClubFollower(clubID string, userID string) error {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}

	userIDAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return err
	}

	return CreateClubFollower(cf.DB, *clubIDAsUUID, *userIDAsUUID)
}

func (cf *ClubFollowerService) DeleteClubFollower(clubID string, userID string) error {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}

	userIDAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return err
	}

	return DeleteClubFollower(cf.DB, *clubIDAsUUID, *userIDAsUUID)
}
