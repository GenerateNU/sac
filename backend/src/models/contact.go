package models

import "github.com/google/uuid"

type ContactType string

const (
	Facebook   ContactType = "facebook"
	Instagram  ContactType = "instagram"
	Twitter    ContactType = "twitter"
	LinkedIn   ContactType = "linkedin"
	YouTube    ContactType = "youtube"
	GitHub     ContactType = "github"
	Slack      ContactType = "slack"
	Discord    ContactType = "discord"
	Email      ContactType = "email"
	CustomSite ContactType = "customSite"
)

type Contact struct {
	Model

	Type    ContactType `gorm:"type:varchar(255);uniqueIndex:idx_contact_type" json:"type" validate:"required,max=255,oneof=facebook instagram twitter linkedin youtube github slack discord email customSite"`
	Content string      `gorm:"type:varchar(255);" json:"content" validate:"required,max=255"`

	ClubID uuid.UUID `gorm:"foreignKey:ClubID;uniqueIndex:idx_contact_type" json:"-" validate:"uuid4"`
}

type PutContactRequestBody struct {
	Type    ContactType  `gorm:"type:varchar(255)" json:"type" validate:"required,max=255,oneof=facebook instagram twitter linkedin youtube github slack discord email customSite"`

	//TODO content  validator
	Content string `gorm:"type:varchar(255)" json:"content" validate:"required,s3_url,max=255"`
}
