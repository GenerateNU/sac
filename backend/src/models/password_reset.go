package models

import (
	"time"

	"github.com/google/uuid"
)

type PasswordReset struct {
	UserID    uuid.UUID `gorm:"type:varchar(36);not null;primaryKey" json:"user_id" validate:"required,uuid4"`
	Token     string    `gorm:"type:varchar(255);unique" json:"token" validate:"required,max=255"`
	ExpiresAt time.Time `gorm:"type:timestamp;not null;primaryKey" json:"expires_at" validate:"required"`
}

type PasswordResetRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyPasswordResetTokenRequestBody struct {
	Token             string `json:"token" validate:"required"`
	NewPassword       string `json:"new_password" validate:"required,min=8,password"`
	VerifyNewPassword string `json:"verify_new_password" validate:"required,min=8,password,eqfield=NewPassword"`
}
