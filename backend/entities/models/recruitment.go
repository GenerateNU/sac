package models

import "github.com/google/uuid"

type Recruitment struct {
	Model

	Cycle RecruitmentCycle `json:"cycle"`
	Type  RecruitmentType  `gorm:"column:_type" json:"type"`

	ClubID uuid.UUID `json:"-"`

	Application []Application `json:"-"`
}

func (r Recruitment) TableName() string {
	return "recruitment"
}
