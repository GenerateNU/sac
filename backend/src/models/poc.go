package models

import (
	"github.com/google/uuid"
)

type PointOfContact struct {
	Model

	Name     string `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`
	Email    string `gorm:"uniqueIndex:compositeindex;index;not null;type:varchar(255)" json:"email" validate:"required,email,max=255"`
	Position string `gorm:"type:varchar(255);" json:"position" validate:"required,max=255"`

	ClubID uuid.UUID `gorm:"uniqueIndex:compositeindex;index;not null;foreignKey:ClubID" json:"-" validate:"min=1"`
	
	PhotoFile File `gorm:"polymorphic:Owner;" json:"photo_file"`
}

type CreatePointOfContactBody struct {
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Position string `json:"position" validate:"required,max=255"`
}
