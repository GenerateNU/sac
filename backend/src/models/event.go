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
	EventType EventType `gorm:"type:varchar(255);default:open" json:"event_type" validate:"required"`

	RSVP         []User         `gorm:"many2many:user_event_rsvps;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_rsvps" validate:"-"`
	Waitlist     []User         `gorm:"many2many:user_event_waitlists;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_waitlists" validate:"-"`
	Club         []Club         `gorm:"many2many:club_events;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"clubs" validate:"required"`
	Tag          []Tag          `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags" validate:"-"`
	Notification []Notification `gorm:"polymorphic:Reference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;;" json:"notifications" validate:"-"`
}
