package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/huandu/go-assert"

	"github.com/goccy/go-json"
)

func TestGetAllUsersWorks(t *testing.T) {
	TestRequest{
		Method: "GET",
		Path:   "/api/v1/users/",
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
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
			},
		},
	)
}

func TestDeleteUserWorks(t *testing.T) {
	app, assert := InitTest(t)
	user := models.User{
		Role:         models.Student,
		NUID:         "12345678",
		FirstName:    "Bob",
		LastName:     "Dylan",
		Email:        "dylan.b@northeastern.edu",
		PasswordHash: "music",
		College:      models.CAMD,
		Year:         models.Second,
	}
	err := app.Conn.Create(&user).Error
	assert.NilError(err)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/users/%d", app.Address, user.ID), nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.App.Test(req)
	assert.Equal(resp.StatusCode, 200)
}

func TestDeleteUserNotExist(t *testing.T) {
	app, assert := InitTest(t)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/users/%d", app.Address, 1000), nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.App.Test(req)
	assert.Equal(resp.StatusCode, 404)
}

func TestDeleteUserInvalidStringID(t *testing.T) {
	app, assert := InitTest(t)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("%s%s", app.Address, "/api/v1/users/hello"), nil)
	resp, err := app.App.Test(req)
	assert.NilError(err)
	assert.Equal(resp.StatusCode, 400)

	body, err := io.ReadAll(resp.Body)
	assert.NilError(err)
	errorMessage := string(body)
	expectedErrorMessage := "wrong or invalid id"
	assert.Assert(strings.Contains(errorMessage, expectedErrorMessage))
}

func TestDeleteUserInvalidNegativeID(t *testing.T) {
	app, assert := InitTest(t)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("%s%s", app.Address, "/api/v1/users/-1"), nil)
	resp, err := app.App.Test(req)
	assert.NilError(err)
	assert.Equal(resp.StatusCode, 400)

	body, err := io.ReadAll(resp.Body)
	assert.NilError(err)
	errorMessage := string(body)
	expectedErrorMessage := "invalid id"
	assert.Assert(strings.Contains(errorMessage, expectedErrorMessage))
}

func TestDeleteUserDatabaseNotConnected(t *testing.T) {
	app, assert := InitTest(t)

	db, err := app.Conn.DB()
	assert.NilError(err)
	db.Close()

	req := httptest.NewRequest("DELETE", fmt.Sprintf("%s%s", app.Address, "/api/v1/users/1"), nil)
	resp, err := app.App.Test(req)
	assert.NilError(err)
	assert.Equal(resp.StatusCode, 500)

	body, err := io.ReadAll(resp.Body)
	assert.NilError(err)
	errorMessage := string(body)
	expectedErrorMessage := "not connected to database"
	assert.Assert(strings.Contains(errorMessage, expectedErrorMessage))
}
