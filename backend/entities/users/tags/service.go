package tags

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/tags"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type UserTagServiceInterface interface {
	GetUserTags(id string) ([]models.Tag, error)
	CreateUserTags(id string, tagIDs CreateUserTagsBody) ([]models.Tag, error)
	DeleteUserTag(id string, tagID string) error
}

type UserTagService struct {
	types.ServiceParams
}

func NewUserTagService(serviceParams types.ServiceParams) UserTagServiceInterface {
	return &UserTagService{serviceParams}
}

func (u *UserTagService) GetUserTags(id string) ([]models.Tag, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return GetUserTags(u.DB, *idAsUUID)
}

func (u *UserTagService) CreateUserTags(id string, tagIDs CreateUserTagsBody) ([]models.Tag, error) {
	// Validate the id:
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(u.Validate, tagIDs); err != nil {
		return nil, err
	}

	// Retrieve a list of valid tags from the ids:
	tags, err := tags.GetTagsByIDs(u.DB, tagIDs.Tags)
	if err != nil {
		return nil, err
	}

	// Update the user to reflect the new tags:
	return CreateUserTags(u.DB, *idAsUUID, tags)
}

func (u *UserTagService) DeleteUserTag(id string, tagID string) error {
	// Validate the userID:
	userIDAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	tagIDAsUUID, err := utilities.ValidateID(tagID)
	if err != nil {
		return err
	}

	return DeleteUserTag(u.DB, *userIDAsUUID, *tagIDAsUUID)
}
