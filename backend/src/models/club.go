package models

import (
	"backend/src/types"
	"time"
)

type RecruitmentCycle string

const (
	Fall       RecruitmentCycle = "fall"
	Spring     RecruitmentCycle = "spring"
	FallSpring RecruitmentCycle = "fallSpring"
	Always     RecruitmentCycle = "always"
)

type RecruitmentType string

const (
	Unrestricted RecruitmentType = "unrestricted"
	Tryout       RecruitmentType = "tryout"
	Application  RecruitmentType = "application"
)

type Club struct {
	types.Model

	SoftDeletedAt time.Time `gorm:"type:timestamptz;default:NULL" json:"-" validate:"-"`

	Name             string           `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Preview          string           `gorm:"type:varchar(255)" json:"preview" validate:"required"`
	Description      string           `gorm:"type:varchar(255)" json:"description" validate:"required"` // MongoDB URI
	NumMembers       int              `gorm:"type:int;default:0" json:"num_members" validate:"required"`
	IsRecruiting     bool             `gorm:"type:bool;default:false" json:"is_recruiting" validate:"required"`
	RecruitmentCycle RecruitmentCycle `gorm:"type:varchar(255);default:always" json:"recruitment_cycle" validate:"required"`
	RecruitmentType  RecruitmentType  `gorm:"type:varchar(255);default:open" json:"recruitment_type" validate:"required"`
	ApplicationLink  string           `gorm:"type:varchar(255);default:NULL" json:"application_link" validate:"required"`
	Logo             string           `gorm:"type:varchar(255);" json:"logo" validate:"required"` // S3 URI

	Parent *uint `gorm:"foreignKey:Parent" json:"parent_club" validate:"-"`
	Tag    []Tag `gorm:"many2many:club_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags" validate:"-"`
	// User
	Member            []User           `gorm:"many2many:user_club_members;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"club_members" validate:"required"`
	Follower          []User           `gorm:"many2many:user_club_followers;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"club_followers" validate:"-"`
	IntendedApplicant []User           `gorm:"many2many:user_club_intended_applicants;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"club_intended_applicants" validate:"-"`
	Comment           []Comment        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments" validate:"-"`
	PointOfContact    []PointOfContact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"point_of_contacts" validate:"required"`
	Contact           []Contact        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"contacts" validate:"required"`
	// Event
	Event       []Event        `gorm:"many2many:club_events;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"events" validate:"-"`
	Notifcation []Notification `gorm:"polymorphic:Reference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"notifications" validate:"-"`
}