package models

import (
	"backend/src/types"
)

type PointOfContact struct {
	types.Model

	Name     string `gorm:"type:varchar(255)" json:"name" validate:"required,len<=255"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required,email,len<=255"`
	Photo    string `gorm:"type:varchar(255);default:NULL" json:"photo" validate:"url,len<=255"` // S3 URL, fallback to default logo if null
	Position string `gorm:"type:varchar(255);" json:"position" validate:"required,len<=255"`

	ClubID uint `gorm:"foreignKey:ClubID" json:"-" validate:"min=1"`
}
