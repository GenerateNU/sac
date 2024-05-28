package models

type Recruitment struct {
	Model

	Cycle        RecruitmentCycle `json:"cycle"`
	Type         RecruitmentType  `gorm:"_type" json:"type"`
	IsRecruiting bool             `json:"is_recruiting"`
	Application  []Application    `json:"applications"`
}

func (r Recruitment) TableName() string {
	return "recruitment"
}
