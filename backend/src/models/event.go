package models

import (
	"backend/src/types"
	"time"
)

type EventType string

const (
	Open        EventType = "open"
	MembersOnly EventType = "membersOnly"
)

type Event struct {
	types.Model
	Name      string    `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Preview   string    `gorm:"type:varchar(255)" json:"preview" validate:"required"`
	Content   string    `gorm:"type:varchar(255)" json:"content" validate:"required"`
	StartTime time.Time `gorm:"type:timestamptz" json:"start_time" validate:"required"`
	EndTime   time.Time `gorm:"type:timestamptz" json:"end_time" validate:"required"`
	Location  string    `gorm:"type:varchar(255)" json:"location" validate:"required"`
	EventType EventType `gorm:"type:event_type;default:open" json:"event_type" validate:"required"`
}
