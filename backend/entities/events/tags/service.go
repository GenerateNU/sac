package tags

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/tags"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type EventTagServiceInterface interface {
	CreateEventTags(id string, clubTagsBody CreateEventTagsRequestBody) ([]models.Tag, error)
	GetEventTags(id string) ([]models.Tag, error)
	DeleteEventTag(id string, tagId string) error
}

type EventTagService struct {
	types.ServiceParams
}

func NewEventTagService(serviceParams types.ServiceParams) EventTagServiceInterface {
	return &EventTagService{serviceParams}
}

func (e *EventTagService) CreateEventTags(id string, eventTagsBody CreateEventTagsRequestBody) ([]models.Tag, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(e.Validate, eventTagsBody); err != nil {
		return nil, err
	}

	tags, err := tags.GetTagsByIDs(e.DB, eventTagsBody.Tags)
	if err != nil {
		return nil, err
	}

	return CreateEventTags(e.DB, *idAsUUID, tags)
}

func (e *EventTagService) GetEventTags(id string) ([]models.Tag, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	return GetEventTags(e.DB, *idAsUUID)
}

func (e *EventTagService) DeleteEventTag(id string, tagId string) error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return err
	}

	tagIdAsUUID, err := utilities.ValidateID(tagId)
	if err != nil {
		return err
	}

	return DeleteEventTag(e.DB, *idAsUUID, *tagIdAsUUID)
}
