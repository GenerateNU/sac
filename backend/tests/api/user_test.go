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

	// Successful update (should return 200)

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
	resp, err := app.App.Test(req)

	var updatedUser models.User
	err = json.NewDecoder(resp.Body).Decode(&updatedUser)
	assert.NilError(err)
	assert.Equal(resp.StatusCode, 200)
	assert.Equal(updatedUser.FirstName, "Michael")
	assert.Equal(updatedUser.LastName, "Brennan")

	// Invalid update values (should return 400)

	// Each entry in invalid_datas represents JSON for a request that should fail (status code 400)
	invalidDatas := []map[string]interface{}{
		// TODO: add the email and password tests in once those validations are complete
		//{"email": "not-northeastern@gmail.com"},
		{"nuid": "1800-123-4567"},
		//{"password": "bad-password"},
		{"year": 1963},
		{"college": "UT-Austin"},
	}

	for i := 0; i < len(invalidDatas); i++ {
		body, err = json.Marshal(invalidDatas[i])
		assert.NilError(err)

		req = httptest.NewRequest(
			"PATCH",
			fmt.Sprintf("%s/api/v1/users/%d", app.Address, user.ID),
			bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err = app.App.Test(req)
		assert.NilError(err)
		assert.Equal(resp.StatusCode, 400)
	}

	// User to update does not exist (should return 400)
	data = map[string]interface{}{
		"first_name": "Michael",
		"last_name":  "Brennan",
	}
	body, err = json.Marshal(data)
	assert.NilError(err)

	req = httptest.NewRequest(
		"PATCH",
		fmt.Sprintf("%s/api/v1/users/%d", app.Address, 12345678),
		bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err = app.App.Test(req)
	assert.NilError(err)
	assert.Equal(resp.StatusCode, 400)
}
