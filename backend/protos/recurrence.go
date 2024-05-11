package protos

import (
	"errors"
	"slices"
	"time"
)

type Cadence int

const (
	Weekly Cadence = iota
	Biweekly
	Monthly
)

type Recurrence struct {
	Cadence            Cadence
	HappensOnMonday    bool
	HappensOnTuesday   bool
	HappensOnWednesday bool
	HappensOnThursday  bool
	HappensOnFriday    bool
	HappensOnSaturday  bool
	HappensOnSunday    bool
}

func NewRecurrence(cadence Cadence, happensOn ...HappensOn) (*Recurrence, error) {
	if len(happensOn) == 0 {
		return nil, errors.New("at least one happensOn function must be provided")
	}

	var r Recurrence
	r.Cadence = cadence
	for _, happens := range happensOn {
		r = happens(r)
	}

	return &r, nil
}

func (r *Recurrence) HappensOn() []time.Weekday {
	var result []time.Weekday
	if r.HappensOnMonday {
		result = append(result, time.Monday)
	}
	if r.HappensOnTuesday {
		result = append(result, time.Tuesday)
	}
	if r.HappensOnWednesday {
		result = append(result, time.Wednesday)
	}
	if r.HappensOnThursday {
		result = append(result, time.Thursday)
	}
	if r.HappensOnFriday {
		result = append(result, time.Friday)
	}
	if r.HappensOnSaturday {
		result = append(result, time.Saturday)
	}
	if r.HappensOnSunday {
		result = append(result, time.Sunday)
	}

	return result
}

type HappensOn func(Recurrence) Recurrence

func Monday(r Recurrence) Recurrence {
	r.HappensOnMonday = true
	return r
}

func Tuesday(r Recurrence) Recurrence {
	r.HappensOnTuesday = true
	return r
}

func Wednesday(r Recurrence) Recurrence {
	r.HappensOnWednesday = true
	return r
}

func Thursday(r Recurrence) Recurrence {
	r.HappensOnThursday = true
	return r
}

func Friday(r Recurrence) Recurrence {
	r.HappensOnFriday = true
	return r
}

func Saturday(r Recurrence) Recurrence {
	r.HappensOnSaturday = true
	return r
}

func Sunday(r Recurrence) Recurrence {
	r.HappensOnSunday = true
	return r
}

type Occurrence struct {
	Start time.Time
	End   time.Time
}

func (r *Recurrence) ReccurUntil(start time.Time, end time.Time, termination time.Time) (result []Occurrence, err error) {
	if start.After(end) {
		return nil, errors.New("start date is after end date")
	}

	if start.After(termination) {
		return nil, errors.New("start date is after termination date")
	}

	if end.After(termination) {
		return nil, errors.New("end date is after termination date")
	}

	if !slices.Contains(r.HappensOn(), start.Weekday()) {
		return nil, errors.New("start date is on a day that the recurrence does not happen on")
	}

	current := start
	currentEnd := end
	for current.Before(termination) {
		if current.After(currentEnd) {
			break
		}

		if r.HappensOnMonday && current.Weekday() == time.Monday {
			result = append(result, Occurrence{current, currentEnd})
		}
		if r.HappensOnTuesday && current.Weekday() == time.Tuesday {
			result = append(result, Occurrence{current, currentEnd})
		}
		if r.HappensOnWednesday && current.Weekday() == time.Wednesday {
			result = append(result, Occurrence{current, currentEnd})
		}
		if r.HappensOnThursday && current.Weekday() == time.Thursday {
			result = append(result, Occurrence{current, currentEnd})
		}
		if r.HappensOnFriday && current.Weekday() == time.Friday {
			result = append(result, Occurrence{current, currentEnd})
		}
		if r.HappensOnSaturday && current.Weekday() == time.Saturday {
			result = append(result, Occurrence{current, currentEnd})
		}
		if r.HappensOnSunday && current.Weekday() == time.Sunday {
			result = append(result, Occurrence{current, currentEnd})
		}

		switch r.Cadence {
		case Weekly:
			current = current.AddDate(0, 0, 7)
			currentEnd = currentEnd.AddDate(0, 0, 7)
		case Biweekly:
			current = current.AddDate(0, 0, 14)
			currentEnd = currentEnd.AddDate(0, 0, 14)
		case Monthly:
			current = current.AddDate(0, 1, 0)
			currentEnd = currentEnd.AddDate(0, 1, 0)
		}
	}

	return result, nil
}

