package tests

import (
	"backend/src/models"
	"backend/src/transactions"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/goccy/go-json"

	"github.com/huandu/go-assert"
)

func TestGetAllUsersWorks(t *testing.T) {
	assert := assert.New(t)
	app, err := SpawnApp()

	assert.NilError(err)

	_, err = app.InsertSampleUser() // THIS SHOULD BE REPLACED BY INSERTING A USER WITH OUR CREATE ENDPOINT

	assert.NilError(err)

	req := httptest.NewRequest("GET", fmt.Sprintf("%s/api/v1/users", app.Address), nil)

	resp, err := app.App.Test(req)

	assert.NilError(err)

	assert.Equal(200, resp.StatusCode)

	var users []models.User

	err = json.NewDecoder(resp.Body).Decode(&users)

	assert.NilError(err)

	assert.Equal(1, len(users))

	respUser := users[0]

	dbUsers, err := transactions.GetAllUsers(app.Conn)

	assert.NilError(err)

	assert.Equal(1, len(dbUsers))

	dbUser := dbUsers[0]

	assert.Equal(dbUser, respUser)
}
