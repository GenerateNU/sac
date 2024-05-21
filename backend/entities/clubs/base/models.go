package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/google/uuid"
)

type CreateClubRequestBody struct {
	UserID               uuid.UUID               `json:"user_id" validate:"required,uuid4"`
	Name                 string                  `json:"name" validate:"required,max=255"`
	Preview              string                  `json:"preview" validate:"required,max=255"`
	Description          string                  `json:"description" validate:"required,http_url,s3_url,max=255"`
	IsRecruiting         bool                    `json:"is_recruiting" validate:"required"`
	RecruitmentCycle     models.RecruitmentCycle `json:"recruitment_cycle" validate:"required,max=255,oneof=fall spring fallSpring always"`
	RecruitmentType      models.RecruitmentType  `json:"recruitment_type" validate:"required,max=255,oneof=unrestricted tryout application"`
	WeeklyTimeCommitment *int                    `json:"weekly_time_commitment" validate:"omitempty,min=1"`
	OneWordToDescribeUs  string                  `json:"one_word_to_describe_us" validate:"omitempty,max=255"`
	ApplicationLink      string                  `json:"application_link" validate:"required,max=255,http_url"`
	Logo                 string                  `json:"logo" validate:"omitempty,http_url,s3_url,max=255"`
}

type UpdateClubRequestBody struct {
	Name                 string                  `json:"name" validate:"omitempty,max=255"`
	Preview              string                  `json:"preview" validate:"omitempty,max=255"`
	Description          string                  `json:"description" validate:"omitempty,http_url,s3_url,max=255"`
	IsRecruiting         bool                    `json:"is_recruiting" validate:"omitempty"`
	RecruitmentCycle     models.RecruitmentCycle `json:"recruitment_cycle" validate:"required,max=255,oneof=fall spring fallSpring always"`
	RecruitmentType      models.RecruitmentType  `json:"recruitment_type" validate:"required,max=255,oneof=unrestricted tryout application"`
	WeeklyTimeCommitment *int                    `json:"weekly_time_commitment" validate:"omitempty,min=1"`
	OneWordToDescribeUs  string                  `json:"one_word_to_describe_us" validate:"omitempty,max=255"`
	ApplicationLink      string                  `json:"application_link" validate:"omitempty,required,max=255,http_url"`
	Logo                 string                  `json:"logo" validate:"omitempty,s3_url,max=255,http_url"`
}
