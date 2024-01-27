package models

import "github.com/google/uuid"

type Tag struct {
	Model

	Name string `gorm:"type:varchar(255)" json:"name" validate:"required,max=255"`

	CategoryID uuid.UUID `gorm:"foreignKey:CategoryID" json:"category_id" validate:"required,uuid4"`

	User  []User  `gorm:"many2many:user_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Club  []Club  `gorm:"many2many:club_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Event []Event `gorm:"many2many:event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}

type TagRequestBody struct {
	Name       string    `json:"name" validate:"required,max=255"`
	CategoryID uuid.UUID `json:"category_id" validate:"required,uuid4"`
}
