package contacts

import "github.com/GenerateNU/sac/backend/entities/models"

type PutContactRequestBody struct {
	Type    models.ContactType `json:"type" validate:"required,max=255,oneof=facebook instagram x linkedin youtube github slack discord email customSite"`
	Content string             `json:"content" validate:"required,contact_pointer,max=255"`
}
