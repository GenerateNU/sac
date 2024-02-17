package models

import (
	"fmt"
	"strings"

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
	Tag    []Tag      `gorm:"many2many:club_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags,omitempty" validate:"-"`
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
	IsRecruiting     bool             `json:"is_recruiting" validate:"omitempty"`
	RecruitmentCycle RecruitmentCycle `gorm:"type:varchar(255);default:always" json:"recruitment_cycle" validate:"required,max=255,oneof=fall spring fallSpring always"`
	RecruitmentType  RecruitmentType  `gorm:"type:varchar(255);default:unrestricted" json:"recruitment_type" validate:"required,max=255,oneof=unrestricted tryout application"`
	ApplicationLink  string           `json:"application_link" validate:"omitempty,required,max=255,http_url"`
	Logo             string           `json:"logo" validate:"omitempty,s3_url,max=255,http_url"` // S3 URL
}

type CreateClubTagsRequestBody struct {
	Tags []uuid.UUID `json:"tags" validate:"required"`
}

type ClubQueryParams struct {
	Tags             []string          `query:"tags"`
	MinMembers       int               `query:"min_members"`
	MaxMembers       int               `query:"max_members"`
	RecruitmentCycle *RecruitmentCycle `query:"recruitment_cycle"`
	IsRecruiting     *bool             `query:"is_recruiting"`
	Limit            int               `query:"limit"`
	Page             int               `query:"page"`
}

func (cqp *ClubQueryParams) IntoWhere() string {
	conditions := make([]string, 0)

	if cqp.MinMembers != 0 {
		conditions = append(conditions, fmt.Sprintf("num_members >= %d", cqp.MinMembers))
	}
	if cqp.MaxMembers != 0 {
		conditions = append(conditions, fmt.Sprintf("num_members <= %d", cqp.MaxMembers))
	}
	if cqp.RecruitmentCycle != nil {
		conditions = append(conditions, fmt.Sprintf("recruitment_cycle = '%s'", *cqp.RecruitmentCycle))
	}
	if cqp.IsRecruiting != nil {
		conditions = append(conditions, fmt.Sprintf("is_recruiting = %t", *cqp.IsRecruiting))
	}

	if len(conditions) == 0 {
		return ""
	}
	return "WHERE " + strings.Join(conditions, " AND ")
}

func (c *Club) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&c).Update("num_members", c.NumMembers+1)
	return
}

func (c *Club) AfterDelete(tx *gorm.DB) (err error) {
	tx.Model(&c).Update("num_members", c.NumMembers-1)
	return
}
