package models

import (
	"github.com/GenerateNU/sac/backend/src/types"
)

type PointOfContact struct {
	types.Model

	Name     string `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required,email,max=255"`
	Photo    string `gorm:"type:varchar(255);default:NULL" json:"photo" validate:"url,max=255"` // S3 URL, fallback to default logo if null
	Position string `gorm:"type:varchar(255);" json:"position" validate:"required,max=255"`

	ClubID uint `gorm:"foreignKey:ClubID" json:"-" validate:"min=1"`
}

type CreatePointOfContactBody struct {
	Name     string `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"required,email,max=255"`
	Photo    string `gorm:"type:varchar(255);default:NULL" json:"photo" validate:"url,max=255"` // S3 URL, fallback to default logo if null
	Position string `gorm:"type:varchar(255);" json:"position" validate:"required,max=255"`
}