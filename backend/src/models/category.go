package models

import "github.com/GenerateNU/sac/backend/src/types"

type Category struct {
	types.Model

	Name string `gorm:"type:varchar(255);unique" json:"name" validate:"required,max=255"`
	Tag  []Tag  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}

type CategoryRequestBody struct {
	Name string `json:"name" validate:"required,max=255"`
}
