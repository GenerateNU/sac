package models

import (
	"backend/src/types"
	"time"
)

type NotificationType string

const (
	EventNotification NotificationType = "event"
	ClubNotification  NotificationType = "club"
)

type Notification struct {
	types.Model

	SendAt   time.Time `gorm:"type:timestamptz" json:"send_at" validate:"required"`
	Title    string    `gorm:"type:varchar(255)" json:"title" validate:"required,len<=255"`
	Content  string    `gorm:"type:varchar(255)" json:"content" validate:"required,len<=255"`
	DeepLink string    `gorm:"type:varchar(255)" json:"deep_link" validate:"required,len<=255"`
	Icon     string    `gorm:"type:varchar(255)" json:"icon" validate:"required,url,len<=255"` // S3 URL

	ReferenceID   uint             `gorm:"type:int" json:"-" validate:"min=1"`
	ReferenceType NotificationType `gorm:"type:varchar(255)" json:"-" validate:"len<=255"`
}
