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

	SendAt   time.Time `json:"send_at"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	DeepLink string    `json:"deep_link"`
	Icon     string    `json:"icon"`

	ReferenceID   uuid.UUID        `json:"-"`
	ReferenceType NotificationType `json:"-"`
}
