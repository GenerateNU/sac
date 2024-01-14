package models

import (
	"backend/src/types"
	"time"
)

type Notification struct {
	types.Model
	SendAt   time.Time `gorm:"type:timestamptz" json:"send_at" validate:"required"`
	Title    string    `gorm:"type:varchar(255)" json:"title" validate:"required"`
	Content  string    `gorm:"type:varchar(255)" json:"content" validate:"required"`
	DeepLink string    `gorm:"type:varchar(255)" json:"deep_link" validate:"required"`
	Icon     string    `gorm:"type:varchar(255)" json:"icon" validate:"required"` // S3 URI
}
