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

func TestCreateCategoryWorks(t *testing.T) {
	app, assert := InitTest(t)

	data := map[string]interface{}{
		"category_name": "Science",
	}
	body, err := json.Marshal(data)

	assert.NilError(err)

	req := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/category/", app.Address), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.App.Test(req)

	assert.NilError(err)

	assert.Equal(201, resp.StatusCode)

	var respCategory models.Category

	err = json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	dbCategory, err := transactions.GetCategory(app.Conn, respCategory.ID)

	assert.NilError(err)

	assert.Equal(dbCategory, &respCategory)
}

func TestCreateCategoryIgnoresid(t *testing.T) {
	app, assert := InitTest(t)

	data := map[string]interface{}{
		"id":            12,
		"category_name": "Science",
	}

	body, err := json.Marshal(data)

	assert.NilError(err)

	req := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/category/", app.Address), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.App.Test(req)

	assert.NilError(err)

	assert.Equal(201, resp.StatusCode)

	var respCategory models.Category

	err = json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	dbCategory, err := transactions.GetCategory(app.Conn, respCategory.ID)

	assert.NilError(err)

	assert.NotEqual(dbCategory.ID, 12)
}

func TestCreateCategoryFailsIfNameIsNotString(t *testing.T) {
	app, assert := InitTest(t)

	body := map[string]interface{}{
		"category_name": 1231,
	}

	marshalledBody, err := json.Marshal(body)

	assert.NilError(err)

	req := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/category/", app.Address), bytes.NewBuffer(marshalledBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.App.Test(req)

	assert.NilError(err)

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	assert.NilError(err)

	msg := string(bodyBytes)

	assert.Equal("failed to process the request", msg)

	assert.Equal(400, resp.StatusCode)
}

func TestCreateCategoryFailsIfNameIsMissing(t *testing.T) {
	app, assert := InitTest(t)

	data := map[string]interface{}{}
	body, err := json.Marshal(data)

	assert.NilError(err)

	req := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/category/", app.Address), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.App.Test(req)

	assert.NilError(err)

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	assert.NilError(err)

	msg := string(bodyBytes)

	assert.Equal("failed to validate the data", msg)

	assert.Equal(400, resp.StatusCode)
}
