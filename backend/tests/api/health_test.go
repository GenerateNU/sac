package tests

import (
	"testing"
)

func TestHealthWorks(t *testing.T) {
	TestRequest{
		Method: "GET",
		Path:   "/health",
	}.TestOnStatus(t, nil,
		200,
	).Close()
}
