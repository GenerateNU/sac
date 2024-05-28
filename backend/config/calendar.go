package config

import "time"

type CalendarSettings struct {
	MaxTerminationDate time.Time
}

type IntermediateCalendarSettings struct {
	MaxTerminationDate string `env:"MAX_TERMINATION_DATE"`
}

func (i *IntermediateCalendarSettings) Into() (*CalendarSettings, error) {
	t, err := time.Parse("01-02-2006", i.MaxTerminationDate)
	if err != nil {
		return nil, err
	}

	return &CalendarSettings{
		MaxTerminationDate: t,
	}, nil
}
