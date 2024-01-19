package tests

import (
	"testing"
)

func TestHealthWorks(t *testing.T) {
	_, assert, resp := RequestTestSetup(t, "GET", "/health", nil)

	assert.Equal(200, resp.StatusCode)
}
