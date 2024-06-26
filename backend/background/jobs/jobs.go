package jobs

import (
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/integrations/email"
	"gorm.io/gorm"
)

type Jobs struct {
	db      *gorm.DB
	emailer email.Emailer
	search  *config.SearchSettings
}

func New(db *gorm.DB, settings *config.Settings) *Jobs {
	return &Jobs{db: db, search: &settings.Search}
}
