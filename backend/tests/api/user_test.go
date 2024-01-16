package tests

import (
	"backend/src/models"
	"backend/src/transactions"
	"fmt"
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
