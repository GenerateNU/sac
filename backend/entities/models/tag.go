package models

import "github.com/google/uuid"

type Tag struct {
	Model

	Name string `json:"name"`

	CategoryID uuid.UUID `json:"category_id"`

	User  []User  `gorm:"many2many:user_tags;" json:"-"`
	Club  []Club  `gorm:"many2many:club_tags;" json:"-"`
	Event []Event `gorm:"many2many:event_tags;" json:"-"`
}

type CreateTagRequestBody struct {
	Name       string    `json:"name" validate:"required,max=255"`
	CategoryID uuid.UUID `json:"category_id" validate:"required,uuid4"`
}

type UpdateTagRequestBody struct {
	Name       string    `json:"name" validate:"omitempty,max=255"`
	CategoryID uuid.UUID `json:"category_id" validate:"omitempty,uuid4"`
}
