package models

import (
	"github.com/google/uuid"
)

func (Follower) TableName() string {
	return "user_club_followers"
}

type Follower struct {
	UserID uuid.UUID `gorm:"primaryKey;" json:"user_id"`
	ClubID uuid.UUID `gorm:"primaryKey" json:"club_id"`

	Club *Club `json:"-"`
	User *User `json:"-"`
}
