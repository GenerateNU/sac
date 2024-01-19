package tests

import (
	"testing"
)

func TestHealthWorks(t *testing.T) {
	app, assert, resp := RequestTestSetup(t, "GET", "/health", nil, nil)
	defer app.DropDB()

	assert.Equal(200, resp.StatusCode)
}
