package base

import (
	"github.com/GenerateNU/sac/backend/entities/events"
	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type EventServiceInterface interface {
	// getters
	GetEvents(limit string, page string) ([]models.Event, error)
	GetEvent(eventID string) (*models.Event, error)
	// event cud
	CreateEvent(body models.CreateEventRequestBody) (*models.Event, error)
	UpdateEvent(eventID string, eventBody models.UpdateEventRequestBody) (*models.Event, error)
	DeleteEvent(eventID string) error
}

type EventService struct {
	types.ServiceParams
}

func NewEventService(serviceParams types.ServiceParams) EventServiceInterface {
	return &EventService{serviceParams}
}

func (e *EventService) GetEvents(limit string, page string) ([]models.Event, error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, err
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, err
	}

	return GetEvents(e.DB, *limitAsInt, *pageAsInt)
}

func (e *EventService) GetEvent(eventID string) (*models.Event, error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, err
	}

	return events.GetEvent(e.DB, *idAsUUID)
}

func (e *EventService) CreateEvent(body models.CreateEventRequestBody) (*models.Event, error) {
	if err := utilities.Validate(e.Validate, body); err != nil {
		return nil, err
	}

	event, err := utilities.MapJsonTags(body, &models.Event{})
	if err != nil {
		return nil, err
	}

	return CreateEvent(e.DB, *event)
}

func (e *EventService) UpdateEvent(eventID string, eventBody models.UpdateEventRequestBody) (*models.Event, error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, err
	}

	if err := e.Validate.Struct(eventBody); err != nil {
		return nil, err
	}

	event, err := utilities.MapJsonTags(eventBody, &models.Event{})
	if err != nil {
		return nil, err
	}

	return UpdateEvent(e.DB, *idAsUUID, *event)
}

func (e *EventService) DeleteEvent(eventID string) error {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return err
	}

	return DeleteEvent(e.DB, *idAsUUID)
}
