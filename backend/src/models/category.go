package models

import "github.com/GenerateNU/sac/backend/src/types"

type Category struct {
	types.Model

	Name string `gorm:"type:varchar(255)" json:"category_name" validate:"required,max=255"`

	Tag []Tag `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}

type CreateUpdateCategoryRequestBody struct {

type PartialCategory struct {
	Name string `json:"category_name" validate:"required,max=255"`
}

type CreateCategoryRequestBody struct {
	PartialCategory
}

type UpdateCategoryRequestBody struct {
	PartialCategory
}
