package models

import (
	"github.com/GenerateNU/sac/backend/src/types"
)

type Media string

const (
	Facebook  Media = "facebook"
	Instagram Media = "instagram"
	Twitter   Media = "twitter"
	LinkedIn  Media = "linkedin"
	YouTube   Media = "youtube"
	GitHub    Media = "github"
	Custom    Media = "custom"
)

type Contact struct {
	types.Model

	Type    Media  `gorm:"type:varchar(255)" json:"type" validate:"required,max=255"`
	Content string `gorm:"type:varchar(255)" json:"content" validate:"required,url,max=255"` // media URL

	ClubID uint `gorm:"foreignKey:ClubID" json:"-" validate:"min=1"`
}
