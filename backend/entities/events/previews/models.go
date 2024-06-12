package previews

import (
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/google/uuid"
)

type EventStub struct {
	ID uuid.UUID `json:"id"`

	Title string `json:"title"`

	EventType models.EventType `json:"event_type"`
	Location  string           `json:"location"`
	Link      string           `json:"link"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	Host uuid.UUID `json:"host"`
}

func (es *EventStub) Into(tags []models.Tag, hostName string, hostLogo string) *EventPreview {
	return &EventPreview{
		ID:        es.ID,
		Title:     es.Title,
		EventType: es.EventType,
		Location:  es.Location,
		Link:      es.Link,
		StartTime: es.StartTime,
		EndTime:   es.EndTime,
		Tags:      tags,
		HostName:  hostName,
		HostLogo:  hostLogo,
	}
}

func (es EventStub) TableName() string {
	return "events"
}

type EventPreview struct {
	ID uuid.UUID `json:"id"`

	Title string `json:"title"`

	EventType models.EventType `json:"event_type"`
	Location  string           `json:"location"`
	Link      string           `json:"link"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	Tags []models.Tag `json:"tags"`

	HostName string `json:"host_name"`
	HostLogo string `json:"host_logo"`
}
