package tags

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/tags"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubTagServiceInterface interface {
	CreateClubTags(id string, clubTagsBody models.CreateClubTagsRequestBody) ([]models.Tag, error)
	GetClubTags(id string) ([]models.Tag, error)
	DeleteClubTag(id string, tagId string) error
}

type ClubTagService struct {
	types.ServiceParams
}

func NewClubTagService(serviceParams types.ServiceParams) ClubTagServiceInterface {
	return &ClubTagService{serviceParams}
}

func (c *ClubTagService) CreateClubTags(id string, clubTagsBody models.CreateClubTagsRequestBody) ([]models.Tag, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(c.Validate, clubTagsBody); err != nil {
		return nil, err
	}

	tags, err := tags.GetTagsByIDs(c.DB, clubTagsBody.Tags)
	if err != nil {
		return nil, err
	}

	return CreateClubTags(c.DB, *idAsUUID, tags)
}

func (c *ClubTagService) GetClubTags(id string) ([]models.Tag, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return GetClubTags(c.DB, *idAsUUID)
}

func (c *ClubTagService) DeleteClubTag(id string, tagId string) error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	tagIdAsUUID, err := utilities.ValidateID(tagId)
	if err != nil {
		return err
	}

	return DeleteClubTag(c.DB, *idAsUUID, *tagIdAsUUID)
}
