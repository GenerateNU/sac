package rsvps

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
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

// handle if event is private and user is not member
// handle if event is full, then return appropriate http status code
// to accomplish ^, add new rsvp capacity and waitlist capacity fields to event. they should be optional and provide a reasonable default
func (h *Handler) CreateEventRSVP(eventID string, userID string) error {
	return nil
}

func (h *Handler) GetEventRSVPs(eventID string, pageInfo fiberpaginate.PageInfo) ([]models.User, error) {
	return nil, nil
}

func (h *Handler) DeleteEventRSVP(eventID string, userID string) error {
	return nil
}
