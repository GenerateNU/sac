package tag

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/tags"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubTagServiceInterface interface {
	CreateClubTags(id string, clubTagsBody models.CreateClubTagsRequestBody) ([]models.Tag, *errors.Error)
	GetClubTags(id string) ([]models.Tag, *errors.Error)
	DeleteClubTag(id string, tagId string) *errors.Error
}

type ClubTagService struct {
	types.ServiceParams
}

func NewClubTagService(serviceParams types.ServiceParams) ClubTagServiceInterface {
	return &ClubTagService{serviceParams}
}

func (c *ClubTagService) CreateClubTags(id string, clubTagsBody models.CreateClubTagsRequestBody) ([]models.Tag, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if err := c.Validate.Struct(clubTagsBody); err != nil {
		return nil, &errors.FailedToValidateClubTags
	}

	tags, err := tags.GetTagsByIDs(c.DB, clubTagsBody.Tags)
	if err != nil {
		return nil, err
	}

	return CreateClubTags(c.DB, *idAsUUID, tags)
}

func (c *ClubTagService) GetClubTags(id string) ([]models.Tag, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return GetClubTags(c.DB, *idAsUUID)
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

	return DeleteClubTag(c.DB, *idAsUUID, *tagIdAsUUID)
}
