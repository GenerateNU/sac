package models

import (
	"time"

	"github.com/google/uuid"
)

type Verification struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"primaryKey"`
	Email     string
	Token     string
	ExpiresAt time.Time
}
