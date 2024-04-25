package models

import (
	"github.com/google/uuid"
)

func (Follower) TableName() string {
	return "user_club_followers"
}

type Follower struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid4"`
	ClubID uuid.UUID `json:"club_id" validate:"required,uuid4"`

	Club *Club `json:"-" validate:"-"`
	User *User `json:"-" validate:"-"`
}
