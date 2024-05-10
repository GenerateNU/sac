package series

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
)

type EventSeriesServiceInterface interface {
	CreateEventSeries(body CreateEventSeriesRequestBody) ([]models.Event, error)
	GetEventSeries(seriesID string) ([]models.Event, error)
	DeleteEventSeries(seriesID string) error
}

type EventSeriesService struct {
	types.ServiceParams
}

func NewEventSeriesService(serviceParams types.ServiceParams) EventSeriesServiceInterface {
	return &EventSeriesService{serviceParams}
}

func (e *EventSeriesService) CreateEventSeries(body CreateEventSeriesRequestBody) ([]models.Event, error) {
	if err := utilities.Validate(e.Validate, body); err != nil {
		return nil, err
	}

	validatedBody, err := body.Into()
	if err != nil {
		return nil, err
	}

	if validatedBody.Termination.After(e.Calendar.MaxTerminationDate) {
		return nil, utilities.BadRequest(fmt.Errorf("termination date is after max termination date. termination: %s, max termination: %s", validatedBody.Termination, e.Calendar.MaxTerminationDate))
	}

	rootEvent, err := utilities.MapRequestToModel(validatedBody, &models.Event{})
	if err != nil {
		return nil, err
	}

	occurrences, err := validatedBody.Recurrence.ReccurUntil(rootEvent.Start, rootEvent.End, validatedBody.Termination)
	if err != nil {
		return nil, fmt.Errorf("error recurring event series: %w", err)
	}

	events := make([]models.Event, len(occurrences))
	seriesID := uuid.New()

	for index, occurrence := range occurrences {
		eventOccurence := *rootEvent
		eventOccurence.SeriesID = &seriesID
		eventOccurence.Start = occurence.Start
		eventOccurence.End = occurence.End
		events[index] = eventOccurence
	}

	return events, nil
}

func (e *EventSeriesService) GetEventSeries(seriesID string) ([]models.Event, error) {
	idAsUUID, err := utilities.ValidateID(seriesID)
	if err != nil {
		return nil, err
	}

	return GetEventSeries(e.DB, *idAsUUID)
}

func (e *EventSeriesService) DeleteEventSeries(seriesID string) error {
	idAsUUID, err := utilities.ValidateID(seriesID)
	if err != nil {
		return err
	}

	return DeleteEventSeries(e.DB, *idAsUUID)
}
