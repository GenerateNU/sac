package models

import (
	"backend/src/types"
)

type Tag struct {
	types.Model

	Name string `gorm:"type:varchar(255)" json:"name" validate:"required,len<=255"`

	CategoryID uint `gorm:"foreignKey:CategoryID" json:"category_id" validate:"-"`

	User  []User  `gorm:"many2many:user_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Club  []Club  `gorm:"many2many:club_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Event []Event `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}
