package jobs

import (
	"github.com/GenerateNU/sac/backend/integrations/email"
	"gorm.io/gorm"
)

type Jobs struct {
	db      *gorm.DB
	emailer email.Emailer
}

func New(db *gorm.DB) *Jobs {
	return &Jobs{db: db}
}
