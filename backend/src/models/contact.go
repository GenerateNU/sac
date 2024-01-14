package models

import (
	"backend/src/types"

	"github.com/google/uuid"
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
	Type    Media     `gorm:"type:media" json:"type" validate:"required"`
	Content string    `gorm:"type:varchar(255)" json:"content" validate:"required"` // media URI
	ClubID  uuid.UUID `gorm:"type:uuid;" json:"club_id" validate:"required"`
}
