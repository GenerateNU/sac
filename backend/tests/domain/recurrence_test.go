package tests

import (
	"testing"
	"time"

	"github.com/GenerateNU/sac/backend/protos"
)

func TestRecurrenceUntil(t *testing.T) {
	recurrence, err := protos.NewRecurrence(
		true,  // isWeekly
		false, // isBiweekly
		false, // isYearly
		false, // happensOnMonday
		false, // happensOnTuesday
		true,  // happensOnWednesday
		false, // happensOnThursday
		false, // happensOnFriday
		false, // happensOnSaturday
		false, // happensOnSunday
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	start := time.Date(2024, 1, 1, 6, 30, 0, 0, time.UTC)
	end := time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC)

	termination := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)

	events, err := recurrence.ReccurUntil(start, end, termination)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := []protos.Occurence{
		{
			Start: time.Date(2024, 1, 3, 6, 30, 0, 0, time.UTC),
			End:   time.Date(2024, 1, 3, 9, 0, 0, 0, time.UTC),
		},
		{
			Start: time.Date(2024, 1, 10, 6, 30, 0, 0, time.UTC),
			End:   time.Date(2024, 1, 10, 9, 0, 0, 0, time.UTC),
		},
		{
			Start: time.Date(2024, 1, 17, 6, 30, 0, 0, time.UTC),
			End:   time.Date(2024, 1, 17, 9, 0, 0, 0, time.UTC),
		},
		{
			Start: time.Date(2024, 1, 24, 6, 30, 0, 0, time.UTC),
			End:   time.Date(2024, 1, 24, 9, 0, 0, 0, time.UTC),
		},
		{
			Start: time.Date(2024, 1, 31, 6, 30, 0, 0, time.UTC),
			End:   time.Date(2024, 1, 31, 9, 0, 0, 0, time.UTC),
		},
	}

	if len(events) != len(expected) {
		t.Errorf("expected %d events, got %d", len(expected), len(events))
	}

	for i, event := range events {
		if event.Start != expected[i].Start {
			t.Errorf("expected start time %v, got %v", expected[i].Start, event.Start)
		}
		if event.End != expected[i].End {
			t.Errorf("expected end time %v, got %v", expected[i].End, event.End)
		}
	}
}
