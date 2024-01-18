package models

import (
	"backend/src/types"
)

type Comment struct {
	types.Model

	Question        string `gorm:"type:varchar(255)" json:"question" validate:"required"`
	Answer          string `gorm:"type:varchar(255)" json:"answer" validate:"-"`
	NumFoundHelpful uint   `gorm:"type:int;default:0" json:"num_found_helpful" validate:"-"`

	AskedByID uint `gorm:"type:uuid" json:"-" validate:"-"`
	AskedBy   User `gorm:"foreignKey:AskedByID" json:"-" validate:"-"`

	ClubID uint `gorm:"type:uuid" json:"-" validate:"-"`
	Club   Club `gorm:"foreignKey:ClubID" json:"-" validate:"-"`

	AnsweredByID *uint `gorm:"type:uuid" json:"-" validate:"-"`
	AnsweredBy   *User `gorm:"foreignKey:AnsweredBy" json:"-" validate:"-"`
}
