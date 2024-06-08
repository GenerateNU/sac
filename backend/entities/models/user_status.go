package models

type UserStatus string

const (
	UserStatusPendingVerification UserStatus = "pending_verification"
	UserStatusVerified            UserStatus = "verified"
)
