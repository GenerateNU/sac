package tests

import (
	"testing"
	"time"

	"github.com/GenerateNU/sac/backend/protos"
)

func TestRecurrenceUntilWeekly(t *testing.T) {
	t.Parallel()

	recurrence, err := protos.NewRecurrence(protos.Weekly, protos.Wednesday)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	start := time.Date(2024, 1, 3, 6, 30, 0, 0, time.UTC)
	end := time.Date(2024, 1, 3, 9, 0, 0, 0, time.UTC)
	termination := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)

	expected := []protos.Occurrence{
		{Start: start, End: end},
		{Start: time.Date(2024, 1, 10, 6, 30, 0, 0, time.UTC), End: time.Date(2024, 1, 10, 9, 0, 0, 0, time.UTC)},
		{Start: time.Date(2024, 1, 17, 6, 30, 0, 0, time.UTC), End: time.Date(2024, 1, 17, 9, 0, 0, 0, time.UTC)},
		{Start: time.Date(2024, 1, 24, 6, 30, 0, 0, time.UTC), End: time.Date(2024, 1, 24, 9, 0, 0, 0, time.UTC)},
		{Start: time.Date(2024, 1, 31, 6, 30, 0, 0, time.UTC), End: time.Date(2024, 1, 31, 9, 0, 0, 0, time.UTC)},
	}

	validateOccurrences(t, recurrence, start, end, termination, expected)
}

func TestRecurrenceUntilBiweekly(t *testing.T) {
	t.Parallel()

	recurrence, err := protos.NewRecurrence(protos.Biweekly, protos.Wednesday)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	start := time.Date(2024, 1, 3, 6, 30, 0, 0, time.UTC)
	end := time.Date(2024, 1, 3, 9, 0, 0, 0, time.UTC)
	termination := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)

	expected := []protos.Occurrence{
		{Start: start, End: end},
		{Start: time.Date(2024, 1, 17, 6, 30, 0, 0, time.UTC), End: time.Date(2024, 1, 17, 9, 0, 0, 0, time.UTC)},
		{Start: time.Date(2024, 1, 31, 6, 30, 0, 0, time.UTC), End: time.Date(2024, 1, 31, 9, 0, 0, 0, time.UTC)},
		{Start: time.Date(2024, 2, 14, 6, 30, 0, 0, time.UTC), End: time.Date(2024, 2, 14, 9, 0, 0, 0, time.UTC)},
		{Start: time.Date(2024, 2, 28, 6, 30, 0, 0, time.UTC), End: time.Date(2024, 2, 28, 9, 0, 0, 0, time.UTC)},
	}

	validateOccurrences(t, recurrence, start, end, termination, expected)
}

func TestRecurrenceUntilMonthly(t *testing.T) {
	t.Parallel()

	recurrence, err := protos.NewRecurrence(protos.Monthly, protos.Wednesday)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	start := time.Date(2024, 1, 3, 6, 30, 0, 0, time.UTC)
	end := time.Date(2024, 1, 3, 9, 0, 0, 0, time.UTC)
	termination := time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC)

	expected := []protos.Occurrence{
		{Start: start, End: end},
		{Start: time.Date(2024, 2, 3, 6, 30, 0, 0, time.UTC), End: time.Date(2024, 2, 3, 9, 0, 0, 0, time.UTC)},
		{Start: time.Date(2024, 3, 3, 6, 30, 0, 0, time.UTC), End: time.Date(2024, 3, 3, 9, 0, 0, 0, time.UTC)},
	}

	validateOccurrences(t, recurrence, start, end, termination, expected)
}

func validateOccurrences(t *testing.T, recurrence *protos.Recurrence, start, end, termination time.Time, expected []protos.Occurrence) {
	events, err := recurrence.ReccurUntil(start, end, termination)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
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
