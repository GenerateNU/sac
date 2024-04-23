package member

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
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

	return CreateMember(u.DB, *userIdAsUUID, *clubIdAsUUID)
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
	return DeleteMember(u.DB, *userIdAsUUID, *clubIdAsUUID)
}

func (u *UserMemberService) GetMembership(userID string) ([]models.Club, *errors.Error) {
	userIdAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return nil, err
	}

	return GetClubMembership(u.DB, *userIdAsUUID)
}
