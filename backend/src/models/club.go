package models

import (
	"backend/src/types"
	"time"

	"github.com/google/uuid"
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
	Accepting   RecruitmentType = "accepting"
	Tryout      RecruitmentType = "tryout"
	Application RecruitmentType = "application"
)

type Club struct {
	types.Model
	SoftDeletedAt    time.Time        `gorm:"type:timestamptz;default:NULL" json:"-" validate:"-"`
	Parent           uuid.UUID        `gorm:"type:uuid;default:NULL" json:"parent" validate:"-"`
	Name             string           `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Preview          string           `gorm:"type:varchar(255)" json:"preview" validate:"required"`
	Description      string           `gorm:"type:varchar(255)" json:"description" validate:"required"`  // MongoDB URI
	NumMembers       int              `gorm:"type:int;default:0" json:"num_members" validate:"required"` // On sac_user join, increment this field in transaction for efficient num_members query
	IsRecruiting     bool             `gorm:"type:bool;default:false" json:"is_recruiting" validate:"required"`
	RecruitmentCycle RecruitmentCycle `gorm:"type:recruitment_cycle;default:always" json:"recruitment_cycle" validate:"required"`
	RecruitmentType  RecruitmentType  `gorm:"type:recruitment_type;default:open" json:"recruitment_type" validate:"required"`
	ApplicationLink  string           `gorm:"type:varchar(255);default:NULL" json:"application_link" validate:"required"`
	Logo             string           `gorm:"type:varchar(255);" json:"logo" validate:"required"` // S3 URI
}
