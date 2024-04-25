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
	UserID uuid.UUID `gorm:"primaryKey;" json:"user_id" validate:"required,uuid4"`
	ClubID uuid.UUID `gorm:"primaryKey;" json:"club_id" validate:"required,uuid4"`

	Club *Club `json:"-" validate:"-"`
	User *User `json:"-" validate:"-"`

	MembershipType MembershipType `json:"membership_type" validate:"required,oneof=member admin"`
}
