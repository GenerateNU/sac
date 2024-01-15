package models

import (
	"backend/src/types"
)

type PointOfContact struct {
	types.Model

	Name     string `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required"`
	Photo    string `gorm:"type:varchar(255);default:NULL" json:"-" validate:"-"` // S3 URI, fallback to default logo if null
	Position string `gorm:"type:varchar(255);" json:"position" validate:"required"`

	ClubID uint `gorm:"foreignKey:ClubID" json:"-" validate:"-"`
}
