package tests

import (
	"backend/src/models"
	"backend/src/transactions"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
)

// CREATE CATEGORY:
func TestCreateCategory(t *testing.T) {
	app, assert := InitTest(t)

	// SUCCESS: a category is created when passed valid data:
	data1 := map[string]interface{}{
		"category_name": "Science",
	}
	body1, err := json.Marshal(data1)

	assert.NilError(err)

	req1 := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/category/", app.Address), bytes.NewBuffer(body1))
	req1.Header.Set("Content-Type", "application/json")
	resp1, err := app.App.Test(req1)

	assert.NilError(err)
	assert.Equal(201, resp1.StatusCode)

	var respCategory models.Category
	err = json.NewDecoder(resp1.Body).Decode(&respCategory)

	assert.NilError(err)

	dbCategory, err := transactions.GetCategory(app.Conn, respCategory.ID)

	assert.NilError(err)
	assert.Equal(dbCategory, &respCategory)

	// SUCCESS: a category ignores the id passed to it on creation:
	data2 := map[string]interface{}{
		"id": 12,
		"category_name": "Science",
	}
	body2, err := json.Marshal(data2)
	req2 := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/category/", app.Address), bytes.NewBuffer(body2))
	req2.Header.Set("Content-Type", "application/json")
	resp2, err := app.App.Test(req2)

	assert.NilError(err)
	assert.Equal(201, resp2.StatusCode)

	err = json.NewDecoder(resp2.Body).Decode(&respCategory)

	assert.NilError(err)

	dbCategory, err = transactions.GetCategory(app.Conn, respCategory.ID)

	assert.NotEqual(dbCategory.ID, 12)

	// FAILURE: it will fail if the category name is not a string:
	data3 := map[string]interface{}{
		"category_name": 1231,
	}
	body3, err := json.Marshal(data3)
	req3 := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/category/", app.Address), bytes.NewBuffer(body3))
	req3.Header.Set("Content-Type", "application/json")
	resp3, err := app.App.Test(req3)
	defer resp3.Body.Close()
	bodyBytes, err := io.ReadAll(resp3.Body)
	msg := string(bodyBytes)

	assert.Equal(msg, "Failed to process the request")
	assert.Equal(resp3.StatusCode, 400)

	// FAILURE: it will fail if the the request body is missing category_name:
	data4 := map[string]interface{}{}
	body4, err := json.Marshal(data4)
	req4 := httptest.NewRequest("POST", fmt.Sprintf("%s/api/v1/category/", app.Address), bytes.NewBuffer(body4))
	req4.Header.Set("Content-Type", "application/json")
	resp4, err := app.App.Test(req4)
	defer resp3.Body.Close()
	bodyBytes, err = io.ReadAll(resp4.Body)
	msg = string(bodyBytes)

	assert.Equal(msg, "Failed to validate the data")
	assert.Equal(resp3.StatusCode, 400)
}
