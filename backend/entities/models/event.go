package models

import (
	"time"

	"github.com/GenerateNU/sac/backend/protos"
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
	Name        string     `json:"name" validate:"required,max=255"`
	Preview     string     `json:"preview" validate:"required,max=255"`
	Description string     `json:"description" validate:"required,max=255"`
	Host        *uuid.UUID `json:"host" validate:"uuid4"`

	// geoshi
	EventType EventType `json:"event_type" validate:"required,max=255,oneof=hybrid in_person virtual"`
	Location  string    `json:"location" validate:"max=255,required_if=EventType in_person,required_if=EventType hybrid"`
	Link      string    `json:"link" validate:"url,max=255,required_if=EventType virtual,required_if=EventType hybrid"`

	// internal housekeeping states
	IsPublic   bool `json:"is_public" validate:"required"`
	IsDraft    bool `json:"is_draft" validate:"required"`
	IsArchived bool `json:"is_archived" validate:"required"`

	// timing
	StartTime time.Time `json:"start_time" validate:"required,ltecsfield=EndTime"`
	EndTime   time.Time `json:"end_time" validate:"required,gtecsfield=StartTime"`

	// series
	SeriesID *uuid.UUID `json:"-" validate:"-"`

	// bridgerton
	RSVP         []User         `gorm:"many2many:user_event_rsvps;" json:"-" validate:"-"`
	Waitlist     []User         `gorm:"many2many:user_event_waitlists;" json:"-" validate:"-"`
	Clubs        []Club         `gorm:"many2many:club_events;" json:"-" validate:"-"`
	Tag          []Tag          `gorm:"many2many:event_tags;" json:"-" validate:"-"`
	Notification []Notification `gorm:"polymorphic:Reference;" json:"-" validate:"-"`
}

type CreateEventRequestBody struct {
	// details
	Name        string    `json:"name" validate:"required,max=255"`
	Preview     string    `json:"preview" validate:"required,max=255"`
	Description string    `json:"description" validate:"required,max=255"`
	HostID      uuid.UUID `json:"host" validate:"uuid4"`

	// geoshi
	EventType EventType `json:"event_type" validate:"required,max=255,oneof=hybrid in_person virtual"`
	Location  string    `json:"location" validate:"max=255,required_if=EventType in_person,required_if=EventType hybrid"`
	Link      string    `json:"link" validate:"url,max=255,required_if=EventType virtual,required_if=EventType hybrid"`

	// internal housekeeping states
	IsPublic   bool `json:"is_public" validate:"required"`
	IsDraft    bool `json:"is_draft" validate:"required"`
	IsArchived bool `json:"is_archived" validate:"required"`

	// timing
	Start time.Time `json:"start_time" validate:"required,ltecsfield=EndTime"`
	End   time.Time `json:"end_time" validate:"required,gtecsfield=StartTime"`
}

type UpdateEventRequestBody struct {
	// details
	Name        string    `json:"name" validate:"omitempty,max=255"`
	Preview     string    `json:"preview" validate:"omitempty,max=255"`
	Description string    `json:"description" validate:"omitempty,max=255"`
	HostID      uuid.UUID `json:"host" validate:"omitempty,uuid4"`

	// geoshi
	EventType EventType `json:"event_type" validate:"omitempty,max=255,oneof=hybrid in_person virtual"`
	Location  string    `json:"location" validate:"omitempty,max=255,required_if=EventType in_person,required_if=EventType hybrid"`
	Link      string    `json:"link" validate:"omitempty,url,max=255,required_if=EventType virtual,required_if=EventType hybrid"`

	// internal housekeeping states
	IsPublic   bool `json:"is_public" validate:"omitempty"`
	IsDraft    bool `json:"is_draft" validate:"omitempty"`
	IsArchived bool `json:"is_archived" validate:"omitempty"`

	// timing
	StartTime  time.Time               `json:"start_time" validate:"omitempty,ltecsfield=EndTime"`
	EndTime    time.Time               `json:"end_time" validate:"omitempty,gtecsfield=StartTime"`
	Recurrence protos.StringRecurrence `json:"recurrence" validate:"omitempty"`

	UpdateAllInSeries bool `json:"update_all_in_series" validate:"required"`
}
