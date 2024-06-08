package models

import "github.com/google/uuid"

type VerificationDeletionQueue struct {
	VerificationToken string    `gorm:"primaryKey"`
	UserID            uuid.UUID `gorm:"primaryKey"`
}
