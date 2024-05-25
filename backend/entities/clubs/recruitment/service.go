package recruitment

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubRecruitmentServicer interface {
	CreateClubRecruitment(clubID string, body CreateClubRecruitmentRequestBody) (*models.Club, error)
	UpdateClubRecruitment(clubID string, body UpdateClubRecruitmentRequestBody) (*models.Club, error)
	DeleteClubRecruitment(clubID string) error
}

type ClubRecruitmentService struct {
	types.ServiceParams
}

func NewClubRecruitmentService(params types.ServiceParams) ClubRecruitmentServicer {
	return &ClubRecruitmentService{params}
}

func (s *ClubRecruitmentService) CreateClubRecruitment(clubID string, body CreateClubRecruitmentRequestBody) (*models.Club, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(s.Validate, body); err != nil {
		return nil, err
	}

	recruitment, err := utilities.MapJsonTags(body, &models.Recruitment{})
	if err != nil {
		return nil, err
	}

	return CreateClubRecruitment(s.DB, idAsUUID, recruitment)
}

func (s *ClubRecruitmentService) UpdateClubRecruitment(clubID string, body UpdateClubRecruitmentRequestBody) (*models.Club, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	if utilities.AtLeastOne(body, UpdateClubRecruitmentRequestBody{}) {
		return nil, errors.New("at least one field must be present")
	}

	if err := utilities.Validate(s.Validate, body); err != nil {
		return nil, err
	}

	recruitment, err := utilities.MapJsonTags(body, &models.Recruitment{})
	if err != nil {
		return nil, err
	}

	return UpdateClubRecruitment(s.DB, idAsUUID, recruitment)
}

func (s *ClubRecruitmentService) DeleteClubRecruitment(clubID string) error {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}

	return DeleteClubRecruitment(s.DB, idAsUUID)
}
