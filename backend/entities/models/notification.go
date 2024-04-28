package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationType string

const (
	EventNotification NotificationType = "event"
	ClubNotification  NotificationType = "club"
)

type Notification struct {
	Model

	SendAt   time.Time `json:"send_at" validate:"required"`
	Title    string    `json:"title" validate:"required,max=255"`
	Content  string    `json:"content" validate:"required,max=255"`
	DeepLink string    `json:"deep_link" validate:"required,max=255"`
	Icon     string    `json:"icon" validate:"required,s3_url,http_url,max=255"`

	ReferenceID   uuid.UUID        `json:"-" validate:"uuid4"`
	ReferenceType NotificationType `json:"-" validate:"max=255"`
}
