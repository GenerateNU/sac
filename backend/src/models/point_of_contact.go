package models

import (
	"backend/src/types"
)

type PointOfContact struct {
	types.Model

	Name     string `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required,email"`
	Photo    string `gorm:"type:varchar(255);default:NULL" json:"photo" validate:"url"` // S3 URL, fallback to default logo if null
	Position string `gorm:"type:varchar(255);" json:"position" validate:"required"`

	ClubID uint `gorm:"foreignKey:ClubID" json:"-" validate:"-"`
}
