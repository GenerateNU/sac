package tests

import (
	"backend/src/models"
	"backend/src/transactions"
	"testing"

	"github.com/goccy/go-json"
)

func TestGetAllUsersWorks(t *testing.T) {
	// setup the test
	app, assert, resp := RequestTestSetup(t, "GET", "/api/v1/users/", nil)
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
