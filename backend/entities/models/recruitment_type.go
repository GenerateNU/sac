package models

type RecruitmentType string

const (
	RecruitmentTypeUnrestricted RecruitmentType = "unrestricted"
	RecruitmentTypeTryout       RecruitmentType = "tryout"
	RecruitmentTypeApplication  RecruitmentType = "application"
)
