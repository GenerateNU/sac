package config

import "time"

type CalendarSettings struct {
	MaxTerminationDate time.Time `yaml:"maxterminationdate"`
}
