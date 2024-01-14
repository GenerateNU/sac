package models

import "backend/src/types"

type Interest struct {
	types.Model
	Title string `gorm:"type:varchar(255)" json:"title" validate:"required"`
	Icon  rune   `gorm:"type:char" json:"icon" validate:"required"`
}
