package models

import "github.com/google/uuid"

type Comment struct {
	Model

	Question        string `json:"question"`
	Answer          string `json:"answer"`
	NumFoundHelpful uint   `json:"num_found_helpful"`

	AskedByID uuid.UUID `json:"-"`
	AskedBy   User      `json:"-"`

	ClubID uuid.UUID `json:"-"`
	Club   Club      `json:"-"`

	AnsweredByID *uuid.UUID `json:"-"`
	AnsweredBy   *User      `json:"-"`
}
