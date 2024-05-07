package protos

import (
	"errors"
	"time"
)

const (
	// housekeeping
	PROTO_VERSION byte = 1
	PROTO_SIZE    int  = 11

	// indices
	VERSION_INDEX            int = 0
	IS_WEEKLY_INDEX          int = 1
	IS_BIWEEKLY_INDEX        int = 2
	IS_MONTHLY_INDEX         int = 3
	START_DAYS_OF_WEEK_INDEX int = 4
	END_DAYS_OF_WEEK_INDEX   int = 11
)

// PROTO v1
// 0: [proto version]
// 1: [is this weekly?]
// 2: [is this biweekly?]
// 3: [is this monthly?]
// 4: [does this occur on Monday?]
// 5: [does this occur on Tuesday?]
// 6: [does this occur on Wednesday?]
// 7: [does this occur on Thursday?]
// 8: [does this occur on Friday?]
// 9: [does this occur on Saturday?]
// 10: [does this occur on Sunday?]
type BinaryRecurrence []byte

type Recurrence struct {
	IsWeekly           bool
	IsBiweekly         bool
	IsMonthly          bool
	HappensOnMonday    bool
	HappensOnTuesday   bool
	HappensOnWednesday bool
	HappensOnThursday  bool
	HappensOnFriday    bool
	HappensOnSaturday  bool
	HappensOnSunday    bool
}

func NewRecurrence(isWeekly bool, isBiweekly bool, isMonthly bool, happensOnMonday bool, happensOnTuesday bool, happensOnWednesday bool, happensOnThursday bool, happensOnFriday bool, happensOnSaturday bool, happensOnSunday bool) (*Recurrence, error) {
	if !isWeekly && !isBiweekly && !isMonthly {
		return nil, errors.New("recurrence must be weekly, biweekly, or monthly")
	}

	if !happensOnMonday && !happensOnTuesday && !happensOnWednesday && !happensOnThursday && !happensOnFriday && !happensOnSaturday && !happensOnSunday {
		return nil, errors.New("recurrence must happen on at least one day")
	}

	return &Recurrence{
		IsWeekly:           isWeekly,
		IsBiweekly:         isBiweekly,
		IsMonthly:          isMonthly,
		HappensOnMonday:    happensOnMonday,
		HappensOnTuesday:   happensOnTuesday,
		HappensOnWednesday: happensOnWednesday,
		HappensOnThursday:  happensOnThursday,
		HappensOnFriday:    happensOnFriday,
		HappensOnSaturday:  happensOnSaturday,
		HappensOnSunday:    happensOnSunday,
	}, nil
}

type Occurence struct {
	Start time.Time
	End   time.Time
}

func (r *Recurrence) ReccurUntil(start time.Time, end time.Time, termination time.Time) (result []Occurence, err error) {
	if start.After(end) {
		return nil, errors.New("start date is after end date")
	}

	if start.After(termination) {
		return nil, errors.New("start date is after termination date")
	}

	if end.After(termination) {
		return nil, errors.New("end date is after termination date")
	}

	current := start
	for current.Before(termination) {
		if current.After(end) {
			break
		}

		if r.HappensOnMonday && current.Weekday() == time.Monday {
			result = append(result, Occurence{current, end})
		}
		if r.HappensOnTuesday && current.Weekday() == time.Tuesday {
			result = append(result, Occurence{current, end})
		}
		if r.HappensOnWednesday && current.Weekday() == time.Wednesday {
			result = append(result, Occurence{current, end})
		}
		if r.HappensOnThursday && current.Weekday() == time.Thursday {
			result = append(result, Occurence{current, end})
		}
		if r.HappensOnFriday && current.Weekday() == time.Friday {
			result = append(result, Occurence{current, end})
		}
		if r.HappensOnSaturday && current.Weekday() == time.Saturday {
			result = append(result, Occurence{current, end})
		}
		if r.HappensOnSunday && current.Weekday() == time.Sunday {
			result = append(result, Occurence{current, end})
		}

		if r.IsWeekly {
			current = current.AddDate(0, 0, 7)
		} else if r.IsBiweekly {
			current = current.AddDate(0, 0, 14)
		} else if r.IsMonthly {
			current = current.AddDate(0, 1, 0)
		}
	}

	return result, nil
}

func (r *Recurrence) MarshallBinary() (data []byte, err error) {
	data = make([]byte, PROTO_SIZE)

	data[VERSION_INDEX] = PROTO_VERSION

	if r.IsWeekly {
		data[IS_WEEKLY_INDEX] = 1
	}
	if r.IsBiweekly {
		data[IS_BIWEEKLY_INDEX] = 1
	}
	if r.IsMonthly {
		data[IS_MONTHLY_INDEX] = 1
	}
	if r.HappensOnMonday {
		data[START_DAYS_OF_WEEK_INDEX] = 1
	}
	if r.HappensOnTuesday {
		data[START_DAYS_OF_WEEK_INDEX+1] = 1
	}
	if r.HappensOnWednesday {
		data[START_DAYS_OF_WEEK_INDEX+2] = 1
	}
	if r.HappensOnThursday {
		data[START_DAYS_OF_WEEK_INDEX+3] = 1
	}
	if r.HappensOnFriday {
		data[START_DAYS_OF_WEEK_INDEX+4] = 1
	}
	if r.HappensOnSaturday {
		data[START_DAYS_OF_WEEK_INDEX+5] = 1
	}
	if r.HappensOnSunday {
		data[START_DAYS_OF_WEEK_INDEX+6] = 1
	}

	return data, nil
}

func (r *Recurrence) UnmarshallBinary(data []byte) (err error) {
	if data[VERSION_INDEX] != PROTO_VERSION {
		return errors.New("invalid version")
	}

	if len(data) != PROTO_SIZE {
		return errors.New("invalid data length")
	}

	r.IsWeekly = data[IS_WEEKLY_INDEX] == 1
	r.IsBiweekly = data[IS_BIWEEKLY_INDEX] == 1
	r.IsMonthly = data[IS_MONTHLY_INDEX] == 1
	r.HappensOnMonday = data[START_DAYS_OF_WEEK_INDEX] == 1
	r.HappensOnTuesday = data[START_DAYS_OF_WEEK_INDEX+1] == 1
	r.HappensOnWednesday = data[START_DAYS_OF_WEEK_INDEX+2] == 1
	r.HappensOnThursday = data[START_DAYS_OF_WEEK_INDEX+3] == 1
	r.HappensOnFriday = data[START_DAYS_OF_WEEK_INDEX+4] == 1
	r.HappensOnSaturday = data[START_DAYS_OF_WEEK_INDEX+5] == 1
	r.HappensOnSunday = data[START_DAYS_OF_WEEK_INDEX+6] == 1

	return nil
}
