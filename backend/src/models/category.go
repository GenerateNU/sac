package models

import "backend/src/types"

type CategoryName string

const (
	Academic     CategoryName = "academic"
	Arts         CategoryName = "arts"
	Business     CategoryName = "business"
	Cultural     CategoryName = "cultural"
	Health       CategoryName = "health"
	Political    CategoryName = "political"
	Professional CategoryName = "professional"
	Religious    CategoryName = "religious"
	Social       CategoryName = "social"
	Sports       CategoryName = "sports"
	Technology   CategoryName = "technology"
	Other        CategoryName = "other"
)

type Category struct {
	types.Model
	Name CategoryName `gorm:"type:category_name" json:"category_name" validate:"required"`
}
