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

type CreateEventRequestBody struct {
	// details
	Name        string    `json:"name" validate:"required,max=255"`
	Preview     string    `json:"preview" validate:"required,max=255"`
	Description string    `json:"description" validate:"required,max=255"`
	HostID      uuid.UUID `json:"host" validate:"uuid4"`

	// geoshi
	EventType EventType `json:"event_type" validate:"required,max=255,oneof=hybrid in_person virtual"`
	Location  string    `json:"location" validate:"max=255,required_if=EventType in_person,required_if=EventType hybrid"`
	Link      string    `json:"link" validate:"http_url,max=255,required_if=EventType virtual,required_if=EventType hybrid"`

	// internal housekeeping states
	IsPublic   *bool `json:"is_public" validate:"required"`
	IsDraft    *bool `json:"is_draft" validate:"required"`
	IsArchived *bool `json:"is_archived" validate:"required"`

	// timing
	StartTime time.Time `json:"start_time" validate:"required,ltefield=EndTime"`
	EndTime   time.Time `json:"end_time" validate:"required,gtefield=StartTime"`
}

func (c *CreateEventRequestBody) Into() *Event {
	return &Event{
		Name:        c.Name,
		Preview:     c.Preview,
		Description: c.Description,
		Host:        &c.HostID,
		EventType:   c.EventType,
		Location:    c.Location,
		Link:        c.Link,
		IsPublic:    *c.IsPublic,
		IsDraft:     *c.IsDraft,
		IsArchived:  *c.IsArchived,
		StartTime:   c.StartTime,
		EndTime:     c.EndTime,
	}
}

type UpdateEventRequestBody struct {
	// details
	Name        string     `json:"name" validate:"omitempty,max=255"`
	Preview     string     `json:"preview" validate:"omitempty,max=255"`
	Description string     `json:"description" validate:"omitempty,max=255"`
	HostID      *uuid.UUID `json:"host" validate:"omitempty,uuid4"`

	// geoshi
	EventType EventType `json:"event_type" validate:"omitempty,max=255,oneof=hybrid in_person virtual"`
	Location  string    `json:"location" validate:"omitempty,max=255,required_if=EventType in_person,required_if=EventType hybrid"`
	Link      string    `json:"link" validate:"omitempty,url,max=255,required_if=EventType virtual,required_if=EventType hybrid"`

	// internal housekeeping states
	IsPublic   bool `json:"is_public" validate:"omitempty"`
	IsDraft    bool `json:"is_draft" validate:"omitempty"`
	IsArchived bool `json:"is_archived" validate:"omitempty"`

	// timing
	StartTime time.Time `json:"start_time" validate:"omitempty,ltecsfield=EndTime"`
	EndTime   time.Time `json:"end_time" validate:"omitempty,gtecsfield=StartTime"`

	UpdateAllInSeries *bool `json:"update_all_in_series" validate:"required"`
}
