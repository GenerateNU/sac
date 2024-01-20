package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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
	TestRequest{
		Method: "DELETE",
		Path: "/api/v1/users/1000",
	}.TestOnStatusAndMessage(t, nil,
		MessageWithStatus{
			Status:  404,
			Message: "user not found",
		},
	)
}

func TestDeleteUserBadRequest(t *testing.T) {
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"hello",
		"null",
	}
	for _, badRequest := range badRequests {
		TestRequest{
			Method: "DELETE",
			Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: "failed to validate id",
			},
		)
	}
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
}