const (
	// housekeeping
	PROTO_VERSION rune = '1'
	PROTO_SIZE    int  = 11

	// indices
	VERSION_INDEX            int = 0
	CADENCE_INDEX            int = 1
	START_DAYS_OF_WEEK_INDEX int = 2
	END_DAYS_OF_WEEK_INDEX   int = 9
)

// PROTO v1
// 0: [proto version]
// 1: [cadence]
// 2: [does this occur on Monday?]
// 3: [does this occur on Tuesday?]
// 4: [does this occur on Wednesday?]
// 5: [does this occur on Thursday?]
// 6: [does this occur on Friday?]
// 7: [does this occur on Saturday?]
// 8: [does this occur on Sunday?]
type StringRecurrence string

func (r *Recurrence) Into() string {
	data := make([]rune, PROTO_SIZE)

	data[VERSION_INDEX] = PROTO_VERSION

	switch r.Cadence {
	case Weekly:
		data[CADENCE_INDEX] = '0'
	case Biweekly:
		data[CADENCE_INDEX] = '1'
	case Monthly:
		data[CADENCE_INDEX] = '2'
	}

	if r.HappensOnMonday {
		data[START_DAYS_OF_WEEK_INDEX] = '1'
	}
	if r.HappensOnTuesday {
		data[START_DAYS_OF_WEEK_INDEX+1] = '1'
	}
	if r.HappensOnWednesday {
		data[START_DAYS_OF_WEEK_INDEX+2] = '1'
	}
	if r.HappensOnThursday {
		data[START_DAYS_OF_WEEK_INDEX+3] = '1'
	}
	if r.HappensOnFriday {
		data[START_DAYS_OF_WEEK_INDEX+4] = '1'
	}
	if r.HappensOnSaturday {
		data[START_DAYS_OF_WEEK_INDEX+5] = '1'
	}
	if r.HappensOnSunday {
		data[START_DAYS_OF_WEEK_INDEX+6] = '1'
	}

	return string(data)
}

func RecurrenceFrom(raw StringRecurrence) (recurrence *Recurrence, err error) {
	runes := []rune(raw)

	if runes[VERSION_INDEX] != PROTO_VERSION {
		return nil, errors.New("invalid version")
	}

	if len(runes) != PROTO_SIZE {
		return nil, errors.New("invalid data length")
	}

	if runes[CADENCE_INDEX] != '0' && runes[CADENCE_INDEX] != '1' && runes[CADENCE_INDEX] != '2' {
		return nil, errors.New("invalid cadence")
	}

	var happensOn []HappensOn
	if runes[START_DAYS_OF_WEEK_INDEX] == '1' {
		happensOn = append(happensOn, Monday)
	}
	if runes[START_DAYS_OF_WEEK_INDEX+1] == '1' {
		happensOn = append(happensOn, Tuesday)
	}
	if runes[START_DAYS_OF_WEEK_INDEX+2] == '1' {
		happensOn = append(happensOn, Wednesday)
	}
	if runes[START_DAYS_OF_WEEK_INDEX+3] == '1' {
		happensOn = append(happensOn, Thursday)
	}
	if runes[START_DAYS_OF_WEEK_INDEX+4] == '1' {
		happensOn = append(happensOn, Friday)
	}
	if runes[START_DAYS_OF_WEEK_INDEX+5] == '1' {
		happensOn = append(happensOn, Saturday)
	}
	if runes[START_DAYS_OF_WEEK_INDEX+6] == '1' {
		happensOn = append(happensOn, Sunday)
	}

	return NewRecurrence(
		Cadence(runes[CADENCE_INDEX]),
		happensOn...,
	)
}
