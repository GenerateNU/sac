package tests

import (
	"fmt"
	"math/rand"
	"net/http"
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

func CreateSampleUser(t *testing.T, email string, nuid string) ExistingAppAssert {
	return TestRequest{
		Method: "POST",
		Path:   "/api/v1/users/",
		Body: &map[string]interface{}{
			"first_name": "TestX",
			"last_name":  "TestY",
			"email":      email,
			"password":   "1234567890",
			"nuid":       nuid,
			"college":    "CAMD",
			"year":       3,
		},
	}.TestOnStatusAndDBKeepDB(t, nil,
		DBTesterWithStatus{
			Status: 201,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {

				var respUser models.User

				err := json.NewDecoder(resp.Body).Decode(&respUser)

				assert.NilError(err)

				dbUser, err := transactions.GetUser(app.Conn, respUser.ID)
				assert.NilError(err)

				// This is done because password hash is ommitted in response
				respUser.PasswordHash = dbUser.PasswordHash

				assert.Equal(dbUser, &respUser)
			},
		},
	)
}

func TestCreateUserWorks(t *testing.T) {
	appAssert := CreateSampleUser(t, "test@northeastern.edu", "001159263")
	appAssert.App.DropDB()
}

func TestCreateUserFailsIfCategoryWithEmailAlreadyExists(t *testing.T) {
	email := "test@northeastern.edu"

	existingAppAssert := CreateSampleUser(t, "test@northeastern.edu", "001159263")

	for _, permutation := range AllCasingPermutations(email) {
		fmt.Println(permutation)
		var numberRunes = []rune("1234567890")

		nuid_arr := make([]rune, 9)

		for i := range nuid_arr {
			nuid_arr[i] = numberRunes[rand.Intn(len(numberRunes))]
		}
		nuid := string(nuid_arr)

		TestRequest{
			Method: "POST",
			Path:   "/api/v1/users/",
			Body: &map[string]interface{}{
				"first_name": "TestX",
				"last_name":  "TestY",
				"email":      email,
				"password":   "1234567890",
				"nuid":       nuid,
				"college":    "CAMD",
				"year":       3,
			},
		}.TestOnStatusAndMessageKeepDB(t, &existingAppAssert,
			MessageWithStatus{
				Status:  400,
				Message: "user with that email already exists",
			},
		)
	}

	existingAppAssert.App.DropDB()
}


func TestCreateUserFailsIfCategoryWithNUIDAlreadyExists(t *testing.T) {

	existingAppAssert := CreateSampleUser(t, "test@northeastern.edu", "001159263")

	TestRequest{
		Method: "POST",
		Path:   "/api/v1/users/",
		Body: &map[string]interface{}{
			"first_name": "TestX",
			"last_name":  "TestY",
			"email":      "test2@northeastern.edu",
			"password":   "1234567890",
			"nuid":       "001159263",
			"college":    "CAMD",
			"year":       3,
		},
	}.TestOnStatusAndMessageKeepDB(t, &existingAppAssert,
		MessageWithStatus{
			Status:  400,
			Message: "user with that nuid already exists",
		},
	)

	existingAppAssert.App.DropDB()
}
