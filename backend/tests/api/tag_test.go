package tests

import (
	"backend/src/models"
	"backend/src/transactions"
	"bytes"
	"fmt"
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

	req := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/tags/create", app.Address), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.App.Test(req)

	assert.NilError(err)

	assert.Equal(200, resp.StatusCode)

	var respTag models.Tag

	err = json.NewDecoder(resp.Body).Decode(&respTag)

	assert.NilError(err)

	dbTag, err := transactions.GetTag(app.Conn, respTag.ID)

	assert.NilError(err)

	assert.Equal(dbTag, respTag)
}
