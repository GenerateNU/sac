package models

import "backend/src/types"

type Category struct {
	types.Model

	Name string `gorm:"type:varchar(255)" json:"category_name" validate:"required"`

	Tag []Tag `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}

type CreateCategoryRequestBody struct {
	Name string `gorm:"type:varchar(255)" json:"category_name" validate:"required"`
}
