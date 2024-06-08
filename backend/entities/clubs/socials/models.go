package socials

import "github.com/GenerateNU/sac/backend/entities/models"

type PutSocialRequestBody struct {
	Type    models.SocialType `json:"type" validate:"required,max=255,oneof=facebook instagram x linkedin youtube github slack discord email customSite"`
	Content string            `json:"content" validate:"required,social_pointer,max=255"`
}
