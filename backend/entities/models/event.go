package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

// How Events are represented in the OpenSearch index.
type EventSearchDocument struct {
	Name        string    `json:"name"`
	Preview     string    `json:"preview"`
	Description string    `json:"description"`
	EventType   EventType `json:"event_type"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Clubs       []string  `json:"clubs"`
	Tags        []string  `json:"tags"`
}

func (c *Event) ToSearchDocument() interface{} {
	tagIds := make([]string, len(c.Tag))
	for i, tag := range c.Tag {
		tagIds[i] = tag.ID.String()
	}

	clubIds := make([]string, len(c.Clubs))
	for i, club := range c.Clubs {
		clubIds[i] = club.ID.String()
	}

	return EventSearchDocument{
		Name:        c.Name,
		Preview:     c.Preview,
		Description: c.Description,
		EventType:   c.EventType,
		StartTime:   c.StartTime,
		EndTime:     c.EndTime,
		Tags:        tagIds,
		Clubs:       clubIds,
	}
}

func (c Event) Preload(db *gorm.DB) *gorm.DB {
	return db.Preload("Tag").Preload("Club")
}
