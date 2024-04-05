package constants

import "time"

const (
	ACCESS_TOKEN_EXPIRY  time.Duration = time.Minute * 60
	REFRESH_TOKEN_EXPIRY time.Duration = time.Hour * 24 * 30
)
