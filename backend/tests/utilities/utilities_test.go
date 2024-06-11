package tests

import (
	"testing"
	"time"

	"github.com/GenerateNU/sac/backend/utilities"
)

func TestParseTime(t *testing.T) {
	t.Parallel()

	parsedTime, err := utilities.ParseTime("2024-06-01", utilities.YYYY_dash_MM_dash_DD)
	if err != nil {
		t.Errorf("Error parsing time: %v", err)
	}

	if *parsedTime != time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Expected time to be 2024-06-01, got %v", parsedTime)
	}
}

func TestParseTimeEmptyString(t *testing.T) {
	t.Parallel()

	parsedTime, err := utilities.ParseTime("", utilities.YYYY_dash_MM_dash_DD)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	if parsedTime != nil {
		t.Errorf("Expected nil, got %v", parsedTime)
	}

	if err.Error() != "parsing time \"\" as \"2006-01-02\": cannot parse \"\" as \"2006\"" {
		t.Errorf("Expected error message to be 'time string is empty', got %v", err)
	}
}
