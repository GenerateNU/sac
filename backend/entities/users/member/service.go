package member

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type UserMemberServiceInterface interface {
	GetMembership(userID string) ([]models.Club, *errors.Error)
}

type UserMemberService struct {
	types.ServiceParams
}

func NewUserMemberService(serviceParams types.ServiceParams) UserMemberServiceInterface {
	return &UserMemberService{serviceParams}
}

func (u *UserMemberService) GetMembership(userID string) ([]models.Club, *errors.Error) {
	userIdAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return nil, err
	}

	return GetClubMembership(u.DB, *userIdAsUUID)
}
