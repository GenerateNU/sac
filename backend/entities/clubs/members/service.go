package members

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type ClubMemberServiceInterface interface {
	GetClubMembers(clubID string, pageInfo fiberpaginate.PageInfo) ([]models.User, error)
	CreateClubMember(clubID string, userID string) error
	DeleteClubMember(clubID string, userID string) error
}

type ClubMemberService struct {
	types.ServiceParams
}

func NewClubMemberService(params types.ServiceParams) ClubMemberServiceInterface {
	return &ClubMemberService{params}
}

func (cms *ClubMemberService) GetClubMembers(clubID string, pageInfo fiberpaginate.PageInfo) ([]models.User, error) {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	return GetClubMembers(cms.DB, *clubIDAsUUID, pageInfo)
}

func (cms *ClubMemberService) CreateClubMember(clubID string, userID string) error {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}

	userIDAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return err
	}

	return CreateClubMember(cms.DB, *clubIDAsUUID, *userIDAsUUID)
}

func (cms *ClubMemberService) DeleteClubMember(clubID string, userID string) error {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}

	userIDAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return err
	}

	return DeleteClubMember(cms.DB, *clubIDAsUUID, *userIDAsUUID)
}
