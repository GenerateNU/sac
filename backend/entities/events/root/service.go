package root

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type EventServiceInterface interface {
	GetEvents(limit string, page string) ([]models.Event, *errors.Error)
	GetEvent(eventID string) ([]models.Event, *errors.Error)
	GetSeriesByEventID(eventID string) ([]models.Event, *errors.Error)
	GetSeriesByID(seriesID string) ([]models.Event, *errors.Error)
	CreateEvent(eventBodies models.CreateEventRequestBody) ([]models.Event, *errors.Error)
	UpdateEvent(eventID string, eventBody models.UpdateEventRequestBody) ([]models.Event, *errors.Error)
	UpdateSeries(seriesID string, seriesBody models.UpdateSeriesRequestBody) ([]models.Event, *errors.Error)
	UpdateSeriesByEventID(eventID string, seriesBody models.UpdateSeriesRequestBody) ([]models.Event, *errors.Error)
	DeleteEvent(eventID string) *errors.Error
	DeleteSeriesByEventID(seriesID string) *errors.Error
	DeleteSeriesByID(seriesID string) *errors.Error
	GetHostsByEventID(eventID string) ([]models.Club, *errors.Error)
	GetTagsByEventID(eventID string) ([]models.Tag, *errors.Error)
}

type EventService struct {
	types.ServiceParams
}

func NewEventService(serviceParams types.ServiceParams) EventServiceInterface {
	return &EventService{serviceParams}
}

func (e *EventService) GetEvents(limit string, page string) ([]models.Event, *errors.Error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	return GetEvents(e.DB, *limitAsInt, *pageAsInt)
}

// right now we are always returning a slice
func (e *EventService) CreateEvent(eventBody models.CreateEventRequestBody) ([]models.Event, *errors.Error) {
	if err := e.Validate.Struct(eventBody); err != nil {
		return nil, &errors.FailedToValidateEvent
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
		event, err := CreateEvent(e.DB, *event)
		if err != nil {
			return nil, &errors.FailedToCreateEvent
		}
		return event, err
	}

	if err := e.Validate.Struct(eventBody.Series); err != nil {
		return nil, &errors.FailedToValidateEventSeries
	}

	series, err := utilities.MapRequestToModel(eventBody.Series, &models.Series{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	// Create other events in series and update field in series (for join table)
	events := createEventSlice(event, *series)
	series.Events = events

	return CreateEventSeries(e.DB, *series)
}

func (e *EventService) GetEvent(eventID string) ([]models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return GetEvent(e.DB, *idAsUUID)
}

func (e *EventService) GetSeriesByEventID(eventID string) ([]models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return GetSeriesByEventID(e.DB, *idAsUUID)
}

func (e *EventService) GetSeriesByID(seriesID string) ([]models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(seriesID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return GetSeriesByID(e.DB, *idAsUUID)
}

func (e *EventService) UpdateEvent(id string, eventBody models.UpdateEventRequestBody) ([]models.Event, *errors.Error) {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	if utilities.AtLeastOne(eventBody, models.UpdateEventRequestBody{}) {
		return nil, &errors.FailedToValidateTag
	}

	if err := e.Validate.Struct(eventBody); err != nil {
		return nil, &errors.FailedToValidateEvent
	}

	updatedEvent := mapToEvent(eventBody)

	return UpdateEvent(e.DB, *idAsUUID, *updatedEvent)
}

func (e *EventService) UpdateSeries(seriesID string, seriesBody models.UpdateSeriesRequestBody) ([]models.Event, *errors.Error) {
	seriesIDAsUUID, idErr := utilities.ValidateID(seriesID)
	if idErr != nil {
		return nil, idErr
	}

	if err := e.Validate.Struct(seriesBody); err != nil {
		return nil, &errors.FailedToValidateEventSeries
	}

	series, err := utilities.MapRequestToModel(seriesBody, &models.Series{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	updatedEventBody := mapToEvent(seriesBody.EventDetails)

	newEvents := createEventSlice(updatedEventBody, *series)
	series.Events = newEvents

	return UpdateSeries(e.DB, *seriesIDAsUUID, *series)
}

func (e *EventService) UpdateSeriesByEventID(eventID string, seriesBody models.UpdateSeriesRequestBody) ([]models.Event, *errors.Error) {
	eventIDAsUUID, idErr := utilities.ValidateID(eventID)
	if idErr != nil {
		return nil, idErr
	}

	if err := e.Validate.Struct(seriesBody); err != nil {
		return nil, &errors.FailedToValidateEventSeries
	}

	series, err := utilities.MapRequestToModel(seriesBody, &models.Series{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	updatedEventBody := mapToEvent(seriesBody.EventDetails)

	newEvents := createEventSlice(updatedEventBody, *series)
	series.Events = newEvents

	return UpdateSeriesByEventID(e.DB, *eventIDAsUUID, *series)
}

func (e *EventService) DeleteEvent(eventID string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return DeleteEvent(e.DB, *idAsUUID)
}

func (e *EventService) DeleteSeriesByEventID(eventID string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return DeleteSeriesByEventID(e.DB, *idAsUUID)
}

func (e *EventService) DeleteSeriesByID(seriesID string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(seriesID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return DeleteSeriesByID(e.DB, *idAsUUID)
}

func (e *EventService) GetHostsByEventID(eventID string) ([]models.Club, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return GetHostsByEventID(e.DB, *idAsUUID)
}

func (e *EventService) GetTagsByEventID(eventID string) ([]models.Tag, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, &errors.FailedToValidateID
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
