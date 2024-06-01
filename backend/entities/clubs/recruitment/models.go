package recruitment

import "github.com/GenerateNU/sac/backend/entities/models"

type CreateClubRecruitmentRequestBody struct {
	Cycle models.RecruitmentCycle `json:"cycle" validate:"required,oneof=fall spring fallSpring always"`
	Type  models.RecruitmentType  `json:"type" validate:"required,oneof=unrestricted tryout application"`
}

type CreateApplicationRequestBody struct {
	Title string `json:"title" validate:"required,max=255"`
	Link  string `json:"link" validate:"required,http_url,max=255"`
}

type UpdateClubRecruitmentRequestBody struct {
	Cycle        models.RecruitmentCycle `json:"cycle" validate:"omitempty,oneof=fall spring fallSpring always"`
	Type         models.RecruitmentType  `json:"type" validate:"omitempty,oneof=unrestricted tryout application"`
	IsRecruiting *bool                   `json:"is_recruiting" validate:"omitempty"`
}

type UpdateApplicationRequestBody struct {
	Title string `json:"title" validate:"omitempty,max=255"`
	Link  string `json:"link" validate:"omitempty,http_url,max=255"`
}
