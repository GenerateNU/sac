package jobs

import (
	"gorm.io/gorm"
)

type Jobs struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Jobs {
	return &Jobs{db: db}
}
