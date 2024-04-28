package models

import "github.com/google/uuid"

type Comment struct {
	Model

	Question        string `json:"question" validate:"required,max=255"`
	Answer          string `json:"answer" validate:",max=255"`
	NumFoundHelpful uint   `json:"num_found_helpful" validate:"min=0"`

	AskedByID uuid.UUID `json:"-" validate:"uuid4"`
	AskedBy   User      `json:"-" validate:"-"`

	ClubID uuid.UUID `json:"-" validate:"uuid4"`
	Club   Club      `json:"-" validate:"-"`

	AnsweredByID *uuid.UUID `json:"-" validate:"uuid4"`
	AnsweredBy   *User      `json:"-" validate:"-"`
}
