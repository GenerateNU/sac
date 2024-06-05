package utilities

import "time"

type TimeFormat string

const (
	YYYY_dash_MM_dash_DD TimeFormat = "2006-01-02"
)

func ParseTime(s string, format TimeFormat) (time.Time, error) {
	return time.Parse(string(format), s)
}
