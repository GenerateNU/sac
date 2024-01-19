package tests

import (
	"backend/src/models"
	"backend/src/transactions"
	"bytes"
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
}
