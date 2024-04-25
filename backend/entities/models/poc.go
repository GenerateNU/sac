package models

import (
	"github.com/google/uuid"
)

type PointOfContact struct {
	Model

	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Position string `json:"position" validate:"required,max=255"`

	ClubID uuid.UUID `json:"-" validate:"required,uuid4"`

	PhotoFile File `json:"photo_file"`
}

type CreatePointOfContactBody struct {
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Position string `json:"position" validate:"required,max=255"`
}

type UpdatePointOfContactBody struct {
	Name     string `json:"name" validate:"omitempty,max=255"`
	Email    string `json:"email" validate:"omitempty,email,max=255"`
	Position string `json:"position" validate:"omitempty,max=255"`
}
