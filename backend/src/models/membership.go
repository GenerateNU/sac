package models

import "github.com/google/uuid"

type MembershipType string

const (
	MembershipTypeMember MembershipType = "member"
	MembershipTypeAdmin MembershipType = "admin"
)

type Tabler interface {
	TableName() string
}

func (Membership) TableName() string {
	return "user_club_members"
}

type Membership struct {
	Model

	ClubID uuid.UUID `gorm:"primaryKey" json:"club_id" validate:"required,uuid4`
	UserID uuid.UUID `gorm:"primaryKey" json:"user_id" validate:"required,uuid4"`
	Type MembershipType `gorm:"type:varchar(255);default:member" json:"membership_type" validate:"required,max=255"`

	Club *Club `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	User *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}
