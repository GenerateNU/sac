package models

import (
	"backend/src/types"
)

type Comment struct {
	types.Model

	Question        string `gorm:"type:varchar(255)" json:"question" validate:"required,len<=255"`
	Answer          string `gorm:"type:varchar(255)" json:"answer" validate:",len<=255"`
	NumFoundHelpful uint   `gorm:"type:int;default:0" json:"num_found_helpful" validate:"min=0"`

	AskedByID uint `gorm:"type:uuid" json:"-" validate:"min=1"`
	AskedBy   User `gorm:"foreignKey:AskedByID" json:"-" validate:"-"`

	ClubID uint `gorm:"type:uuid" json:"-" validate:"min=1"`
	Club   Club `gorm:"foreignKey:ClubID" json:"-" validate:"-"`

	AnsweredByID *uint `gorm:"type:uuid" json:"-" validate:"min=1"`
	AnsweredBy   *User `gorm:"foreignKey:AnsweredBy" json:"-" validate:"-"`
}
