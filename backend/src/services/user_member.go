package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type UserMemberServiceInterface interface {
	CreateMembership(userID string, clubID string) *errors.Error
	DeleteMembership(userID string, clubID string) *errors.Error
	GetMembership(userID string) ([]models.Club, *errors.Error)
}

type UserMemberService struct {
	types.ServiceParams
}

func NewUserMemberService(params types.ServiceParams) UserMemberServiceInterface {
	return &UserMemberService{params}
}

func (u *UserMemberService) CreateMembership(userID string, clubID string) *errors.Error {
	userIdAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return err
	}
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}

	return transactions.CreateMember(u.DB, *userIdAsUUID, *clubIdAsUUID)
}

func (u *UserMemberService) DeleteMembership(userID string, clubID string) *errors.Error {
	userIdAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return err
	}

	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}
	return transactions.DeleteMember(u.DB, *userIdAsUUID, *clubIdAsUUID)
}

func (u *UserMemberService) GetMembership(userID string) ([]models.Club, *errors.Error) {
	userIdAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return nil, err
	}

	return transactions.GetClubMembership(u.DB, *userIdAsUUID)
}
