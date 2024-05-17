package models

import (
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	Hybrid   EventType = "hybrid"
	InPerson EventType = "in_person"
	Virtual  EventType = "virtual"
)

type Event struct {
	Model

	// details
	Name        string     `json:"name"`
	Preview     string     `json:"preview"`
	Description string     `json:"description"`
	Host        *uuid.UUID `json:"host"`

	// geoshi
	EventType EventType `json:"event_type"`
	Location  string    `json:"location"`
	Link      string    `json:"link"`

	// internal housekeeping states
	IsPublic   bool `json:"is_public"`
	IsDraft    bool `json:"is_draft"`
	IsArchived bool `json:"is_archived"`

	// timing
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	// series
	SeriesID *uuid.UUID `json:"-"`

	// bridgerton
	RSVP         []User         `gorm:"many2many:user_event_rsvps;" json:"-"`
	Waitlist     []User         `gorm:"many2many:user_event_waitlists;" json:"-"`
	Clubs        []Club         `gorm:"many2many:club_events;" json:"-"`
	Tag          []Tag          `gorm:"many2many:event_tags;" json:"-"`
	Notification []Notification `gorm:"polymorphic:Reference;" json:"-"`
}
