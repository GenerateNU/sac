package constants

import "time"

const (
	DELETE_EXPIRED_VERIFICATION_INTERVAL time.Duration = 30 * time.Minute
	DELETE_EXPIRED_VERIFICATION_LIMIT    int           = 100
	EMAIL_SENDER_INTERVAL                time.Duration = 5 * time.Second
	MAX_EMAIL_ATTEMPTS                   int           = 3
)
