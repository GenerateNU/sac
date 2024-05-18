package models

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Club struct {
	Model

	SoftDeletedAt gorm.DeletedAt `json:"-"`

	Name                 string           `json:"name"`
	Preview              string           `json:"preview"`
	Description          string           `json:"description"`
	NumMembers           int              `json:"num_members"`
	IsRecruiting         bool             `json:"is_recruiting"`
	RecruitmentCycle     RecruitmentCycle `json:"recruitment_cycle"`
	RecruitmentType      RecruitmentType  `json:"recruitment_type"`
	WeeklyTimeCommitment int              `json:"weekly_time_commitment"`
	OneWordToDescribeUs  string           `json:"one_word_to_describe_us"`
	ApplicationLink      string           `json:"application_link"`
	Logo                 string           `json:"logo"`

	Parent            *uuid.UUID       `gorm:"foreignKey:Parent;" json:"-"`
	Tag               []Tag            `gorm:"many2many:club_tags;" json:"-"`
	Member            []User           `gorm:"many2many:user_club_members;" json:"-"`
	Follower          []User           `gorm:"many2many:user_club_followers;" json:"-"`
	IntendedApplicant []User           `gorm:"many2many:user_club_intended_applicants;" json:"-"`
	Comment           []Comment        `json:"-"`
	PointOfContact    []PointOfContact `json:"-"`
	Contact           []Contact        `json:"-"`

	// Event
	HostEvent   []Event        `gorm:"foreignKey:Host;" json:"-"`
	Event       []Event        `gorm:"many2many:club_events;" json:"-"`
	Notifcation []Notification `gorm:"polymorphic:Reference;" json:"-"`
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
