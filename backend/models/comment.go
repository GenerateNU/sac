package models

import "github.com/google/uuid"

type Comment struct {
	Model

	Question        string `gorm:"type:varchar(255);not null" json:"question" validate:"required,max=255"`
	Answer          string `gorm:"type:varchar(255);not null" json:"answer" validate:",max=255"`
	NumFoundHelpful uint   `gorm:"type:int;default:0;not null" json:"num_found_helpful" validate:"min=0"`

	AskedByID uuid.UUID `gorm:"type:uuid;not null" json:"-" validate:"uuid4"`
	AskedBy   User      `gorm:"foreignKey:AskedByID;not null" json:"-" validate:"-"`

	ClubID uuid.UUID `gorm:"type:uuid;not null" json:"-" validate:"uuid4"`
	Club   Club      `gorm:"foreignKey:ClubID;not null" json:"-" validate:"-"`

	AnsweredByID *uuid.UUID `gorm:"type:uuid" json:"-" validate:"uuid4"`
	AnsweredBy   *User      `gorm:"foreignKey:AnsweredBy" json:"-" validate:"-"`
}
