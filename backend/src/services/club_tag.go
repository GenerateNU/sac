package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type ClubTagServiceInterface interface {
	CreateClubTags(id string, clubTagsBody models.CreateClubTagsRequestBody) ([]models.Tag, *errors.Error)
	GetClubTags(id string) ([]models.Tag, *errors.Error)
	DeleteClubTag(id string, tagId string) *errors.Error
}

type ClubTagService struct {
	types.ServiceParams
}

func NewClubTagService(params types.ServiceParams) ClubTagServiceInterface {
	return &ClubTagService{params}
}

func (c *ClubTagService) CreateClubTags(id string, clubTagsBody models.CreateClubTagsRequestBody) ([]models.Tag, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if err := c.Validate.Struct(clubTagsBody); err != nil {
		return nil, &errors.FailedToValidateClubTags
	}

	tags, err := transactions.GetTagsByIDs(c.DB, clubTagsBody.Tags)
	if err != nil {
		return nil, err
	}

	return transactions.CreateClubTags(c.DB, *idAsUUID, tags)
}

func (c *ClubTagService) GetClubTags(id string) ([]models.Tag, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetClubTags(c.DB, *idAsUUID)
}

func (c *ClubTagService) DeleteClubTag(id string, tagId string) *errors.Error {
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
