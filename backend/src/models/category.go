package models

import "backend/src/types"

type CategoryName string

const (
	ArtsAndDesign                 CategoryName = "artsAndDesign"
	Business                      CategoryName = "business"
	Cultural                      CategoryName = "cultural"
	EquityAndInclusion            CategoryName = "equityAndInclusion"
	PoliticalAndLaw               CategoryName = "politicalAndLaw"
	Religious                     CategoryName = "religious"
	Social                        CategoryName = "social"
	Sports                        CategoryName = "sports"
	Sciences                      CategoryName = "sciences"
	ComputerScienceAndEngineering CategoryName = "computerScienceAndEngineering"
)

type Category struct {
	types.Model

	Name CategoryName `gorm:"type:varchar(255)" json:"category_name" validate:"required"`

	Tag []Tag `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags" validate:"-"`
}
