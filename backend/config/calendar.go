package config

import "time"

type CalendarSettings struct {
	MaxTerminationDate time.Time
}

type intermediateCalendarSettings struct {
	MaxTerminationDate string `env:"MAX_TERMINATION_DATE"`
}

func (i *intermediateCalendarSettings) into() (*CalendarSettings, error) {
	t, err := time.Parse("01-02-2006", i.MaxTerminationDate)
	if err != nil {
		return nil, err
	}

	return &CalendarSettings{
		MaxTerminationDate: t,
	}, nil
}
