package constants

import "time"

const (
	REDIS_MAX_IDLE_CONNECTIONS int           = 10
	REDIS_MAX_OPEN_CONNECTIONS int           = 100
	RATE_LIMIT_DURATION        time.Duration = 5 * time.Minute
	RATE_LIMIT_MAX_REQUESTS    int           = 5
	REDIS_TIMEOUT              time.Duration = 200 * time.Millisecond
)
