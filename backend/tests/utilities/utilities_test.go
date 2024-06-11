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

	if parsedTime != time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Expected time to be 2024-06-01, got %v", parsedTime)
	}
}
