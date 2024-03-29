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

	SendAt   time.Time `gorm:"type:timestamptz" json:"send_at" validate:"required"`
	Title    string    `gorm:"type:varchar(255)" json:"title" validate:"required,max=255"`
	Content  string    `gorm:"type:varchar(255)" json:"content" validate:"required,max=255"`
	DeepLink string    `gorm:"type:varchar(255)" json:"deep_link" validate:"required,max=255"`
	Icon     string    `gorm:"type:varchar(255)" json:"icon" validate:"required,http_url,max=255"` // S3 URL

	ReferenceID   uuid.UUID        `gorm:"type:int" json:"-" validate:"uuid4"`
	ReferenceType NotificationType `gorm:"type:varchar(255)" json:"-" validate:"max=255"`
}
