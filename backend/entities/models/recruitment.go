package models

type Recruitment struct {
	Cycle        RecruitmentCycle `json:"cycle"`
	Type         RecruitmentType  `gorm:"_type" json:"type"`
	IsRecruiting bool             `json:"is_recruiting"`
	Applications []Application    `json:"applications"`
}
