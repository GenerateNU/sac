package models

import (
	"backend/src/types"

	"github.com/google/uuid"
)

type PointOfContact struct {
	types.Model
	Name     string    `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Email    string    `gorm:"type:varchar(255)" json:"email" validate:"required"`
	Profile  string    `gorm:"type:varchar(255);default:NULL" json:"-" validate:"-"` // S3 URI, fallback to default logo if null
	Position string    `gorm:"type:varchar(255);" json:"position" validate:"required"`
	ClubID   uuid.UUID `gorm:"type:uuid;" json:"club_id" validate:"required"`
}
