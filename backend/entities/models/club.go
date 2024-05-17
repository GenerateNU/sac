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
	UserID               uuid.UUID        `json:"user_id" validate:"required,uuid4"` // TODO: huh?
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

type ClubQueryParams struct {
	Tags                    []string          `query:"tags"`
	MinMembers              int               `query:"min_members"`
	MaxMembers              int               `query:"max_members"`
	RecruitmentCycle        *RecruitmentCycle `query:"recruitment_cycle"`
	IsRecruiting            *bool             `query:"is_recruiting"`
	MinWeeklyTimeCommitment int               `query:"min_weekly_time_commitment"`
	MaxWeeklyTimeCommitment int               `query:"max_weekly_time_commitment"`
	OneWordToDescribeUs     string            `query:"one_word_to_describe_us"`
	Search                  string            `query:"search"`
}

type ClubSearch struct {
	SearchString string `query:"search"`
}

func NewClubSearch(searchQuery string) *ClubSearch {
	return &ClubSearch{
		SearchString: searchQuery,
	}
}

// dummy searchID
func (cs *ClubSearch) SearchId() string {
	return ""
}

func (cs *ClubSearch) Namespace() string {
	return "clubs"
}

func (cs *ClubSearch) EmbeddingString() string {
	return cs.SearchString
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
	if cqp.MinWeeklyTimeCommitment != 0 {
		conditions = append(conditions, fmt.Sprintf("weekly_time_commitment >= %d", cqp.MinWeeklyTimeCommitment))
	}
	if cqp.MaxWeeklyTimeCommitment != 0 {
		conditions = append(conditions, fmt.Sprintf("weekly_time_commitment <= %d", cqp.MaxWeeklyTimeCommitment))
	}
	if cqp.OneWordToDescribeUs != "" {
		conditions = append(conditions, fmt.Sprintf("one_word_to_describe_us = '%s'", cqp.OneWordToDescribeUs))
	}

	if len(conditions) == 0 {
		return ""
	}
	return strings.Join(conditions, " AND ")
}

func (c *Club) SearchId() string {
	return c.ID.String()
}

func (c *Club) Namespace() string {
	return "clubs"
}

func (c *Club) EmbeddingString() string {
	return fmt.Sprintf("%s %s %s %s %s", c.Name, c.Name, c.Name, c.Name, c.Description)
}

type Clubs []Club

func (c Clubs) String(i int) string {
	return c[i].Name
}

func (c Clubs) Len() int {
	return len(c)
}
