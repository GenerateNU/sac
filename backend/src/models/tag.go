package models

import (
	"backend/src/types"
)

type Tag struct {
	types.Model

	Name string `gorm:"type:varchar(255)" json:"name" validate:"required"`

	CategoryID uint `gorm:"foreignKey:CategoryID" json:"category_id" validate:"required"`

	User  []User  `gorm:"many2many:user_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"users" validate:"-"`
	Club  []Club  `gorm:"many2many:club_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"clubs" validate:"-"`
	Event []Event `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"events" validate:"-"`
}

type TagCreateRequestBody struct {
	Name       string `json:"name" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}
