package tests

import (
	"backend/src/models"
	"backend/src/transactions"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"
)

func TestGetAllUsersWorks(t *testing.T) {
	// setup the test
	app, assert, resp := RequestTester(t, "GET", "/api/v1/users/", nil, nil, nil, nil)
	defer app.DropDB()

	assert.Equal(200, resp.StatusCode)

	// decode the response body into a slice of users
	var users []models.User

	err := json.NewDecoder(resp.Body).Decode(&users)

	assert.NilError(err)

	assert.Equal(1, len(users))

	respUser := users[0]

	// get all users from the database
	dbUsers, err := transactions.GetAllUsers(app.Conn)

	assert.NilError(err)

	assert.Equal(1, len(dbUsers))

	dbUser := dbUsers[0]

	// assert that the user returned from the database is the same as the user returned from the API
	assert.Equal(dbUser, respUser)
}

func TestGetUserHappyPath(t *testing.T) {
	// initialize the test
	app, assert := InitTest(t)

	// create a GET request to the APP/api/v1/users/:id endpoint
	req := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/1", app.Address), nil)

	// send the request to the app
	resp, err := app.App.Test(req)

	assert.NilError(err)

	assert.Equal(200, resp.StatusCode)

	// decode the response body into a slice of users
	var user *models.User

	err = json.NewDecoder(resp.Body).Decode(&user)

	assert.NilError(err)

	dbUser, err := transactions.GetUser(app.Conn, "1")

	assert.NilError(err)

	assert.Equal(dbUser.Email, user.Email)
	assert.Equal(user.Email, "generatesac@gmail.com")
}

func TestGetUserBadRequest(t *testing.T) {
	app, assert := InitTest(t)
	// create a GET request to the APP/api/v1/users/:id endpoint
	req := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/letters", app.Address), nil)
	resp, err := app.App.Test(req)
	assert.NilError(err)
	defer resp.Body.Close()
	bodyBytes, io_err := io.ReadAll(resp.Body)
	assert.NilError(io_err)
	msg := string(bodyBytes)
	assert.Equal("id must be a positive number", msg)
	assert.Equal(400, resp.StatusCode)

	// create a GET request to the APP/api/v1/users/:id endpoint
	req2 := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/null", app.Address), nil)
	resp2, err2 := app.App.Test(req2)
	assert.NilError(err2)
	defer resp2.Body.Close()
	bodyBytes2, io_err2 := io.ReadAll(resp2.Body)
	assert.NilError(io_err2)
	msg2 := string(bodyBytes2)
	assert.Equal("id must be a positive number", msg2)
	assert.Equal(400, resp2.StatusCode)
}
func TestGetUserBadId(t *testing.T) {
	app, assert := InitTest(t)
	// create a GET request to the APP/api/v1/users/:id endpoint
	req1 := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/-1", app.Address), nil)
	resp1, err1 := app.App.Test(req1)
	assert.NilError(err1)
	defer resp1.Body.Close()
	bodyBytes1, err1 := io.ReadAll(resp1.Body)
	msg1 := string(bodyBytes1)
	assert.Equal("id must be a positive number", msg1)
	assert.Equal(400, resp1.StatusCode)

	// create a GET request to the APP/api/v1/users/:id endpoint
	req2 := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/0", app.Address), nil)
	resp2, err2 := app.App.Test(req2)
	assert.NilError(err2)
	defer resp2.Body.Close()
	bodyBytes2, io_err2 := io.ReadAll(resp2.Body)
	assert.NilError(io_err2)
	msg2 := string(bodyBytes2)
	assert.Equal("id must be a positive number", msg2)
	assert.Equal(400, resp2.StatusCode)
}
func TestGetUserNotFound(t *testing.T) {
	app, assert := InitTest(t)
	// create a GET request to the APP/api/v1/users/:id endpoint
	req := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/100", app.Address), nil)
	resp, err := app.App.Test(req)
	assert.NilError(err)
	defer resp.Body.Close()
	bodyBytes, io_err := io.ReadAll(resp.Body)
	assert.NilError(io_err)
	msg := string(bodyBytes)
	assert.Equal("record not found", msg)
	assert.Equal(404, resp.StatusCode)
}
