package models

import (
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

type ClubSearchDocument struct {
	Name        string   `json:"name"`
	Preview     string   `json:"preview"`
	Description string   `json:"description"`
	NumMembers  int      `json:"num_members"`
	Tags        []string `json:"tags"`
}

func (c *Club) ToSearchDocument() interface{} {
	tagIds := make([]string, len(c.Tag))
	for i, tag := range c.Tag {
		tagIds[i] = tag.ID.String()
	}

	return ClubSearchDocument{
		Name:        c.Name,
		Preview:     c.Preview,
		Description: c.Description,
		NumMembers:  c.NumMembers,
		Tags:        tagIds,
	}
}
