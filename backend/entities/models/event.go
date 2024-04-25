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

	Name        string     `json:"name" validate:"required,max=255"`
	Preview     string     `json:"preview" validate:"required,max=255"`
	Content     string     `json:"content" validate:"required,max=255"`
	StartTime   time.Time  `json:"start_time" validate:"required,ltecsfield=EndTime"`
	EndTime     time.Time  `json:"end_time" validate:"required,gtecsfield=StartTime"`
	Location    string     `json:"location" validate:"required,max=255"`
	EventType   EventType  `json:"event_type" validate:"required,max=255,oneof=open membersOnly"`
	IsRecurring bool       `json:"is_recurring" validate:"-"`
	MeetingLink string     `json:"meeting_link" validate:"url"`
	Host        *uuid.UUID `json:"host" validate:"uuid4"`

	RSVP         []User         `json:"-" validate:"-"`
	Waitlist     []User         `json:"-" validate:"-"`
	Clubs        []Club         `json:"-" validate:"-"`
	Tag          []Tag          `json:"-" validate:"-"`
	Notification []Notification `json:"-" validate:"-"`
}

type Series struct {
	Model
	RecurringType  RecurringType `json:"recurring_type" validate:"max=255"`
	MaxOccurrences int           `json:"max_occurrences" validate:"min=1"`
	Events         []Event       `json:"events" validate:"-"`
}

type EventSeries struct {
	EventID  uuid.UUID `json:"event_id" validate:"uuid4"`
	Event    Event     `json:"-" validate:"-"`
	SeriesID uuid.UUID `json:"series_id" validate:"uuid4"`
	Series   Series    `json:"-" validate:"-"`
}

// Not needed for now, we will just update the events separately
type EventInstanceException struct {
	Model
	EventID       int `json:"event_id" validate:"required"`
	Event         Event
	IsRescheduled bool      `json:"is_rescheduled" validate:"required"`
	IsCancelled   bool      `json:"is_cancelled" validate:"required"`
	StartTime     time.Time `json:"start_time" validate:"required,datetime,ltecsfield=EndTime"`
	EndTime       time.Time `json:"end_time" validate:"required,datetime,gtecsfield=StartTime"`
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
