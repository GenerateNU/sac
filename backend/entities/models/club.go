package models

import (
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

	SoftDeletedAt gorm.DeletedAt `json:"-" validate:"-"`

	Name                 string           `json:"name" validate:"required,max=255"`
	Preview              string           `json:"preview" validate:"required,max=255"`
	Description          string           `json:"description" validate:"required,http_url,s3_url,max=255"`
	NumMembers           int              `json:"num_members" validate:"required,min=1"`
	IsRecruiting         bool             `json:"is_recruiting" validate:"required"`
	RecruitmentCycle     RecruitmentCycle `json:"recruitment_cycle" validate:"required,max=255,oneof=fall spring fallSpring always"`
	RecruitmentType      RecruitmentType  `json:"recruitment_type" validate:"required,max=255,oneof=unrestricted tryout application"`
	WeeklyTimeCommitment int              `json:"weekly_time_commitment" validate:"omitempty,min=1"`
	OneWordToDescribeUs  string           `json:"one_word_to_describe_us" validate:"omitempty,max=255"`
	ApplicationLink      string           `json:"application_link" validate:"required,max=255,http_url"`
	Logo                 string           `json:"logo" validate:"omitempty,http_url,s3_url,max=255"`

	Parent            *uuid.UUID       `gorm:"foreignKey:Parent;" json:"-" validate:"uuid4"`
	Tag               []Tag            `gorm:"many2many:club_tags;" json:"-" validate:"-"`
	Member            []User           `gorm:"many2many:user_club_members;" json:"-" validate:"required"`
	Follower          []User           `gorm:"many2many:user_club_followers;" json:"-" validate:"-"`
	IntendedApplicant []User           `gorm:"many2many:user_club_intended_applicants;" json:"-" validate:"-"`
	Comment           []Comment        `json:"-" validate:"-"`
	PointOfContact    []PointOfContact `json:"-" validate:"-"`
	Contact           []Contact        `json:"-" validate:"-"`
	// Event
	HostEvent   []Event        `gorm:"foreignKey:Host;" json:"-" validate:"-"`
	Event       []Event        `gorm:"many2many:club_events;" json:"-" validate:"-"`
	Notifcation []Notification `gorm:"polymorphic:Reference;" json:"-" validate:"-"`
}

type CreateClubRequestBody struct {
	UserID               uuid.UUID        `json:"user_id" validate:"required,uuid4"`
	Name                 string           `json:"name" validate:"required,max=255"`
	Preview              string           `json:"preview" validate:"required,max=255"`
	Description          string           `json:"description" validate:"required,http_url,s3_url,max=255"`
	IsRecruiting         bool             `json:"is_recruiting" validate:"required"`
	RecruitmentCycle     RecruitmentCycle `json:"recruitment_cycle" validate:"required,max=255,oneof=fall spring fallSpring always"`
	RecruitmentType      RecruitmentType  `json:"recruitment_type" validate:"required,max=255,oneof=unrestricted tryout application"`
	WeeklyTimeCommitment *int             `json:"weekly_time_commitment" validate:"omitempty,min=1"`
	OneWordToDescribeUs  string           `json:"one_word_to_describe_us" validate:"omitempty,max=255"`
	ApplicationLink      string           `json:"application_link" validate:"required,max=255,http_url"`
	Logo                 string           `json:"logo" validate:"omitempty,http_url,s3_url,max=255"`
}

type UpdateClubRequestBody struct {
	Name                 string           `json:"name" validate:"omitempty,max=255"`
	Preview              string           `json:"preview" validate:"omitempty,max=255"`
	Description          string           `json:"description" validate:"omitempty,http_url,s3_url,max=255"`
	IsRecruiting         bool             `json:"is_recruiting" validate:"omitempty"`
	RecruitmentCycle     RecruitmentCycle `json:"recruitment_cycle" validate:"required,max=255,oneof=fall spring fallSpring always"`
	RecruitmentType      RecruitmentType  `json:"recruitment_type" validate:"required,max=255,oneof=unrestricted tryout application"`
	WeeklyTimeCommitment *int             `json:"weekly_time_commitment" validate:"omitempty,min=1"`
	OneWordToDescribeUs  string           `json:"one_word_to_describe_us" validate:"omitempty,max=255"`
	ApplicationLink      string           `json:"application_link" validate:"omitempty,required,max=255,http_url"`
	Logo                 string           `json:"logo" validate:"omitempty,s3_url,max=255,http_url"`
}

type CreateClubTagsRequestBody struct {
	Tags []uuid.UUID `json:"tags" validate:"required"`
}
