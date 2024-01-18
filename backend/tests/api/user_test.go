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
	// initialize the test
	app, assert := InitTest(t)

	// create a GET request to the APP/api/v1/users/ endpoint
	req := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/", app.Address), nil)

	// send the request to the app
	resp, err := app.App.Test(req)

	assert.NilError(err)

	assert.Equal(200, resp.StatusCode)

	// decode the response body into a slice of users
	var users []models.User

	err = json.NewDecoder(resp.Body).Decode(&users)

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

func TestGetUser(t *testing.T) {
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

	// create a GET request to the APP/api/v1/users/:id endpoint
	req2 := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/letters", app.Address), nil)
	resp2, err := app.App.Test(req2)
	defer resp2.Body.Close()
	bodyBytes2, err := io.ReadAll(resp2.Body)
	msg2 := string(bodyBytes2)
	assert.Equal("id must be a positive number", msg2)
	assert.Equal(400, resp2.StatusCode)

	// create a GET request to the APP/api/v1/users/:id endpoint
	req3 := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/null", app.Address), nil)
	resp3, err := app.App.Test(req3)
	defer resp3.Body.Close()
	bodyBytes3, err := io.ReadAll(resp3.Body)
	msg3 := string(bodyBytes3)
	assert.Equal("id must be a positive number", msg3)
	assert.Equal(400, resp3.StatusCode)

	// create a GET request to the APP/api/v1/users/:id endpoint
	req4 := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/100", app.Address), nil)
	resp4, err := app.App.Test(req4)
	defer resp4.Body.Close()
	bodyBytes4, err := io.ReadAll(resp4.Body)
	msg4 := string(bodyBytes4)
	assert.Equal("record not found", msg4)
	assert.Equal(404, resp4.StatusCode)

	// create a GET request to the APP/api/v1/users/:id endpoint
	req5 := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/-1", app.Address), nil)
	resp5, err := app.App.Test(req5)
	defer resp5.Body.Close()
	bodyBytes5, err := io.ReadAll(resp5.Body)
	msg5 := string(bodyBytes5)
	assert.Equal("id must be a positive number", msg5)
	assert.Equal(400, resp5.StatusCode)

	// create a GET request to the APP/api/v1/users/:id endpoint
	req6 := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/0", app.Address), nil)
	resp6, err := app.App.Test(req6)
	defer resp5.Body.Close()
	bodyBytes6, err := io.ReadAll(resp6.Body)
	msg6 := string(bodyBytes6)
	assert.Equal("id must be a positive number", msg6)
	assert.Equal(400, resp6.StatusCode)
}
