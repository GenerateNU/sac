package models

import (
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	Model

	SoftDeletedAt gorm.DeletedAt `gorm:"type:timestamptz;default:NULL" json:"-" validate:"-"`

	Name             string           `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`
	Preview          string           `gorm:"type:varchar(255)" json:"preview" validate:"required,max=255"`
	Description      string           `gorm:"type:varchar(255)" json:"description" validate:"required,http_url,mongo_url,max=255"` // MongoDB URL
	NumMembers       int              `gorm:"type:int" json:"num_members" validate:"required,min=1"`
	IsRecruiting     bool             `gorm:"type:bool;default:false" json:"is_recruiting" validate:"required"`
	RecruitmentCycle RecruitmentCycle `gorm:"type:varchar(255);default:always" json:"recruitment_cycle" validate:"required,max=255,oneof=fall spring fallSpring always"`
	RecruitmentType  RecruitmentType  `gorm:"type:varchar(255);default:unrestricted" json:"recruitment_type" validate:"required,max=255,oneof=unrestricted tryout application"`
	ApplicationLink  string           `gorm:"type:varchar(255);default:NULL" json:"application_link" validate:"required,max=255,http_url"`
	Logo             string           `gorm:"type:varchar(255);default:NULL" json:"logo" validate:"omitempty,http_url,s3_url,max=255"` // S3 URL

	Parent *uuid.UUID `gorm:"foreignKey:Parent" json:"-" validate:"uuid4"`
	Tag    []Tag      `gorm:"many2many:club_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	// User
	Admin             []User           `gorm:"many2many:user_club_admins;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"required"`
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

type CreateClubRequestBody struct {
	UserID           uuid.UUID        `json:"user_id" validate:"required,uuid4"`
	Name             string           `json:"name" validate:"required,max=255"`
	Preview          string           `json:"preview" validate:"required,max=255"`
	Description      string           `json:"description" validate:"required,http_url,mongo_url,max=255"` // MongoDB URL
	NumMembers       int              `json:"num_members" validate:"required,min=1"`
	IsRecruiting     bool             `json:"is_recruiting" validate:"required"`
	RecruitmentCycle RecruitmentCycle `gorm:"type:varchar(255);default:always" json:"recruitment_cycle" validate:"required,max=255,oneof=fall spring fallSpring always"`
	RecruitmentType  RecruitmentType  `gorm:"type:varchar(255);default:unrestricted" json:"recruitment_type" validate:"required,max=255,oneof=unrestricted tryout application"`
	ApplicationLink  string           `json:"application_link" validate:"required,max=255,http_url"`
	Logo             string           `json:"logo" validate:"omitempty,http_url,s3_url,max=255"` // S3 URL
}

type UpdateClubRequestBody struct {
	Name             string           `json:"name" validate:"omitempty,max=255"`
	Preview          string           `json:"preview" validate:"omitempty,max=255"`
	Description      string           `json:"description" validate:"omitempty,http_url,mongo_url,max=255"` // MongoDB URL
	NumMembers       int              `json:"num_members" validate:"omitempty,min=1"`
	IsRecruiting     bool             `json:"is_recruiting" validate:"omitempty"`
	RecruitmentCycle RecruitmentCycle `gorm:"type:varchar(255);default:always" json:"recruitment_cycle" validate:"required,max=255,oneof=fall spring fallSpring always"`
	RecruitmentType  RecruitmentType  `gorm:"type:varchar(255);default:unrestricted" json:"recruitment_type" validate:"required,max=255,oneof=unrestricted tryout application"`
	ApplicationLink  string           `json:"application_link" validate:"omitempty,required,max=255,http_url"`
	Logo             string           `json:"logo" validate:"omitempty,http_url,s3_url,max=255"` // S3 URL
}
