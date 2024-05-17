package models

import (
	"github.com/google/uuid"
)

type PointOfContact struct {
	Model

	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`

	ClubID uuid.UUID `json:"-"`

	PhotoFile File `gorm:"polymorphic:Owner;" json:"photo_file"`
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
