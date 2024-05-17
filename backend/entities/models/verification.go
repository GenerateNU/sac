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

type VerifyEmailRequestBody struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required,len=6"`
}

type VerifyPasswordResetTokenRequestBody struct {
	Token             string `json:"token" validate:"required"`
	NewPassword       string `json:"new_password" validate:"required,min=8,password"`
	VerifyNewPassword string `json:"verify_new_password" validate:"required,min=8,password,eqfield=NewPassword"`
}

type EmailRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}
