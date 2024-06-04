package config

import (
	"fmt"
	"time"
)

type CalendarSettings struct {
	MaxTerminationDate time.Time
}

type intermediateCalendarSettings struct {
	MaxTerminationDate string `env:"MAX_TERMINATION_DATE"`
}

func (i *intermediateCalendarSettings) into() (*CalendarSettings, error) {
	maxTerminationDate, err := time.Parse("01-02-2006", i.MaxTerminationDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse max termination date: %s", err.Error())
	}

	return &CalendarSettings{
		MaxTerminationDate: maxTerminationDate,
	}, nil
}
