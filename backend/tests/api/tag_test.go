package tests

import (
	"backend/src/models"
	"backend/src/transactions"
	"bytes"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"
)

func TestCreateTagWorks(t *testing.T) {
	app, assert := InitTest(t)

	data := map[string]interface{}{
		"name":        "Generate",
		"category_id": 1,
	}

	body, err := json.Marshal(data)

	assert.NilError(err)

	req := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/tags/", app.Address), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.App.Test(req)

	assert.NilError(err)

	assert.Equal(201, resp.StatusCode)

	var respTag models.Tag

	err = json.NewDecoder(resp.Body).Decode(&respTag)

	assert.NilError(err)

	dbTag, err := transactions.GetTag(app.Conn, respTag.ID)

	assert.NilError(err)

	assert.Equal(dbTag, respTag)
}

func TestCreateTagFailsBadRequest(t *testing.T) {
	app, assert := InitTest(t)

	badReqs := []map[string]interface{}{
		{
			"name":        "Generate",
			"category_id": "1",
		},
		{
			"name":        1,
			"category_id": 1,
		},
	}

	for _, badReq := range badReqs {
		body, err := json.Marshal(badReq)

		assert.NilError(err)

		req := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/tags/", app.Address), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.App.Test(req)

		assert.NilError(err)

		assert.Equal(400, resp.StatusCode)

		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)

		assert.NilError(err)

		assert.Equal("Failed to process the request", string(body))
	}
}

func TestCreateTagFailsValidation(t *testing.T) {
	app, assert := InitTest(t)

	badReqs := []map[string]interface{}{
		{
			"name": "Generate",
		},
		{
			"category_id": 1,
		},
		{},
	}

	for _, badReq := range badReqs {
		body, err := json.Marshal(badReq)

		assert.NilError(err)

		req := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/tags/", app.Address), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.App.Test(req)

		assert.NilError(err)

		assert.Equal(400, resp.StatusCode)

		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)

		assert.NilError(err)

		assert.Equal("Failed to validate the data", string(body))
	}
}
