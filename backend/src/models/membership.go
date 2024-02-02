package models

import "github.com/google/uuid"

type MembershipType string

const (
	MembershipTypeMember MembershipType = "member"
	MembershipTypeAdmin  MembershipType = "admin"
)

type Tabler interface {
	TableName() string
}

func (Membership) TableName() string {
	return "user_club_members"
}

type Membership struct {
	Model

	UserID         uuid.UUID       `gorm:"type:uuid;not null" json:"user_id" validate:"required,uuid4"`
	ClubID         uuid.UUID       `gorm:"type:uuid;not null" json:"club_id" validate:"required,uuid4"`
	
	Club *Club `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	User *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	
	MembershipType MembershipType `gorm:"type:varchar(255);not null;default:member" json:"membership_type" validate:"required,oneof=member admin"`
}