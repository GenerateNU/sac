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
	GetClubEvents(id string) ([]models.Event, *errors.Error)
	GetEvent(id string) (*models.Event, *errors.Error)
	GetSeriesByEventId(id string) ([]models.Event, *errors.Error)
	GetSeriesById(id string) ([]models.Event, *errors.Error)
	CreateEvent(eventBodies models.CreateEventRequestBody) ([]models.Event, *errors.Error)
	UpdateEvent(id string, eventBody models.UpdateEventRequestBody) (*models.Event, *errors.Error)
	DeleteEvent(id string) *errors.Error
	DeleteSeriesByEventId(id string) *errors.Error
	DeleteSeriesById(id string) *errors.Error
}

type EventService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewEventService(db *gorm.DB, validate *validator.Validate) *EventService {
	return &EventService{DB: db, Validate: validate}
}

func (c *EventService) GetEvents(limit string, page string) ([]models.Event, *errors.Error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	offset := (*pageAsInt - 1) * *limitAsInt

	return transactions.GetEvents(c.DB, *limitAsInt, offset)
}

func (c *EventService) GetClubEvents(id string) ([]models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetClubEvents(c.DB, *idAsUUID)
}

// TODO: add logic for creating the []event here
// TODO Q: should we always return a slice of events? or should we return a slice of events if it's a series and a single event if it's not?
func (c *EventService) CreateEvent(eventBody models.CreateEventRequestBody) ([]models.Event, *errors.Error) {
	if err := c.Validate.Struct(eventBody); err != nil {
		return nil, &errors.FailedToValidateEvent
	}

	// map requestToModels only works well with maps
	// event, err := utilities.MapRequestToModel(eventBody, &models.Event{})

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
		event, err := transactions.CreateEvent(c.DB, *event)
		if err != nil {
			return nil, &errors.FailedToCreateEvent
		}
		return []models.Event{*event}, err
	}

	if err := c.Validate.Struct(eventBody.Series); err != nil {
		return nil, &errors.FailedToValidateEventSeries
	}

	series, err := utilities.MapRequestToModel(eventBody.Series, &models.Series{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	// Create other events in series and update field in series (for join table)
	events := CreateEventSlice(event, *series)
	series.Events = events

	return transactions.CreateEventSeries(c.DB, *series)
}

func (c *EventService) GetEvent(id string) (*models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetEvent(c.DB, *idAsUUID)
}

func (c *EventService) GetSeriesByEventId(id string) ([]models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetSeriesByEventId(c.DB, *idAsUUID)
}

func (c *EventService) GetSeriesById(id string) ([]models.Event, *errors.Error){
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetSeriesById(c.DB, *idAsUUID)
}

func (c *EventService) UpdateEvent(id string, eventBody models.UpdateEventRequestBody) (*models.Event, *errors.Error) {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	if err := c.Validate.Struct(eventBody); err != nil {
		return nil, &errors.FailedToValidateEvent
	}

	event, err := utilities.MapRequestToModel(eventBody, &models.UpdateEventRequestBody{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.UpdateEvent(c.DB, *idAsUUID, *event)
}

func (c *EventService) DeleteEvent(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteEvent(c.DB, *idAsUUID)
}

func (c *EventService) DeleteSeriesByEventId(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteSeriesByEventId(c.DB, *idAsUUID)
}


func (c* EventService) DeleteSeriesById(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteSeriesById(c.DB, *idAsUUID)
}

// Helper to create other events in a given series using the given firstEvent
func CreateEventSlice(firstEvent *models.Event, series models.Series) []models.Event {
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
