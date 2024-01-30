package models

import "github.com/google/uuid"

type Comment struct {
	Model

	Question        string `gorm:"type:varchar(255)" json:"question" validate:"required,max=255"`
	Answer          string `gorm:"type:varchar(255)" json:"answer" validate:",max=255"`
	NumFoundHelpful uint   `gorm:"type:int;default:0" json:"num_found_helpful" validate:"min=0"`

	AskedByID uuid.UUID `gorm:"type:uuid" json:"-" validate:"uuid4"`
	AskedBy   User      `gorm:"foreignKey:AskedByID" json:"-" validate:"-"`

	ClubID uuid.UUID `gorm:"type:uuid" json:"-" validate:"uuid4"`
	Club   Club      `gorm:"foreignKey:ClubID" json:"-" validate:"-"`

	AnsweredByID *uuid.UUID `gorm:"type:uuid" json:"-" validate:"uuid4"`
	AnsweredBy   *User      `gorm:"foreignKey:AnsweredBy" json:"-" validate:"-"`
}
