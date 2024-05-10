package follower

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type UserFollowerServiceInterface interface {
	CreateFollowing(userId string, clubId string) error
	DeleteFollowing(userId string, clubId string) error
	GetFollowing(userId string) ([]models.Club, error)
}

type UserFollowerService struct {
	types.ServiceParams
}

func NewUserFollowerService(serviceParams types.ServiceParams) UserFollowerServiceInterface {
	return &UserFollowerService{serviceParams}
}

func (u *UserFollowerService) CreateFollowing(userId string, clubId string) error {
	userIdAsUUID, err := utilities.ValidateID(userId)
	if err != nil {
		return err
	}
	clubIdAsUUID, err := utilities.ValidateID(clubId)
	if err != nil {
		return err
	}
	return CreateFollowing(u.DB, *userIdAsUUID, *clubIdAsUUID)
}

func (u *UserFollowerService) DeleteFollowing(userId string, clubId string) error {
	userIdAsUUID, err := utilities.ValidateID(userId)
	if err != nil {
		return err
	}
	clubIdAsUUID, err := utilities.ValidateID(clubId)
	if err != nil {
		return err
	}
	return DeleteFollowing(u.DB, *userIdAsUUID, *clubIdAsUUID)
}

func (u *UserFollowerService) GetFollowing(userId string) ([]models.Club, error) {
	userIdAsUUID, err := utilities.ValidateID(userId)
	if err != nil {
		return nil, err
	}

	return GetClubFollowing(u.DB, *userIdAsUUID)
}
