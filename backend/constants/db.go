package constants

import "time"

const (
	MAX_IDLE_CONNECTIONS int           = 10
	MAX_OPEN_CONNECTIONS int           = 100
	DB_TIMEOUT           time.Duration = 200 * time.Millisecond
)
