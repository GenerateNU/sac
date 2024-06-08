package constants

import "time"

const (
	DELETE_EXPIRED_VERIFICATION_INTERVAL time.Duration = 30 * time.Minute
	DELETE_EXPIRED_VERIFICATION_LIMIT    int           = 100
)
