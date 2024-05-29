package recruitment

import (
	"context"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errs"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type ClubRecruitmentServicer interface {
	CreateClubRecruitment(ctx context.Context, clubID string, body CreateClubRecruitmentRequestBody) (*models.Recruitment, error)
	GetClubRecruitment(ctx context.Context, clubID string) (*models.Recruitment, error)
	UpdateClubRecruitment(ctx context.Context, clubID string, body UpdateClubRecruitmentRequestBody) (*models.Recruitment, error)
	DeleteClubRecruitment(ctx context.Context, clubID string) error

	CreateApplicationRequestBody(ctx context.Context, clubID string, recruitmentID string, body CreateApplicationRequestBody) (*models.Application, error)
	GetClubRecruitmentApplications(ctx context.Context, clubID string, recruitmentID string, pageInfo fiberpaginate.PageInfo) ([]models.Application, error)

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

func (s *ClubRecruitmentService) CreateClubRecruitment(ctx context.Context, clubID string, body CreateClubRecruitmentRequestBody) (*models.Recruitment, error) {
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

func (s *ClubRecruitmentService) GetClubRecruitment(ctx context.Context, clubID string) (*models.Recruitment, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	getClubRecruitmentCtx, getClubRecruitmentCancel := context.WithTimeoutCause(ctx, constants.DB_TIMEOUT, errs.ErrDatabaseTimeout)
	defer getClubRecruitmentCancel()

	return GetClubRecruitment(getClubRecruitmentCtx, s.DB, idAsUUID)
}

func (s *ClubRecruitmentService) UpdateClubRecruitment(ctx context.Context, clubID string, body UpdateClubRecruitmentRequestBody) (*models.Recruitment, error) {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	if utilities.AtLeastOne(body, UpdateClubRecruitmentRequestBody{}) {
		return nil, errs.ErrAtLeastOne
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

	return UpdateClubRecruitment(updateClubRecruitmentCtx, s.DB, clubIDAsUUID, recruitment)
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

func (s *ClubRecruitmentService) CreateApplicationRequestBody(ctx context.Context, clubID string, recruitmentID string, body CreateApplicationRequestBody) (*models.Application, error) {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	recruitmentIDAsUUID, err := utilities.ValidateID(recruitmentID)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(s.Validate, body); err != nil {
		return nil, err
	}

	application, err := utilities.MapJsonTags(body, &models.Application{})
	if err != nil {
		return nil, err
	}

	createApplicationCtx, createApplicationCancel := context.WithTimeoutCause(ctx, constants.DB_TIMEOUT, errs.ErrDatabaseTimeout)
	defer createApplicationCancel()

	return CreateApplication(createApplicationCtx, s.DB, clubIDAsUUID, recruitmentIDAsUUID, application)
}

func (s *ClubRecruitmentService) GetClubRecruitmentApplications(ctx context.Context, clubID string, recruitmentID string, pageInfo fiberpaginate.PageInfo) ([]models.Application, error) {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	recruitmentIDAsUUID, err := utilities.ValidateID(recruitmentID)
	if err != nil {
		return nil, err
	}

	getClubRecruitmentApplicationsCtx, getClubRecruitmentApplicationsCancel := context.WithTimeoutCause(ctx, constants.DB_TIMEOUT, errs.ErrDatabaseTimeout)
	defer getClubRecruitmentApplicationsCancel()

	return GetClubRecruitmentApplications(getClubRecruitmentApplicationsCtx, s.DB, clubIDAsUUID, recruitmentIDAsUUID, pageInfo)
}

func (s *ClubRecruitmentService) GetClubRecruitmentApplication(ctx context.Context, clubID string, recruitmentID string, applicationID string) (*models.Application, error) {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	recruitmentIDAsUUID, err := utilities.ValidateID(recruitmentID)
	if err != nil {
		return nil, err
	}

	applicationIDAsUUID, err := utilities.ValidateID(applicationID)
	if err != nil {
		return nil, err
	}

	getClubRecruitmentApplicationCtx, getClubRecruitmentApplicationCancel := context.WithTimeoutCause(ctx, constants.DB_TIMEOUT, errs.ErrDatabaseTimeout)
	defer getClubRecruitmentApplicationCancel()

	return GetClubRecruitmentApplication(getClubRecruitmentApplicationCtx, s.DB, clubIDAsUUID, recruitmentIDAsUUID, applicationIDAsUUID)
}

func (s *ClubRecruitmentService) UpdateClubRecruitmentApplication(ctx context.Context, clubID string, recruitmentID string, applicationID string, body UpdateApplicationRequestBody) (*models.Application, error) {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	recruitmentIDAsUUID, err := utilities.ValidateID(recruitmentID)
	if err != nil {
		return nil, err
	}

	applicationIDAsUUID, err := utilities.ValidateID(applicationID)
	if err != nil {
		return nil, err
	}

	if utilities.AtLeastOne(body, UpdateApplicationRequestBody{}) {
		return nil, errs.ErrAtLeastOne
	}

	if err := utilities.Validate(s.Validate, body); err != nil {
		return nil, err
	}

	application, err := utilities.MapJsonTags(body, &models.Application{})
	if err != nil {
		return nil, err
	}

	updateClubRecruitmentApplicationCtx, updateClubRecruitmentApplicationCancel := context.WithTimeoutCause(ctx, constants.DB_TIMEOUT, errs.ErrDatabaseTimeout)
	defer updateClubRecruitmentApplicationCancel()

	return UpdateClubRecruitmentApplication(updateClubRecruitmentApplicationCtx, s.DB, clubIDAsUUID, recruitmentIDAsUUID, applicationIDAsUUID, application)
}

func (s *ClubRecruitmentService) DeleteClubRecruitmentApplication(ctx context.Context, clubID string, recruitmentID string, applicationID string) error {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}

	recruitmentIDAsUUID, err := utilities.ValidateID(recruitmentID)
	if err != nil {
		return err
	}

	applicationIDAsUUID, err := utilities.ValidateID(applicationID)
	if err != nil {
		return err
	}

	deleteClubRecruitmentApplicationCtx, deleteClubRecruitmentApplicationCancel := context.WithTimeoutCause(ctx, constants.DB_TIMEOUT, errs.ErrDatabaseTimeout)
	defer deleteClubRecruitmentApplicationCancel()

	return DeleteClubRecruitmentApplication(deleteClubRecruitmentApplicationCtx, s.DB, clubIDAsUUID, recruitmentIDAsUUID, applicationIDAsUUID)
}
