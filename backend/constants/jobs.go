package constants

import "time"

const (
	DELETE_VERIFICATION_INTERVAL         time.Duration = 30 * time.Second
	DELETE_EXPIRED_VERIFICATION_INTERVAL time.Duration = 30 * time.Minute
	DELETE_EXPIRED_TOKEN_LIMIT           int           = 100
)
