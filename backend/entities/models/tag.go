package models

import "github.com/google/uuid"

type Tag struct {
	Model

	Name string `json:"name" validate:"required,max=255"`

	CategoryID uuid.UUID `json:"category_id" validate:"required,uuid4"`

	User  []User  `json:"-" validate:"-"`
	Club  []Club  `json:"-" validate:"-"`
	Event []Event `json:"-" validate:"-"`
}

type CreateTagRequestBody struct {
	Name       string    `json:"name" validate:"required,max=255"`
	CategoryID uuid.UUID `json:"category_id" validate:"required,uuid4"`
}

type UpdateTagRequestBody struct {
	Name       string    `json:"name" validate:"omitempty,max=255"`
	CategoryID uuid.UUID `json:"category_id" validate:"omitempty,uuid4"`
}
