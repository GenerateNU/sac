package models

import "github.com/google/uuid"

type SocialType string

const (
	Facebook   SocialType = "facebook"
	Instagram  SocialType = "instagram"
	X          SocialType = "x"
	LinkedIn   SocialType = "linkedin"
	YouTube    SocialType = "youtube"
	GitHub     SocialType = "github"
	Slack      SocialType = "slack"
	Discord    SocialType = "discord"
	Email      SocialType = "email"
	CustomSite SocialType = "customSite"
)

func GetSocialPrefix(contactType SocialType) string {
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

type Social struct {
	Model

	Type    SocialType `json:"type"`
	Content string     `json:"content"`

	ClubID uuid.UUID `json:"-"`
}
