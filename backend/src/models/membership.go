package models

import (
	"github.com/google/uuid"
)

type MembershipType string

const (
	MembershipTypeMember  MembershipType = "member"
	MembershipTypeOfficer MembershipType = "officer"
)

type MembershipPermission string

const (
	MembershipPermissionRead   MembershipPermission = "read"
	MembershipPermissionCreate MembershipPermission = "create"
)

func (Membership) TableName() string {
	return "user_club_members"
}

type Membership struct {
	UserID uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"user_id" validate:"required,uuid4"`
	ClubID uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"club_id" validate:"required,uuid4"`

	Club *Club `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	User *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`

	RoleName              string                 `gorm:"type:varchar(255);not null;default:member" json:"role_name" validate:"required"`
	MembershipType        MembershipType         `gorm:"type:varchar(255);not null;default:member" json:"membership_type" validate:"required,oneof=member admin"`
	MembershipPermissions []MembershipPermission `gorm:"type:varchar(255)[]" json:"permissions" validate:"required,dive,oneof=read create"`
}
