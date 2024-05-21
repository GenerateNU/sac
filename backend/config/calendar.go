package config

import (
	"time"
)

type CalendarSettings struct {
	MaxTerminationDate time.Time
}

type intermediateCalendarSettings struct {
	MaxTerminationDate string `yaml:"maxterminationdate"`
}

func (int *intermediateCalendarSettings) into() (*CalendarSettings, error) {
	t, err := time.Parse("01-02-2006", int.MaxTerminationDate)
	if err != nil {
		return nil, err
	}

	return &CalendarSettings{
		MaxTerminationDate: t,
	}, nil
}
