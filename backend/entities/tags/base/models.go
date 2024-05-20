package base

import "github.com/google/uuid"

type CreateTagRequestBody struct {
	Name       string    `json:"name" validate:"required,max=255"`
	CategoryID uuid.UUID `json:"category_id" validate:"required,uuid4"`
}

type UpdateTagRequestBody struct {
	Name       string    `json:"name" validate:"omitempty,max=255"`
	CategoryID uuid.UUID `json:"category_id" validate:"omitempty,uuid4"`
}
