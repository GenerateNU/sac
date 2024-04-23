package models

import (
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	Open        EventType = "open"
	MembersOnly EventType = "membersOnly"
)

type RecurringType string

const (
	Daily   RecurringType = "daily"
	Weekly  RecurringType = "weekly"
	Monthly RecurringType = "monthly"
)

type Event struct {
	Model

	Name        string     `gorm:"type:varchar(255);not null" json:"name" validate:"required,max=255"`
	Preview     string     `gorm:"type:varchar(255);not null" json:"preview" validate:"required,max=255"`
	Content     string     `gorm:"type:text;not null" json:"content" validate:"required,max=255"`
	StartTime   time.Time  `gorm:"type:timestamptz;not null" json:"start_time" validate:"required,ltecsfield=EndTime"`
	EndTime     time.Time  `gorm:"type:timestamptz;not null" json:"end_time" validate:"required,gtecsfield=StartTime"`
	Location    string     `gorm:"type:varchar(255);not null" json:"location" validate:"required,max=255"`
	EventType   EventType  `gorm:"type:varchar(255);default:open;not null" json:"event_type" validate:"required,max=255,oneof=open membersOnly"`
	IsRecurring bool       `gorm:"type:bool;default:false;not null" json:"is_recurring" validate:"-"`
	MeetingLink string     `gorm:"type:varchar(255);" json:"meeting_link" validate:"url"`
	Host        *uuid.UUID `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null;" json:"host" validate:"uuid4"`

	RSVP         []User         `gorm:"many2many:user_event_rsvps;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Waitlist     []User         `gorm:"many2many:user_event_waitlists;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Clubs        []Club         `gorm:"many2many:club_events;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Tag          []Tag          `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Notification []Notification `gorm:"polymorphic:Reference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}

type Series struct {
	Model
	RecurringType  RecurringType `gorm:"type:varchar(255);default:open;not null" json:"recurring_type" validate:"max=255"`
	MaxOccurrences int           `gorm:"type:int;not null" json:"max_occurrences" validate:"min=1"`
	Events         []Event       `gorm:"many2many:event_series;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"events" validate:"-"`
}

type EventSeries struct {
	EventID  uuid.UUID `gorm:"type:uuid;not null" json:"event_id" validate:"uuid4"`
	Event    Event     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	SeriesID uuid.UUID `gorm:"type:uuid;not null" json:"series_id" validate:"uuid4"`
	Series   Series    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}

// Not needed for now, we will just update the events separately
type EventInstanceException struct {
	Model
	EventID       int `gorm:"type:uuid;not null" json:"event_id" validate:"required"`
	Event         Event
	IsRescheduled bool      `gorm:"type:bool;default:true;not null" json:"is_rescheduled" validate:"required"`
	IsCancelled   bool      `gorm:"type:bool;default:false;not null" json:"is_cancelled" validate:"required"`
	StartTime     time.Time `gorm:"type:timestamptz;not null" json:"start_time" validate:"required,datetime,ltecsfield=EndTime"`
	EndTime       time.Time `gorm:"type:timestamptz;not null" json:"end_time" validate:"required,datetime,gtecsfield=StartTime"`
}

// TODO We will likely need to update the create and update structs to account for recurring series
type CreateEventRequestBody struct {
	Name        string    `json:"name" validate:"required,max=255"`
	Preview     string    `json:"preview" validate:"required,max=255"`
	Content     string    `json:"content" validate:"required,max=255"`
	StartTime   time.Time `json:"start_time" validate:"required,ltecsfield=EndTime"`
	EndTime     time.Time `json:"end_time" validate:"required,gtecsfield=StartTime"`
	Location    string    `json:"location" validate:"required,max=255"`
	EventType   EventType `json:"event_type" validate:"required,max=255,oneof=open membersOnly"`
	IsRecurring *bool     `json:"is_recurring" validate:"required"`

	// TODO club/tag/notification logic
	Host         *uuid.UUID     `json:"host" validate:"required,uuid4"`
	Clubs        []Club         `json:"-" validate:"omitempty"`
	Tag          []Tag          `json:"-" validate:"omitempty"`
	Notification []Notification `json:"-" validate:"omitempty"`

	// TODO validate if isRecurring, then series is required
	Series CreateSeriesRequestBody `json:"series" validate:"-"`
}

type CreateSeriesRequestBody struct {
	RecurringType  RecurringType `json:"recurring_type" validate:"required,max=255,oneof=daily weekly monthly"`
	MaxOccurrences int           `json:"max_occurrences" validate:"required,min=2"`
}

type UpdateEventRequestBody struct {
	Name      string    `json:"name" validate:"omitempty,max=255"`
	Preview   string    `json:"preview" validate:"omitempty,max=255"`
	Content   string    `json:"content" validate:"omitempty,max=255"`
	StartTime time.Time `json:"start_time" validate:"omitempty,ltecsfield=EndTime"`
	EndTime   time.Time `json:"end_time" validate:"omitempty,gtecsfield=StartTime"`
	Location  string    `json:"location" validate:"omitempty,max=255"`
	EventType EventType `json:"event_type" validate:"omitempty,max=255,oneof=open membersOnly"`

	Host         *uuid.UUID     `json:"host" validate:"omitempty"`
	RSVP         []User         `json:"-" validate:"omitempty"`
	Waitlist     []User         `json:"-" validate:"omitempty"`
	Clubs        []Club         `json:"-" validate:"omitempty"`
	Tag          []Tag          `json:"-" validate:"omitempty"`
	Notification []Notification `json:"-" validate:"omitempty"`
}

// TODO: probably need to make changes to this to update the events as well
type UpdateSeriesRequestBody struct {
	RecurringType  RecurringType `json:"recurring_type" validate:"omitempty,max=255,oneof=daily weekly monthly"`
	MaxOccurrences int           `json:"max_occurrences" validate:"omitempty,min=2"`

	EventDetails UpdateEventRequestBody `json:"event_details" validate:"omitempty"`
}
