package models

import "github.com/google/uuid"

type Tag struct {
	Model

	Name string `json:"name" validate:"required,max=255"`

	CategoryID uuid.UUID `json:"category_id" validate:"required,uuid4"`

	User  []User  `gorm:"many2many:user_tags;" json:"-" validate:"-"`
	Club  []Club  `gorm:"many2many:club_tags;" json:"-" validate:"-"`
	Event []Event `gorm:"many2many:event_tags;" json:"-" validate:"-"`
}

type CreateTagRequestBody struct {
	Name       string    `json:"name" validate:"required,max=255"`
	CategoryID uuid.UUID `json:"category_id" validate:"required,uuid4"`
}

type UpdateTagRequestBody struct {
	Name       string    `json:"name" validate:"omitempty,max=255"`
	CategoryID uuid.UUID `json:"category_id" validate:"omitempty,uuid4"`
}
