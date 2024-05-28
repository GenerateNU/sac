package recruitment

import (
	"context"
	"errors"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errs"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type ClubRecruitmentServicer interface {
	CreateClubRecruitment(ctx context.Context, clubID string, body CreateClubRecruitmentRequestBody) (*models.Club, error)
	UpdateClubRecruitment(ctx context.Context, clubID string, body UpdateClubRecruitmentRequestBody) (*models.Club, error)
	DeleteClubRecruitment(ctx context.Context, clubID string) error

	GetClubRecruitmentApplications(ctx context.Context, clubID string, recruitmentID string, pageInfo fiberpaginate.PageInfo) ([]models.Application, error)
	CreateApplicationRequestBody(ctx context.Context, clubID string, recruitmentID string, body CreateApplicationRequestBody) (*models.Application, error)

	GetClubRecruitmentApplication(ctx context.Context, clubID string, recruitmentID string, applicationID string) (*models.Application, error)
	UpdateClubRecruitmentApplication(ctx context.Context, clubID string, recruitmentID string, applicationID string, body UpdateApplicationRequestBody) (*models.Application, error)
	DeleteClubRecruitmentApplication(ctx context.Context, clubID string, recruitmentID string, applicationID string) error
}

type ClubRecruitmentService struct {
	types.ServiceParams
}

func NewClubRecruitmentService(params types.ServiceParams) ClubRecruitmentServicer {
	return &ClubRecruitmentService{params}
}

func (s *ClubRecruitmentService) CreateClubRecruitment(ctx context.Context, clubID string, body CreateClubRecruitmentRequestBody) (*models.Club, error) {
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

	createClubRecruitmentCtx, createClubRecruitmentCancel := context.WithTimeoutCause(ctx, constants.DB_TIMEOUT, errs.ErrDatabaseTimeout)
	defer createClubRecruitmentCancel()

	return CreateClubRecruitment(createClubRecruitmentCtx, s.DB, idAsUUID, recruitment)
}

func (s *ClubRecruitmentService) UpdateClubRecruitment(ctx context.Context, clubID string, body UpdateClubRecruitmentRequestBody) (*models.Club, error) {
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

	updateClubRecruitmentCtx, updateClubRecruitmentCancel := context.WithTimeoutCause(ctx, constants.DB_TIMEOUT, errs.ErrDatabaseTimeout)
	defer updateClubRecruitmentCancel()

	return UpdateClubRecruitment(updateClubRecruitmentCtx, s.DB, idAsUUID, recruitment)
}

func (s *ClubRecruitmentService) DeleteClubRecruitment(ctx context.Context, clubID string) error {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}

	deleteClubRecruitmentCtx, deleteClubRecruitmentCancel := context.WithTimeoutCause(ctx, constants.DB_TIMEOUT, errs.ErrDatabaseTimeout)
	defer deleteClubRecruitmentCancel()

	return DeleteClubRecruitment(deleteClubRecruitmentCtx, s.DB, idAsUUID)
}

func (s *ClubRecruitmentService) GetClubRecruitmentApplications(ctx context.Context, clubID string, recruitmentID string, pageInfo fiberpaginate.PageInfo) ([]models.Application, error) {
	return nil, nil
}

func (s *ClubRecruitmentService) CreateApplicationRequestBody(ctx context.Context, clubID string, recruitmentID string, body CreateApplicationRequestBody) (*models.Application, error) {
	return nil, nil
}

func (s *ClubRecruitmentService) GetClubRecruitmentApplication(ctx context.Context, clubID string, recruitmentID string, applicationID string) (*models.Application, error) {
	return nil, nil
}

func (s *ClubRecruitmentService) UpdateClubRecruitmentApplication(ctx context.Context, clubID string, recruitmentID string, applicationID string, body UpdateApplicationRequestBody) (*models.Application, error) {
	return nil, nil
}

func (s *ClubRecruitmentService) DeleteClubRecruitmentApplication(ctx context.Context, clubID string, recruitmentID string, applicationID string) error {
	return nil
}
