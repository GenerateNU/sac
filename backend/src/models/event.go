package models

import (
	"time"
)

type EventType string

const (
	Open        EventType = "open"
	MembersOnly EventType = "membersOnly"
)

type Event struct {
	Model

	Name      string    `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`
	Preview   string    `gorm:"type:varchar(255)" json:"preview" validate:"required,max=255"`
	Content   string    `gorm:"type:varchar(255)" json:"content" validate:"required,max=255"`
	StartTime time.Time `gorm:"type:timestamptz" json:"start_time" validate:"required,datetime,ltecsfield=EndTime"`
	EndTime   time.Time `gorm:"type:timestamptz" json:"end_time" validate:"required,datetime,gtecsfield=StartTime"`
	Location  string    `gorm:"type:varchar(255)" json:"location" validate:"required,max=255"`
	EventType EventType `gorm:"type:varchar(255);default:open" json:"event_type" validate:"required,max=255"`

	RSVP         []User         `gorm:"many2many:user_event_rsvps;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Waitlist     []User         `gorm:"many2many:user_event_waitlists;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Club         []Club         `gorm:"many2many:club_events;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Tag          []Tag          `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Notification []Notification `gorm:"polymorphic:Reference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;;" json:"-" validate:"-"`
}
