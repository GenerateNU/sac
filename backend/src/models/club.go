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

	Name             string           `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`
	Preview          string           `gorm:"type:varchar(255)" json:"preview" validate:"required,max=255"`
	Description      string           `gorm:"type:varchar(255)" json:"description" validate:"required,url,max=255"` // MongoDB URL
	NumMembers       int              `gorm:"type:int;default:0" json:"num_members" validate:"required,min=1"`
	IsRecruiting     bool             `gorm:"type:bool;default:false" json:"is_recruiting" validate:"required"`
	RecruitmentCycle RecruitmentCycle `gorm:"type:varchar(255);default:always" json:"recruitment_cycle" validate:"required,max=255"`
	RecruitmentType  RecruitmentType  `gorm:"type:varchar(255);default:unrestricted" json:"recruitment_type" validate:"required,max=255"`
	ApplicationLink  string           `gorm:"type:varchar(255);default:NULL" json:"application_link" validate:"required,max=255"`
	Logo             string           `gorm:"type:varchar(255);default:NULL" json:"logo" validate:"url,max=255"` // S3 URL

	Parent *uint `gorm:"foreignKey:Parent" json:"-" validate:"min=1"`
	Tag    []Tag `gorm:"many2many:club_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	// User
	Member            []User           `gorm:"many2many:user_club_members;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"required"`
	Follower          []User           `gorm:"many2many:user_club_followers;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	IntendedApplicant []User           `gorm:"many2many:user_club_intended_applicants;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Comment           []Comment        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	PointOfContact    []PointOfContact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Contact           []Contact        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	// Event
	Event       []Event        `gorm:"many2many:club_events;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Notifcation []Notification `gorm:"polymorphic:Reference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}
