package base

import (
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errs"
	"github.com/GenerateNU/sac/backend/search"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"
)

type ClubServiceInterface interface {
	GetClubs(pageInfo fiberpaginate.PageInfo) ([]models.Club, error)
	GetClub(id string) (*models.Club, error)
	CreateClub(userID uuid.UUID, clubBody CreateClubRequestBody) (*models.Club, error)
	UpdateClub(id string, clubBody UpdateClubRequestBody) (*models.Club, error)
	DeleteClub(id string) error
}

type ClubService struct {
	types.ServiceParams
}

func NewClubService(serviceParams types.ServiceParams) ClubServiceInterface {
	return &ClubService{serviceParams}
}

func (c *ClubService) GetClubs(pagination fiberpaginate.PageInfo) ([]models.Club, error) {
	return GetClubs(c.DB, pagination)
}

func (c *ClubService) CreateClub(userID uuid.UUID, clubBody CreateClubRequestBody) (*models.Club, error) {
	if err := utilities.Validate(c.Validate, clubBody); err != nil {
		return nil, err
	}

	club, err := utilities.MapJsonTags(clubBody, &models.Club{})
	if err != nil {
		return nil, err
	}

	result, err := CreateClub(c.DB, userID, *club)
	if err != nil {
		return nil, err
	}

	go search.UpsertClubEmbedding(c.DB, c.Search, result)
	return result, err
}

func (c *ClubService) GetClub(id string) (*models.Club, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return clubs.GetClub(c.DB, *idAsUUID)
}

func (c *ClubService) UpdateClub(id string, clubBody UpdateClubRequestBody) (*models.Club, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if utilities.AtLeastOne(clubBody, UpdateClubRequestBody{}) {
		return nil, errs.ErrAtLeastOne
	}

	if err := utilities.Validate(c.Validate, clubBody); err != nil {
		return nil, err
	}

	club, err := utilities.MapJsonTags(clubBody, &models.Club{})
	if err != nil {
		return nil, err
	}

	return UpdateClub(c.DB, *idAsUUID, *club)
}

func (c *ClubService) DeleteClub(id string) error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return DeleteClub(c.DB, *idAsUUID)
}
