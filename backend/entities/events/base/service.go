package base

import (
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/protos"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type EventServiceInterface interface {
	// getters
	GetEvents(limit string, page string) ([]models.Event, *errors.Error)
	GetEvent(eventID string) (models.Event, *errors.Error)
	// event cud
	CreateEvent(body models.CreateEventRequestBody) ([]models.Event, *errors.Error)
	UpdateEvent(eventID string, eventBody models.UpdateEventRequestBody) ([]models.Event, *errors.Error)
	DeleteEvent(eventID string) *errors.Error

	// bridgerton getters (do these belong here?)
	// TODO: move to events/tags, like clubs
	GetTags(eventID string) ([]models.Tag, *errors.Error)
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

func (e *EventService) GetEvent(eventID string) (models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return GetEvent(e.DB, *idAsUUID)
}

func (e *EventService) CreateEvent(body models.CreateEventRequestBody) ([]models.Event, *errors.Error) {
	if err := e.Validate.Struct(body); err != nil {
		return nil, &errors.FailedToValidateEvent
	}

	event, err := utilities.MapRequestToModel(body, &models.Event{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	if body.IsSeries {
		handleEventSeries(event, body.Recurrence, body.Termination)
	}

}

func handleEventSeries(event *models.Event, binaryRecurrence protos.BinaryRecurrence, termination time.Time) ([]models.Event, *errors.Error) {
	var recurrence protos.Recurrence
	if unmarshallErr := recurrence.UnmarshallBinary(binaryRecurrence); unmarshallErr != nil {
		return nil, &errors.FailedToUnmarshallRecurrence
	}

	if termination.Before(event.Start) || termination.After(event.End) {
		return nil, &errors.FailedToValidateTermination
	}

	currentYear, currentMonth, currentDay := event.Start.Date()

	numOccurences, occurencesBetweenErr := recurrence.OccurencesBetween(event.Start, termination)
	if occurencesBetweenErr != nil {
		return nil, &errors.FailedToCalculateOccurences
	}

	events := make([]models.Event, *numOccurences)

	for index := 0; index < *numOccurences; index++ {
		events[index] = *event
		events[index].Start = currentTime
		events[index].End = currentTime.Add(event.End.Sub(event.Start))
		currentTime = currentTime.Add(recurrence.GetDuration())
	}
}
