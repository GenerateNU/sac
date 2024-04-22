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

	SendAt   time.Time `gorm:"type:timestamptz;not null" json:"send_at" validate:"required"`
	Title    string    `gorm:"type:varchar(255);not null" json:"title" validate:"required,max=255"`
	Content  string    `gorm:"type:varchar(255);not null" json:"content" validate:"required,max=255"`
	DeepLink string    `gorm:"type:varchar(255);not null" json:"deep_link" validate:"required,max=255"`
	Icon     string    `gorm:"type:varchar(255);not null" json:"icon" validate:"required,s3_url,http_url,max=255"`

	ReferenceID   uuid.UUID        `gorm:"type:int;not null" json:"-" validate:"uuid4"`
	ReferenceType NotificationType `gorm:"type:varchar(255);not null" json:"-" validate:"max=255"`
}
