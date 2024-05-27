package constants

import "time"

const (
	REDIS_MAX_IDLE_CONNECTIONS int           = 10
	REDIS_MAX_OPEN_CONNECTIONS int           = 100
	TOKEN_BLACKLIST_SKEY       string        = "token_blacklist"
	TOKEN_BLACKLIST_KEY        string        = "blacklisted"
	RATE_LIMIT_DURATION        time.Duration = 5 * time.Minute
	RATE_LIMIT_MAX_REQUESTS    int           = 5
)
