package series

import (
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/protos"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
)

type EventSeriesServiceInterface interface {
	// cud
	CreateEvent(body models.CreateEventRequestBody) ([]models.Event, *errors.Error)
	UpdateEvent(eventID string, eventBody models.UpdateEventRequestBody) ([]models.Event, *errors.Error)
	DeleteEvent(eventID string) *errors.Error

	// bridgerton getters (do these belong here?)
	// TODO: move to events/tags, like clubs
	GetTags(eventID string) ([]models.Tag, *errors.Error)
}

type EventSeriesService struct {
	types.ServiceParams
}

func NewEventSeriesService(serviceParams types.ServiceParams) EventSeriesService {
	return &EventSeriesService{serviceParams}
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

func (e *EventService) GetEvent(eventID string) (*models.Event, *errors.Error) {
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
		events, seriesID, err := handleEventSeries(event, body.Recurrence, body.Termination, e.Calendar.MaxTerminationDate)
		if err != nil {
			return nil, err
		}

		return CreateEventSeries(e.DB, events, seriesID)
	} else {
		return CreateEvent(e.DB, *event)
	}
}

func handleEventSeries(event *models.Event, binaryRecurrence protos.BinaryRecurrence, termination time.Time, maxTermination time.Time) ([]models.Event, *uuid.UUID, *errors.Error) {
	if termination.After(maxTermination) {
		return nil, nil, &errors.FailedToValidateTermination
	}

	var recurrence protos.Recurrence
	if unmarshallErr := recurrence.UnmarshallBinary(binaryRecurrence); unmarshallErr != nil {
		return nil, nil, &errors.FailedToUnmarshallRecurrence
	}

	occurences, err := recurrence.ReccurUntil(event.Start, event.End, termination)
	if err != nil {
		return nil, nil, &errors.FailedToReccurUntil
	}

	events := make([]models.Event, len(occurences))
	seriesID := uuid.New()

	for index, occurence := range occurences {
		eventCopy := *event
		eventCopy.SeriesID = &seriesID
		eventCopy.Start = occurence.Start
		eventCopy.End = occurence.End
		events[index] = eventCopy
	}

	return events, &seriesID, nil
}
