package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ClubServiceInterface interface {
	GetClubs(queryParams *models.ClubQueryParams) ([]models.Club, *errors.Error)
	GetClub(id string) (*models.Club, *errors.Error)
	CreateClub(clubBody models.CreateClubRequestBody) (*models.Club, *errors.Error)
	UpdateClub(id string, clubBody models.UpdateClubRequestBody) (*models.Club, *errors.Error)
	DeleteClub(id string) *errors.Error
}

type ClubService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewClubService(db *gorm.DB, validate *validator.Validate) *ClubService {
	return &ClubService{DB: db, Validate: validate}
}

func (c *ClubService) GetClubs(queryParams *models.ClubQueryParams) ([]models.Club, *errors.Error) {
	if queryParams.Limit < 0 {
		return nil, &errors.FailedToValidateLimit
	}

	if queryParams.Page < 0 {
		return nil, &errors.FailedToValidatePage
	}

	return transactions.GetClubs(c.DB, queryParams)
}

func (c *ClubService) CreateClub(clubBody models.CreateClubRequestBody) (*models.Club, *errors.Error) {
	if err := c.Validate.Struct(clubBody); err != nil {
		return nil, &errors.FailedToValidateClub
	}

	club, err := utilities.MapRequestToModel(clubBody, &models.Club{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.CreateClub(c.DB, clubBody.UserID, *club)
}

func (c *ClubService) GetClub(id string) (*models.Club, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetClub(c.DB, *idAsUUID)
}

func (c *ClubService) UpdateClub(id string, clubBody models.UpdateClubRequestBody) (*models.Club, *errors.Error) {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	if utilities.AtLeastOne(clubBody, models.UpdateClubRequestBody{}) {
		return nil, &errors.FailedToValidateClub
	}

	if err := c.Validate.Struct(clubBody); err != nil {
		return nil, &errors.FailedToValidateClub
	}

	club, err := utilities.MapRequestToModel(clubBody, &models.Club{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.UpdateClub(c.DB, *idAsUUID, *club)
}

func (c *ClubService) DeleteClub(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteClub(c.DB, *idAsUUID)
}
