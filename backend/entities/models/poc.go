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
