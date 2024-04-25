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
	UserID    uuid.UUID        `gorm:"primaryKey" json:"user_id" validate:"required,uuid4"`
	Token     string           `json:"token" validate:"required,max=255"`
	ExpiresAt time.Time        `gorm:"primaryKey" json:"expires_at" validate:"required"`
	Type      VerificationType `json:"type" validate:"required,oneof=email_verification password_reset"`
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
