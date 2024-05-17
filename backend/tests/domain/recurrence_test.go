package tests

import (
	"testing"
	"time"

	"github.com/GenerateNU/sac/backend/protos"
)

func TestRecurrenceUntil(t *testing.T) {
	recurrence, err := protos.NewRecurrence(
		protos.Weekly,
		protos.Wednesday,
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	start := time.Date(2024, 1, 3, 6, 30, 0, 0, time.UTC)      // Wednesday January 3, 2024 6:30 AM
	end := time.Date(2024, 1, 3, 9, 0, 0, 0, time.UTC)         // Wednesday January 3, 2024 9:00 AM
	termination := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC) // Thursday February 1, 2024 12:00 AM

	events, err := recurrence.ReccurUntil(start, end, termination)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := []protos.Occurrence{
		{
			Start: start,
			End:   end,
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
