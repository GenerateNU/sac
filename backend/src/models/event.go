package models

import (
	"github.com/google/uuid"
	"time"
)

type EventType string

const (
	Open        EventType = "open"
	MembersOnly EventType = "membersOnly"
)

type RecurringType string

// excluding annually for now bc most clubs have meetings per semester
const (
	Daily   RecurringType = "daily"
	Weekly  RecurringType = "weekly"
	Monthly RecurringType = "monthly"
)

type Event struct {
	Model

	Name        string    `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`
	Preview     string    `gorm:"type:varchar(255)" json:"preview" validate:"required,max=255"`
	Content     string    `gorm:"type:varchar(255)" json:"content" validate:"required,max=255"`
	StartTime   time.Time `gorm:"type:timestamptz" json:"start_time" validate:"required,datetime,ltecsfield=EndTime"`
	EndTime     time.Time `gorm:"type:timestamptz" json:"end_time" validate:"required,datetime,gtecsfield=StartTime"`
	Location    string    `gorm:"type:varchar(255)" json:"location" validate:"required,max=255"`
	EventType   EventType `gorm:"type:varchar(255);default:open" json:"event_type" validate:"required,max=255"`
	IsRecurring bool      `gorm:"type:bool;default:false" json:"is_recurring" validate:"required"`

	ParentEvent  *uuid.UUID     `gorm:"foreignKey:ParentEvent" json:"-" validate:"uuid4"`
	RSVP         []User         `gorm:"many2many:user_event_rsvps;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Waitlist     []User         `gorm:"many2many:user_event_waitlists;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Club         []Club         `gorm:"many2many:club_events;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Tag          []Tag          `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Notification []Notification `gorm:"polymorphic:Reference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;;" json:"-" validate:"-"`
}

type RecurringPattern struct {
	EventID         *uuid.UUID    `gorm:"type:uuid;primary_key; foreignKey:EventID" json:"event_id" validate:"uuid4"`
	RecurringType   RecurringType `gorm:"type:varchar(255);default:open" json:"event_type" validate:"required,max=255"`
	SeparationCount int           `gorm:"type:int" json:"separation_count" validate:"required,min=0"`
	MaxOccurrences  int           `gorm:"type:int" json:"max_occurrences" validate:"required,min=1"`
	DayOfWeek       int           `gorm:"type:int" json:"day_of_week" validate:"required,min=1,max=7"`
	WeekOfMonth     int           `gorm:"type:int" json:"week_of_month" validate:"required,min=1,max=5"`
	DayOfMonth      int           `gorm:"type:int" json:"day_of_month" validate:"required,min=1,max=31"`
}

type EventInstanceException struct {
	Model

	EventID       *uuid.UUID `gorm:"foreignKey:EventID" json:"event_id" validate:"uuid4"`
	IsRescheduled bool       `gorm:"type:bool;default:true" json:"is_rescheduled" validate:"required"`
	IsCancelled   bool       `gorm:"type:bool;default:false" json:"is_cancelled" validate:"required"`
	StartTime     time.Time  `gorm:"type:timestamptz" json:"start_time" validate:"required,datetime,ltecsfield=EndTime"`
	EndTime       time.Time  `gorm:"type:timestamptz" json:"end_time" validate:"required,datetime,gtecsfield=StartTime"`
}

// TODO We will likely need to update the create and update structs to account for recurring series
type CreateEventRequestBody struct {
	Name    string `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`
	Preview string `gorm:"type:varchar(255)" json:"preview" validate:"required,max=255"`
	Content string `gorm:"type:varchar(255)" json:"content" validate:"required,max=255"`
	//TODO validations for start/end/location
	StartTime   time.Time `gorm:"type:timestamptz" json:"start_time" validate:"required,datetime,ltecsfield=EndTime"`
	EndTime     time.Time `gorm:"type:timestamptz" json:"end_time" validate:"required,datetime,gtecsfield=StartTime"`
	Location    string    `gorm:"type:varchar(255)" json:"location" validate:"required,max=255"`
	EventType   EventType `gorm:"type:varchar(255);default:open" json:"event_type" validate:"required,max=255"`
	IsRecurring bool      `gorm:"type:bool;default:false" json:"is_recurring" validate:"required"`

	ParentEvent  *uuid.UUID     `gorm:"foreignKey:ParentEvent" json:"-" validate:"uuid4"`
	Club         []Club         `gorm:"many2many:club_events;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Tag          []Tag          `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Notification []Notification `gorm:"polymorphic:Reference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;;" json:"-" validate:"-"`
}

type UpdateEventRequestBody struct {
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
