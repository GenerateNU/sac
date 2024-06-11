package base

import (
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/google/uuid"
)

type EventPreview struct {
	ID uuid.UUID `json:"id"`

	Title string `json:"title"`

	EventType models.EventType `json:"event_type"`
	Location  string           `json:"location"`
	Link      string           `json:"link"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (ep EventPreview) TableName() string {
	return "events"
}
