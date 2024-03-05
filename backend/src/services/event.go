package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type EventServiceInterface interface {
	GetEvents(limit string, page string) ([]models.Event, *errors.Error)
	GetEvent(eventID string) (*models.Event, *errors.Error)
	GetSeriesByEventID(eventID string) ([]models.Event, *errors.Error)
	GetSeriesByID(seriesID string) ([]models.Event, *errors.Error)
	CreateEvent(eventBodies models.CreateEventRequestBody) ([]models.Event, *errors.Error)
	UpdateEvent(eventID string, eventBody models.UpdateEventRequestBody) (*models.Event, *errors.Error)
	UpdateSeries(eventID string, seriesID string, seriesBody models.UpdateSeriesRequestBody) ([]models.Event, *errors.Error)
	DeleteEvent(eventID string) *errors.Error
	DeleteSeriesByEventID(seriesID string) *errors.Error
	DeleteSeriesByID(seriesID string) *errors.Error
}

type EventService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewEventService(db *gorm.DB, validate *validator.Validate) *EventService {
	return &EventService{DB: db, Validate: validate}
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

	offset := (*pageAsInt - 1) * *limitAsInt

	return transactions.GetEvents(e.DB, *limitAsInt, offset)
}

// TODO: add logic for creating the []event here
// TODO Q: should we always return a slice of events? or should we return a slice of events if it's a series and a single event if it's not?
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
	}

	if !event.IsRecurring {
		event, err := transactions.CreateEvent(e.DB, *event)
		if err != nil {
			return nil, &errors.FailedToCreateEvent
		}
		return []models.Event{*event}, err
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

	return transactions.CreateEventSeries(e.DB, *series)
}

func (e *EventService) GetEvent(id string) (*models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetEvent(e.DB, *idAsUUID)
}

func (e *EventService) GetSeriesByEventID(id string) ([]models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetSeriesByEventID(e.DB, *idAsUUID)
}

func (e *EventService) GetSeriesByID(id string) ([]models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetSeriesByID(e.DB, *idAsUUID)
}

func (e *EventService) UpdateEvent(id string, eventBody models.UpdateEventRequestBody) (*models.Event, *errors.Error) {
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

	event, err := utilities.MapRequestToModel(eventBody, &models.UpdateEventRequestBody{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.UpdateEvent(e.DB, *idAsUUID, *event)
}

func (e *EventService) UpdateSeries(eventID string, seriesID string, seriesBody models.UpdateSeriesRequestBody) ([]models.Event, *errors.Error) {
	eventIDAsUUID, idErr := utilities.ValidateID(eventID)
	if idErr != nil {
		return nil, idErr
	}

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

	return transactions.UpdateSeries(e.DB, *eventIDAsUUID, *seriesIDAsUUID, *series)
}

func (e *EventService) DeleteEvent(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteEvent(e.DB, *idAsUUID)
}

func (e *EventService) DeleteSeriesByEventID(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteSeriesByEventID(e.DB, *idAsUUID)
}

func (e *EventService) DeleteSeriesByID(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteSeriesByID(e.DB, *idAsUUID)
}

// Helper to create other events in a given series using the given firstEvent
func createEventSlice(firstEvent *models.Event, series models.Series) []models.Event {
	eventBodies := []models.Event{*firstEvent}
	months, days := 0, 0

	// currently 0-indexed (separation count 0 means every week/day/month)
	switch series.RecurringType {
	case "daily":
		days = series.SeparationCount + 1
	case "weekly":
		days = 7 * (series.SeparationCount + 1)
	case "monthly":
		months = series.SeparationCount + 1
	}

	for i := 1; i < series.MaxOccurrences; i++ {
		eventToAdd := *firstEvent
		eventToAdd.StartTime = eventToAdd.StartTime.AddDate(0, i*months, i*days)
		eventToAdd.EndTime = eventToAdd.EndTime.AddDate(0, i*months, i*days)
		eventBodies = append(eventBodies, eventToAdd)
	}

	return eventBodies
}
