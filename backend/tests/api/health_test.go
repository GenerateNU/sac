package tests

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestHealthWorks(t *testing.T) {
	// initialize the test
	app, assert := InitTest(t)

	// create a GET request to the APP/health endpoint
	req := httptest.NewRequest("GET", fmt.Sprintf("%s/health", app.Address), nil)

	resp, err := app.App.Test(req)

	assert.NilError(err)

	assert.Equal(200, resp.StatusCode)
}
