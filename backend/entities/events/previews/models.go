package previews

import (
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/google/uuid"
)

type common struct {
	ID        uuid.UUID        `json:"id"`
	Title     string           `json:"title"`
	EventType models.EventType `json:"event_type"`
	Location  string           `json:"location"`
	Link      string           `json:"link"`
	StartTime time.Time        `json:"start_time"`
	EndTime   time.Time        `json:"end_time"`
}

type EventStub struct {
	common
	Host uuid.UUID `json:"host"`
}

func (es *EventStub) Into(tags []models.Tag, hostName string, hostLogo string) *EventPreview {
	return &EventPreview{
		common:   es.common,
		Tags:     tags,
		HostName: hostName,
		HostLogo: hostLogo,
	}
}

func (es EventStub) TableName() string {
	return "events"
}

type EventPreview struct {
	common
	Tags     []models.Tag `json:"tags"`
	HostName string       `json:"host_name"`
	HostLogo string       `json:"host_logo"`
}
