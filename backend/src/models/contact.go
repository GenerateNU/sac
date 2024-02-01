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

	Type    ContactType `gorm:"type:varchar(255)" json:"type" validate:"required,max=255"`
	Content string      `gorm:"type:varchar(255)" json:"content" validate:"required,contact_pointer,max=255"`

	ClubID uuid.UUID `gorm:"foreignKey:ClubID" json:"-" validate:"uuid4"`
}

type PutContactRequestBody struct {
	Type    Media  `gorm:"type:varchar(255)" json:"type" validate:"required,max=255"`
	Content string `gorm:"type:varchar(255)" json:"content" validate:"required,url,max=255"`
	ClubID  uint   `gorm:"foreignKey:ClubID" json:"-" validate:"min=1"`
}
