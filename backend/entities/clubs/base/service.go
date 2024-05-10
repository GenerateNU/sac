package base

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubServiceInterface interface {
	GetClubs(queryParams *models.ClubQueryParams) ([]models.Club, error)
	GetClub(id string) (*models.Club, error)
	CreateClub(clubBody models.CreateClubRequestBody) (*models.Club, error)
	UpdateClub(id string, clubBody models.UpdateClubRequestBody) (*models.Club, error)
	DeleteClub(id string) error
}

type ClubService struct {
	types.ServiceParams
}

func NewClubService(serviceParams types.ServiceParams) ClubServiceInterface {
	return &ClubService{serviceParams}
}

func (c *ClubService) GetClubs(queryParams *models.ClubQueryParams) ([]models.Club, error) {
	if queryParams.Limit < 0 {
		return nil, fmt.Errorf("limit must be greater than 0")
	}

	if queryParams.Page < 0 {
		return nil, fmt.Errorf("page must be greater than 0")
	}

	return GetClubs(c.DB, c.Integrations.Search, queryParams)
}

func (c *ClubService) CreateClub(clubBody models.CreateClubRequestBody) (*models.Club, error) {
	if err := utilities.Validate(c.Validate, clubBody); err != nil {
		return nil, err
	}

	club, err := utilities.MapRequestToModel(clubBody, &models.Club{})
	if err != nil {
		return nil, err
	}

	return CreateClub(c.DB, c.Integrations.Search, clubBody.UserID, *club)
}

func (c *ClubService) GetClub(id string) (*models.Club, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return clubs.GetClub(c.DB, *idAsUUID)
}

func (c *ClubService) UpdateClub(id string, clubBody models.UpdateClubRequestBody) (*models.Club, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if utilities.AtLeastOne(clubBody, models.UpdateClubRequestBody{}) {
		return nil, fmt.Errorf("at least one field must be present")
	}

	if err := utilities.Validate(c.Validate, clubBody); err != nil {
		return nil, err
	}

	club, err := utilities.MapRequestToModel(clubBody, &models.Club{})
	if err != nil {
		return nil, err
	}

	return UpdateClub(c.DB, c.Integrations.Search, *idAsUUID, *club)
}

func (c *ClubService) DeleteClub(id string) error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	return DeleteClub(c.DB, c.Integrations.Search, *idAsUUID)
}
