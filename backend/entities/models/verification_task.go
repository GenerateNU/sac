package models

import "github.com/google/uuid"

type VerificationTask struct {
	UserID         uuid.UUID `gorm:"primaryKey"`
	VerificationID uuid.UUID `gorm:"primaryKey"`
	Attempts       int
}
