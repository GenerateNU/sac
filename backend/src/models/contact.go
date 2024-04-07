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

func GetContentPrefix(contactType ContactType) string {
	switch contactType {
	case Facebook:
		return "https://www.facebook.com/"
	case Instagram:
		return "https://www.instagram.com/"
	case Twitter:
		return "https://twitter.com/"
	case LinkedIn:
		return "https://www.linkedin.com/"
	case YouTube:
		return "https://www.youtube.com/"
	case GitHub:
		return "https://www.github.com/"
	case Slack:
		return "https://join.slack.com/"
	case Discord:
		return "https://discord.gg/"
	default:
		return ""
	}
}

type Contact struct {
	Model

	Type    ContactType `gorm:"type:varchar(255);uniqueIndex:idx_contact_type" json:"type" validate:"required,max=255,oneof=facebook instagram twitter linkedin youtube github slack discord email customSite"`
	Content string      `gorm:"type:varchar(255)" json:"content" validate:"required,max=255"`

	ClubID uuid.UUID `gorm:"foreignKey:ClubID;uniqueIndex:idx_contact_type" json:"-" validate:"uuid4"`
}

type PutContactRequestBody struct {
	Type    ContactType `json:"type" validate:"required,max=255,oneof=facebook instagram twitter linkedin youtube github slack discord email customSite,contact_pointer"`
	Content string      `json:"content" validate:"required,contact_pointer,max=255"`
}
