package models

import (
	"github.com/google/uuid"
)

func (Follower) TableName() string {
	return "user_club_followers"
}

type Follower struct {
	UserID uuid.UUID `gorm:"type:uuid;not null;primary_key;unique_index:idx_user_club_members" json:"user_id" validate:"required,uuid4"`
	ClubID uuid.UUID `gorm:"type:uuid;not null;primary_key;unique_index:idx_user_club_members" json:"club_id" validate:"required,uuid4"`

	Club *Club `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	User *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}
