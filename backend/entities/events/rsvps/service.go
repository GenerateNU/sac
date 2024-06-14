package rsvps

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type Service interface {
	CreateEventRSVP(eventID string, userID string) error
	GetEventRSVPs(eventID string, pageInfo fiberpaginate.PageInfo) ([]models.User, error)
	DeleteEventRSVP(eventID string, userID string) error
}

type Handler struct {
	types.ServiceParams
}

func NewHandler(params types.ServiceParams) Service {
	return &Handler{params}
}

func (h *Handler) CreateEventRSVP(eventID string, userID string) error {
	eventIDAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return err
	}

	userIDAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return err
	}

	return CreateEventRSVP(h.DB, *eventIDAsUUID, *userIDAsUUID)
}

func (h *Handler) GetEventRSVPs(eventID string, pageInfo fiberpaginate.PageInfo) ([]models.User, error) {
	eventIDAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return nil, err
	}

	return GetEventRSVPs(h.DB, *eventIDAsUUID, pageInfo)
}

func (h *Handler) DeleteEventRSVP(eventID string, userID string) error {
	eventIDAsUUID, err := utilities.ValidateID(eventID)
	if err != nil {
		return err
	}

	userIDAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return err
	}

	return DeleteEventRSVP(h.DB, *eventIDAsUUID, *userIDAsUUID)
}
