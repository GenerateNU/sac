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
	GetEvent(id string) (*models.Event, *errors.Error)
	CreateEvent(eventBody models.CreateEventRequestBody) (*models.Event, *errors.Error)
	UpdateEvent(id string, eventBody models.UpdateEventRequestBody) (*models.Event, *errors.Error)
	DeleteEvent(id string) *errors.Error
}

type EventService struct {
	DB       *gorm.DB
	Validate *validator.Validate
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

func (c *EventService) CreateEvent(eventBody models.CreateEventRequestBody) (*models.Event, *errors.Error) {
	if err := c.Validate.Struct(eventBody); err != nil {
		return nil, &errors.FailedToValidateEvent
	}

	event, err := utilities.MapRequestToModel(eventBody, &models.UpdateEventRequestBody{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	//TODO convert event from type models.UpdateEventRequestBody to models.Event before passing to transactions.CreateEvent
	return transactions.CreateEvent(c.DB, *event)
}

func (c *EventService) GetEvent(id string) (*models.Event, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetEvent(c.DB, *idAsUUID)
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
