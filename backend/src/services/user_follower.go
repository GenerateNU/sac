package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type UserFollowerServiceInterface interface {
	CreateFollowing(userId string, clubId string) *errors.Error
	DeleteFollowing(userId string, clubId string) *errors.Error
	GetFollowing(userId string) ([]models.Club, *errors.Error)
}

type UserFollowerService struct {
	types.ServiceParams
}

func NewUserFollowerService(params types.ServiceParams) UserFollowerServiceInterface {
	return &UserFollowerService{params}
}

func (u *UserFollowerService) CreateFollowing(userId string, clubId string) *errors.Error {
	userIdAsUUID, err := utilities.ValidateID(userId)
	if err != nil {
		return err
	}
	clubIdAsUUID, err := utilities.ValidateID(clubId)
	if err != nil {
		return err
	}
	return transactions.CreateFollowing(u.DB, *userIdAsUUID, *clubIdAsUUID)
}

func (u *UserFollowerService) DeleteFollowing(userId string, clubId string) *errors.Error {
	userIdAsUUID, err := utilities.ValidateID(userId)
	if err != nil {
		return err
	}
	clubIdAsUUID, err := utilities.ValidateID(clubId)
	if err != nil {
		return err
	}
	return transactions.DeleteFollowing(u.DB, *userIdAsUUID, *clubIdAsUUID)
}

func (u *UserFollowerService) GetFollowing(userId string) ([]models.Club, *errors.Error) {
	userIdAsUUID, err := utilities.ValidateID(userId)
	if err != nil {
		return nil, err
	}

	return transactions.GetClubFollowing(u.DB, *userIdAsUUID)
}
