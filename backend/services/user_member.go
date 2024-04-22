package services

import (
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/models"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type UserMemberServiceInterface interface {
	CreateMembership(userID string, clubID string) *errors.Error
	DeleteMembership(userID string, clubID string) *errors.Error
	GetMembership(userID string) ([]models.Club, *errors.Error)
}

type UserMemberService struct {
	types.ServiceParams
}

func NewUserMemberService(serviceParams types.ServiceParams) UserMemberServiceInterface {
	return &UserMemberService{serviceParams}
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
