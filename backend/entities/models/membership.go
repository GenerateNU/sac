package models

import (
	"github.com/google/uuid"
)

type MembershipType string

const (
	MembershipTypeMember MembershipType = "member"
	MembershipTypeAdmin  MembershipType = "admin"
)

func (Membership) TableName() string {
	return "user_club_members"
}

type Membership struct {
	UserID uuid.UUID `gorm:"primaryKey;" json:"user_id"`
	ClubID uuid.UUID `gorm:"primaryKey;" json:"club_id"`

	Club *Club `json:"-"`
	User *User `json:"-"`

	MembershipType MembershipType `json:"membership_type"`
}
