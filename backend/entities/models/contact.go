package models

import "github.com/google/uuid"

type ContactType string

const (
	Facebook   ContactType = "facebook"
	Instagram  ContactType = "instagram"
	X          ContactType = "x"
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
		return "https://facebook.com/"
	case Instagram:
		return "https://instagram.com/"
	case X:
		return "https://x.com/"
	case LinkedIn:
		return "https://linkedin.com/"
	case YouTube:
		return "https://youtube.com/"
	case GitHub:
		return "https://github.com/"
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

	Type    ContactType `json:"type"`
	Content string      `json:"content"`

	ClubID uuid.UUID `json:"-"`
}

type PutContactRequestBody struct {
	Type    ContactType `json:"type" validate:"required,max=255,oneof=facebook instagram x linkedin youtube github slack discord email customSite,contact_pointer"`
	Content string      `json:"content" validate:"required,contact_pointer,max=255"`
}
