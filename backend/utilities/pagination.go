package utilities

import (
	"github.com/garrettladley/fiberpaginate"
	"gorm.io/gorm"
)

func IntoScope(p fiberpaginate.PageInfo, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.Start()).Limit(p.Limit)
	}
}
