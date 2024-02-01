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
	GetClubs(limit string, page string) ([]models.Club, *errors.Error)
	GetClub(id string) (*models.Club, *errors.Error)
	CreateClub(clubBody models.CreateClubRequestBody) (*models.Club, *errors.Error)
	UpdateClub(id string, clubBody models.UpdateClubRequestBody) (*models.Club, *errors.Error)
	DeleteClub(id string) *errors.Error
	CreateClubTags(id string, clubTagsBody models.CreateClubTagsRequestBody) ([]models.Tag, *errors.Error)
	GetClubTags(id string) ([]models.Tag, *errors.Error)
	DeleteClubTag(id string, tagId string) *errors.Error
}

type ClubService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (c *ClubService) GetClubs(limit string, page string) ([]models.Club, *errors.Error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)

	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)

	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	offset := (*pageAsInt - 1) * *limitAsInt

	return transactions.GetClubs(c.DB, *limitAsInt, offset)
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

func (c *ClubService) CreateClubTags(id string, clubTagsBody models.CreateClubTagsRequestBody) ([]models.Tag, *errors.Error) {
	// Validate the id:
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if err := c.Validate.Struct(clubTagsBody); err != nil {
		return nil, &errors.FailedToValidateClubTags
	}

	// Retrieve a list of valid tags from the ids:
	tags, err := transactions.GetTagsByIDs(c.DB, clubTagsBody.Tags)

	if err != nil {
		return nil, err
	}

	// Update the club to reflect the new tags:
	return transactions.CreateClubTags(c.DB, *idAsUUID, tags)
}

func (c *ClubService) GetClubTags(id string) ([]models.Tag, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetClubTags(c.DB, *idAsUUID)
}

func (c *ClubService) DeleteClubTag(id string, tagId string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	tagIdAsUUID, err := utilities.ValidateID(tagId)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteClubTag(c.DB, *idAsUUID, *tagIdAsUUID)
}