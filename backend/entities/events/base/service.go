package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/events"
	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type EventServiceInterface interface {
	// getters
	GetEvents(pageInfo fiberpaginate.PageInfo, start string, end string) ([]models.Event, error)
	GetEvent(eventID string) (*models.Event, error)
	// event cud
	CreateEvent(body events.CreateEventRequestBody) (*models.Event, error)
	UpdateEvent(eventID string, eventBody events.UpdateEventRequestBody) ([]models.Event, error)
	DeleteEvent(eventID string) error
}

type EventService struct {
	types.ServiceParams
}

func NewEventService(serviceParams types.ServiceParams) EventServiceInterface {
	return &EventService{serviceParams}
}

func (e *EventService) GetEvents(pageInfo fiberpaginate.PageInfo, start string, end string) ([]models.Event, error) {
	if start == "" || end == "" {
		return events.GetEvents(e.DB, pageInfo)
	}

	startTime, err := utilities.ParseTime(start, utilities.YYYY_dash_MM_dash_DD)
	if err != nil {
		return nil, utilities.BadRequest(err)
	}

	endTime, err := utilities.ParseTime(end, utilities.YYYY_dash_MM_dash_DD)
	if err != nil {
		return nil, utilities.BadRequest(err)
	}

	if startTime.After(*endTime) {
		return nil, utilities.BadRequest(errors.New("start time must be before end time"))
	}

	return GetEventsByTime(e.DB, pageInfo, *startTime, *endTime)
}

func (e *EventService) GetEvent(eventID string) (*models.Event, error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, err
	}

	return events.GetEvent(e.DB, *idAsUUID)
}

func (e *EventService) CreateEvent(body events.CreateEventRequestBody) (*models.Event, error) {
	if err := utilities.Validate(e.Validate, body); err != nil {
		return nil, err
	}

	return CreateEvent(e.DB, *body.Into())
}

func (e *EventService) UpdateEvent(eventID string, eventBody events.UpdateEventRequestBody) ([]models.Event, error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(e.Validate, eventBody); err != nil {
		return nil, err
	}

	event, err := utilities.MapJsonTags(eventBody, &models.Event{})
	if err != nil {
		return nil, err
	}

	return UpdateEvent(e.DB, *idAsUUID, *event, *eventBody.UpdateAllInSeries)
}

func (e *EventService) DeleteEvent(eventID string) error {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return err
	}

	return DeleteEvent(e.DB, *idAsUUID)
}
