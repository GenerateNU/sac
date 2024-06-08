package models

import (
	"time"

	"github.com/google/uuid"
)

type Verification struct {
	UserID    uuid.UUID `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
