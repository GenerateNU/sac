package member

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubMemberServiceInterface interface {
	GetClubMembers(clubID string, limit string, page string) ([]models.User, *errors.Error)
}

type ClubMemberService struct {
	types.ServiceParams
}

func NewClubMemberService(params types.ServiceParams) ClubMemberServiceInterface {
	return &ClubMemberService{params}
}

func (cms *ClubMemberService) GetClubMembers(clubID string, limit string, page string) ([]models.User, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	return GetClubMembers(cms.DB, *idAsUUID, *limitAsInt, *pageAsInt)
}
