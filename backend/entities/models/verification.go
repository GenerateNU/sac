package models

import (
	"time"

	"github.com/google/uuid"
)

type VerificationType string

const (
	EmailVerificationType VerificationType = "email_verification"
	PasswordResetType     VerificationType = "password_reset"
)

type Verification struct {
	UserID    uuid.UUID        `json:"user_id"`
	Token     string           `json:"token"`
	ExpiresAt time.Time        `json:"expires_at"`
	Type      VerificationType `json:"type"`
}
