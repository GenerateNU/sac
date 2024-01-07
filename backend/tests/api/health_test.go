package tests

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/huandu/go-assert"
)

func TestHealthWorks(t *testing.T) {
	assert := assert.New(t)
	app, err := SpawnApp()

	assert.NilError(err)

	req := httptest.NewRequest("GET", fmt.Sprintf("%s/health", app.Address), nil)

	resp, err := app.App.Test(req)

	assert.NilError(err)

	assert.Equal(200, resp.StatusCode)
}
