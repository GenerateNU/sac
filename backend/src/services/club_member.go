package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type ClubMemberServiceInterface interface {
	GetClubMembers(clubID string, limit string, page string) ([]models.User, *errors.Error)
}

type ClubMemberService struct {
	types.ServiceParams
}

func NewClubMemberService(params types.ServiceParams) *ClubMemberService {
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

	return transactions.GetClubMembers(cms.DB, *idAsUUID, *limitAsInt, *pageAsInt)
}
