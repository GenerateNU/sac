package models

import "backend/src/types"

type Tag struct {
	types.Model
	Name string `gorm:"type:varchar(255)" json:"name" validate:"required"`
}
