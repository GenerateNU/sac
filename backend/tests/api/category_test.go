package tests

import (
	"backend/src/models"
	"backend/src/transactions"
	"io"
	"testing"

	"github.com/goccy/go-json"
)

func TestCreateCategoryWorks(t *testing.T) {
	app, assert, resp := RequestTesterWithJSONBody(t, "POST", "/api/v1/categories/", &map[string]interface{}{
		"category_name": "Science",
	}, nil, nil, nil)
	defer app.DropDB()

	assert.Equal(201, resp.StatusCode)

	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	dbCategory, err := transactions.GetCategory(app.Conn, respCategory.ID)

	assert.NilError(err)

	assert.Equal(dbCategory, &respCategory)
}

func TestCreateCategoryIgnoresid(t *testing.T) {
	app, assert, resp := RequestTesterWithJSONBody(t, "POST", "/api/v1/categories/", &map[string]interface{}{
		"id":            12,
		"category_name": "Science",
	}, nil, nil, nil)
	defer app.DropDB()

	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	dbCategory, err := transactions.GetCategory(app.Conn, respCategory.ID)

	assert.NilError(err)

	assert.NotEqual(12, dbCategory.ID)
}

func TestCreateCategoryFailsIfNameIsNotString(t *testing.T) {
	app, assert, resp := RequestTesterWithJSONBody(t, "POST", "/api/v1/categories/", &map[string]interface{}{
		"category_name": 1231,
	}, nil, nil, nil)
	defer app.DropDB()

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	assert.NilError(err)

	msg := string(bodyBytes)

	assert.Equal("failed to process the request", msg)

	assert.Equal(400, resp.StatusCode)
}

func TestCreateCategoryFailsIfNameIsMissing(t *testing.T) {
	app, assert, resp := RequestTesterWithJSONBody(t, "POST", "/api/v1/categories/", &map[string]interface{}{}, nil, nil, nil)
	defer app.DropDB()

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	assert.NilError(err)

	msg := string(bodyBytes)

	assert.Equal("failed to validate the data", msg)

	assert.Equal(400, resp.StatusCode)
}

func TestCreateCategoryFailsICategoryWithThatNameAlreadyExists(t *testing.T) {
	app, assert, resp := RequestTesterWithJSONBody(t, "POST", "/api/v1/categories/", &map[string]interface{}{
		"category_name": "Science",
	}, nil, nil, nil)

	assert.Equal(201, resp.StatusCode)

	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	dbCategory, err := transactions.GetCategory(app.Conn, respCategory.ID)

	assert.NilError(err)

	assert.Equal(dbCategory, &respCategory)

	_, _, resp = RequestTesterWithJSONBody(t, "POST", "/api/v1/categories/", &map[string]interface{}{
		"category_name": "Science",
	}, nil, &app, assert)

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	assert.NilError(err)

	msg := string(bodyBytes)

	assert.Equal("category with that name already exists", msg)

	assert.Equal(400, resp.StatusCode)

}
