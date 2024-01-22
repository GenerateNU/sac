package tests

import (
	"bytes"
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
	).Close()
}

var AssertRespUserSameAsDBUser = func(app TestApp, assert *assert.A, resp *http.Response) {
	var respUser models.User

	err := json.NewDecoder(resp.Body).Decode(&respUser)

	assert.NilError(err)

	dbUser, err := transactions.GetUser(app.Conn, respUser.ID)

	assert.NilError(err)

	assert.Equal(dbUser.Role, respUser.Role)
	assert.Equal(dbUser.NUID, respUser.NUID)
	assert.Equal(dbUser.FirstName, respUser.FirstName)
	assert.Equal(dbUser.LastName, respUser.LastName)
	assert.Equal(dbUser.Email, respUser.Email)
	assert.Equal(dbUser.College, respUser.College)
	assert.Equal(dbUser.Year, respUser.Year)
}

func TestGetUserWorks(t *testing.T) {
	TestRequest{
		Method: "GET",
		Path:   "/api/v1/users/1",
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status:   200,
			DBTester: AssertRespUserSameAsDBUser,
		},
	).Close()
}

func TestGetUserFailsBadRequest(t *testing.T) {
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		TestRequest{
			Method: "GET",
			Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: "failed to validate id",
			},
		).Close()
	}
}

func TestGetUserFailsNotFound(t *testing.T) {
	TestRequest{
		Method: "GET",
		Path:   "/api/v1/users/69",
	}.TestOnStatusAndMessage(t, nil,
		MessageWithStatus{
			Status:  404,
			Message: "failed to find tag",
		},
	).Close()
}

<<<<<<< HEAD
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
	assert.Equal(resp.StatusCode, 204)
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
	).Close()
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
		).Close()
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
=======
func TestUpdateUserWorks(t *testing.T) {
	// initialize the test
	app, assert := InitTest(t)

	user := models.User{
		Role:         models.Student,
		NUID:         "123456789",
		FirstName:    "Melody",
		LastName:     "Yu",
		Email:        "melody.yu@northeastern.edu",
		PasswordHash: "rainbows",
		College:      models.KCCS,
		Year:         models.Second,
	}

	err := app.Conn.Create(&user).Error
	assert.NilError(err)

	data := map[string]interface{}{
		"first_name": "Michael",
		"last_name":  "Brennan",
	}
	body, err := json.Marshal(data)
	assert.NilError(err)

	req := httptest.NewRequest(
		"PATCH",
		fmt.Sprintf("%s/api/v1/users/%d", app.Address, user.ID),
		bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.App.Test(req)

	var updatedUser models.User
	err = json.NewDecoder(resp.Body).Decode(&updatedUser)
	assert.NilError(err)
	assert.Equal(resp.StatusCode, 200)
	assert.Equal(updatedUser.FirstName, "Michael")
	assert.Equal(updatedUser.LastName, "Brennan")
}

func TestUpdateUserFailsOnInvalidParams(t *testing.T) {
	// initialize the test
	app, assert := InitTest(t)

	user := models.User{
		Role:         models.Student,
		NUID:         "123456789",
		FirstName:    "Melody",
		LastName:     "Yu",
		Email:        "melody.yu@northeastern.edu",
		PasswordHash: "rainbows",
		College:      models.KCCS,
		Year:         models.Second,
	}

	err := app.Conn.Create(&user).Error
	assert.NilError(err)

	// Each entry in invalid_datas represents JSON for a request that should fail (status code 400)
	invalidDatas := []map[string]interface{}{
		{"email": "not.northeastern@gmail.com"},
		{"nuid": "1800-123-4567"},
		{"password": "1234"},
		{"year": 1963},
		{"college": "UT-Austin"},
	}

	for i := 0; i < len(invalidDatas); i++ {
		body, err := json.Marshal(invalidDatas[i])
		assert.NilError(err)

		req := httptest.NewRequest(
			"PATCH",
			fmt.Sprintf("%s/api/v1/users/%d", app.Address, user.ID),
			bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.App.Test(req)
		assert.NilError(err)
		assert.Equal(resp.StatusCode, 400)
	}
}

func TestUpdateUserFailsOnInvalidId(t *testing.T) {
	// initialize the test
	app, assert := InitTest(t)

	user := models.User{
		Role:         models.Student,
		NUID:         "123456789",
		FirstName:    "Melody",
		LastName:     "Yu",
		Email:        "melody.yu@northeastern.edu",
		PasswordHash: "rainbows",
		College:      models.KCCS,
		Year:         models.Second,
	}

	err := app.Conn.Create(&user).Error
	assert.NilError(err)

	// User to update does not exist (should return 400)
	data := map[string]interface{}{
		"first_name": "Michael",
		"last_name":  "Brennan",
	}
	body, err := json.Marshal(data)
	assert.NilError(err)

	req := httptest.NewRequest(
		"PATCH",
		fmt.Sprintf("%s/api/v1/users/%d", app.Address, 12345678),
		bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.App.Test(req)
	assert.NilError(err)
	assert.Equal(resp.StatusCode, 404)
>>>>>>> de0127b (SAC-5 Update User PATCH (#28))
}
