package models

import (
	"github.com/GenerateNU/sac/backend/src/types"
)

type Tag struct {
	types.Model

	Name string `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`

	CategoryID uint `gorm:"foreignKey:CategoryID" json:"category_id" validate:"required,min=1"`

	User  []User  `gorm:"many2many:user_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Club  []Club  `gorm:"many2many:club_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Event []Event `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}

type PartialTag struct {
	Name       string `json:"name" validate:"required,max=255"`
	CategoryID uint   `json:"category_id" validate:"required,min=1"`
}

type CreateTagRequestBody struct {
	PartialTag
}

type UpdateTagRequestBody struct {
	PartialTag
}
