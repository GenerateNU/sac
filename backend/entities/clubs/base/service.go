package base

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type ClubServiceInterface interface {
	GetClubs(pageInfo fiberpaginate.PageInfo) ([]models.Club, error)
	GetClub(id string) (*models.Club, error)
	CreateClub(clubBody CreateClubRequestBody) (*models.Club, error)
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

func (c *ClubService) CreateClub(clubBody CreateClubRequestBody) (*models.Club, error) {
	if err := utilities.Validate(c.Validate, clubBody); err != nil {
		return nil, err
	}

	club, err := utilities.MapJsonTags(clubBody, &models.Club{})
	if err != nil {
		return nil, err
	}

	return CreateClub(c.DB, clubBody.UserID, *club)
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
		return nil, fmt.Errorf("at least one field must be present")
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
