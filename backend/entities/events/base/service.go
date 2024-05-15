package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type EventServiceInterface interface {
	GetEvents(pageInfo fiberpaginate.PageInfo) ([]models.Event, error)
	GetEvent(eventID string) ([]models.Event, error)
	GetSeriesByEventID(eventID string) ([]models.Event, error)
	GetSeriesByID(seriesID string) ([]models.Event, error)
	CreateEvent(eventBodies models.CreateEventRequestBody) ([]models.Event, error)
	UpdateEvent(eventID string, eventBody models.UpdateEventRequestBody) ([]models.Event, error)
	UpdateSeries(seriesID string, seriesBody models.UpdateSeriesRequestBody) ([]models.Event, error)
	UpdateSeriesByEventID(eventID string, seriesBody models.UpdateSeriesRequestBody) ([]models.Event, error)
	DeleteEvent(eventID string) error
	DeleteSeriesByEventID(seriesID string) error
	DeleteSeriesByID(seriesID string) error
	GetHostsByEventID(eventID string) ([]models.Club, error)
	GetTagsByEventID(eventID string) ([]models.Tag, error)
}

type EventService struct {
	types.ServiceParams
}

func NewEventService(serviceParams types.ServiceParams) EventServiceInterface {
	return &EventService{serviceParams}
}

func (e *EventService) GetEvents(pageInfo fiberpaginate.PageInfo) ([]models.Event, error) {
	return GetEvents(e.DB, pageInfo)
}

// right now we are always returning a slice
func (e *EventService) CreateEvent(eventBody models.CreateEventRequestBody) ([]models.Event, error) {
	if err := utilities.Validate(e.Validate, eventBody); err != nil {
		return nil, err
	}

	event := &models.Event{
		Name:        eventBody.Name,
		Preview:     eventBody.Preview,
		Content:     eventBody.Content,
		StartTime:   eventBody.StartTime,
		EndTime:     eventBody.EndTime,
		Location:    eventBody.Location,
		EventType:   eventBody.EventType,
		IsRecurring: *eventBody.IsRecurring,

		Host:         eventBody.Host,
		Clubs:        eventBody.Clubs,
		Tag:          eventBody.Tag,
		Notification: eventBody.Notification,
	}

	if !event.IsRecurring {
		return CreateEvent(e.DB, *event)
	}

	if err := utilities.Validate(e.Validate, eventBody.Series); err != nil {
		return nil, err
	}

	series, err := utilities.MapRequestToModel(eventBody.Series, &models.Series{})
	if err != nil {
		return nil, err
	}

	// Create other events in series and update field in series (for join table)
	events := createEventSlice(event, *series)
	series.Events = events

	return CreateEventSeries(e.DB, *series)
}

func (e *EventService) GetEvent(eventID string) ([]models.Event, error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, err
	}

	return GetEvent(e.DB, *idAsUUID)
}

func (e *EventService) GetSeriesByEventID(eventID string) ([]models.Event, error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, err
	}

	return GetSeriesByEventID(e.DB, *idAsUUID)
}

func (e *EventService) GetSeriesByID(seriesID string) ([]models.Event, error) {
	idAsUUID, err := utilities.ValidateID(seriesID)
	if err != nil {
		return nil, err
	}

	return GetSeriesByID(e.DB, *idAsUUID)
}

func (e *EventService) UpdateEvent(id string, eventBody models.UpdateEventRequestBody) ([]models.Event, error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, err
	}

	if utilities.AtLeastOne(eventBody, models.UpdateEventRequestBody{}) {
		return nil, err
	}

	if err := utilities.Validate(e.Validate, eventBody); err != nil {
		return nil, err
	}

	updatedEvent := mapToEvent(eventBody)

	return UpdateEvent(e.DB, *idAsUUID, *updatedEvent)
}

func (e *EventService) UpdateSeries(seriesID string, seriesBody models.UpdateSeriesRequestBody) ([]models.Event, error) {
	seriesIDAsUUID, err := utilities.ValidateID(seriesID)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(e.Validate, seriesBody); err != nil {
		return nil, err
	}

	series, err := utilities.MapRequestToModel(seriesBody, &models.Series{})
	if err != nil {
		return nil, err
	}

	updatedEventBody := mapToEvent(seriesBody.EventDetails)

	newEvents := createEventSlice(updatedEventBody, *series)
	series.Events = newEvents

	return UpdateSeries(e.DB, *seriesIDAsUUID, *series)
}

func (e *EventService) UpdateSeriesByEventID(eventID string, seriesBody models.UpdateSeriesRequestBody) ([]models.Event, error) {
	eventIDAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(e.Validate, seriesBody); err != nil {
		return nil, err
	}

	series, err := utilities.MapRequestToModel(seriesBody, &models.Series{})
	if err != nil {
		return nil, err
	}

	updatedEventBody := mapToEvent(seriesBody.EventDetails)

	newEvents := createEventSlice(updatedEventBody, *series)
	series.Events = newEvents

	return UpdateSeriesByEventID(e.DB, *eventIDAsUUID, *series)
}

func (e *EventService) DeleteEvent(eventID string) error {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return err
	}

	return DeleteEvent(e.DB, *idAsUUID)
}

func (e *EventService) DeleteSeriesByEventID(eventID string) error {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return err
	}

	return DeleteSeriesByEventID(e.DB, *idAsUUID)
}

func (e *EventService) DeleteSeriesByID(seriesID string) error {
	idAsUUID, err := utilities.ValidateID(seriesID)
	if err != nil {
		return err
	}

	return DeleteSeriesByID(e.DB, *idAsUUID)
}

func (e *EventService) GetHostsByEventID(eventID string) ([]models.Club, error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, err
	}

	return GetHostsByEventID(e.DB, *idAsUUID)
}

func (e *EventService) GetTagsByEventID(eventID string) ([]models.Tag, error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, err
	}

	return GetTagsByEventID(e.DB, *idAsUUID)
}

func createEventSlice(firstEvent *models.Event, series models.Series) []models.Event {
	eventBodies := []models.Event{*firstEvent}
	months, days := 0, 0

	switch series.RecurringType {
	case "daily":
		days = 1
	case "weekly":
		days = 7
	case "monthly":
		months = 1
	}

	for i := 1; i < series.MaxOccurrences; i++ {
		eventToAdd := *firstEvent
		eventToAdd.StartTime = eventToAdd.StartTime.AddDate(0, i*months, i*days)
		eventToAdd.EndTime = eventToAdd.EndTime.AddDate(0, i*months, i*days)
		eventBodies = append(eventBodies, eventToAdd)
	}

	return eventBodies
}

func mapToEvent(eventBody models.UpdateEventRequestBody) *models.Event {
	return &models.Event{
		Name:      eventBody.Name,
		Preview:   eventBody.Preview,
		Content:   eventBody.Content,
		StartTime: eventBody.StartTime,
		EndTime:   eventBody.EndTime,
		Location:  eventBody.Location,
		EventType: eventBody.EventType,

		Host:         eventBody.Host,
		RSVP:         eventBody.RSVP,
		Waitlist:     eventBody.Waitlist,
		Clubs:        eventBody.Clubs,
		Tag:          eventBody.Tag,
		Notification: eventBody.Notification,
	}
}
