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
	nuid := "001159263"

	existingAppAssert := CreateSampleUser(t, "test@northeastern.edu", nuid)

	TestRequest{
		Method: "POST",
		Path:   "/api/v1/users/",
		Body: &map[string]interface{}{
			"first_name": "TestX",
			"last_name":  "TestY",
			"email":      "test2@northeastern.edu",
			"password":   "1234567890",
			"nuid":       nuid,
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

func CreateInvalidUser(t *testing.T, body map[string]interface{}, expectedMessage string) ExistingAppAssert {
	return TestRequest{
		Method: "POST",
		Path:   "/api/v1/users/",
		Body:   &body,
	}.TestOnStatusAndMessageKeepDB(t, nil,
		MessageWithStatus{
			Status:  400,
			Message: expectedMessage,
		},
	)
}

func TestCreateUserFailsOnInvalidNUID(t *testing.T) {
	first := "Jermaine"
	last := "Cole"
	goodEmail := "test@northeastern.edu"
	goodPassword := "1234567890"
	goodCollege := "CAMD"
	goodYear := 3

	body := map[string]interface{}{
		"first_name": first,
		"last_name":  last,
		"email":      goodEmail,
		"password":   goodPassword,
		"college":    goodCollege,
		"year":       goodYear,
	}

	// test that it fails on <9 digits
	badNUIDLen := "012"
	body["nuid"] = badNUIDLen
	//TODO change error messages to be more readable
	expectedMessageLen := "Key: 'CreateUserRequestBody.NUID' Error:Field validation for 'NUID' failed on the 'len' tag"
	appAssertLen := CreateInvalidUser(t, body, expectedMessageLen)
	appAssertLen.App.DropDB()

	// test that it fails on non-numbers
	badNUIDNumber := "01234578a"
	body["nuid"] = badNUIDNumber
	expectedMessageNumber := "Key: 'CreateUserRequestBody.NUID' Error:Field validation for 'NUID' failed on the 'number' tag"
	appAssertNumber := CreateInvalidUser(t, body, expectedMessageNumber)
	appAssertNumber.App.DropDB()
}

func TestCreateUserFailsOnInvalidEmail(t *testing.T) {
	first := "Jermaine"
	last := "Cole"
	badEmail := "test@gmail.com"
	goodPassword := "1234567890"
	goodCollege := "CAMD"
	goodYear := 3
	goodNUID := "001159263"

	body := map[string]interface{}{
		"first_name": first,
		"last_name":  last,
		"email":      badEmail,
		"password":   goodPassword,
		"college":    goodCollege,
		"year":       goodYear,
		"nuid":       goodNUID,
	}

	expectedMessage := "Key: 'CreateUserRequestBody.Email' Error:Field validation for 'Email' failed on the 'neu_email' tag"

	appAssert := CreateInvalidUser(t, body, expectedMessage)
	appAssert.App.DropDB()
}

func TestCreateUserFailsOnInvalidPassword(t *testing.T) {
	badPassword := "123"
	first := "Jermaine"
	last := "Cole"
	goodEmail := "test@northeastern.edu"
	goodCollege := "CAMD"
	goodYear := 3
	goodNUID := "001159263"
	expectedMessage := "Key: 'CreateUserRequestBody.Password' Error:Field validation for 'Password' failed on the 'password' tag"

	body := map[string]interface{}{
		"first_name": first,
		"last_name":  last,
		"email":      goodEmail,
		"password":   badPassword,
		"college":    goodCollege,
		"year":       goodYear,
		"nuid":       goodNUID,
	}

	appAssert := CreateInvalidUser(t, body, expectedMessage)
	appAssert.App.DropDB()
}

func TestCreateUserFailsOnInvalidYear(t *testing.T) {
	goodPassword := "123456789"
	first := "Jermaine"
	last := "Cole"
	goodEmail := "test@northeastern.edu"
	goodCollege := "CAMD"
	badYear := 7
	goodNUID := "001159263"
	expectedMessage := "Key: 'CreateUserRequestBody.Year' Error:Field validation for 'Year' failed on the 'max' tag"

	body := map[string]interface{}{
		"first_name": first,
		"last_name":  last,
		"email":      goodEmail,
		"password":   goodPassword,
		"college":    goodCollege,
		"year":       badYear,
		"nuid":       goodNUID,
	}

	appAssert := CreateInvalidUser(t, body, expectedMessage)
	appAssert.App.DropDB()
}

func TestCreateUserFailsOnInvalidCollege(t *testing.T) {
	goodPassword := "123456789"
	first := "Jermaine"
	last := "Cole"
	goodEmail := "test@northeastern.edu"
	badCollege := "oopsies"
	goodYear := 6
	goodNUID := "001159263"
	expectedMessage := "Key: 'CreateUserRequestBody.College' Error:Field validation for 'College' failed on the 'oneof' tag"

	body := map[string]interface{}{
		"first_name": first,
		"last_name":  last,
		"email":      goodEmail,
		"password":   goodPassword,
		"college":    badCollege,
		"year":       goodYear,
		"nuid":       goodNUID,
	}

	appAssert := CreateInvalidUser(t, body, expectedMessage)
	appAssert.App.DropDB()
}

func TestCreateUserFailsOnMissingField(t *testing.T) {

	// tests that:
	//		 if a field is missing, the body parser should fail and return a 400
	// 		 if a field is present but empty, the body parser should fail and return a 400

	password := "123456789"
	first := "Jermaine"
	last := "Cole"
	email := "test@northeastern.edu"
	college := "CS"
	year := 6
	nuid := "001159263"
	expectedMessage := "Key: 'CreateUserRequestBody.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag"

	body := map[string]interface{}{
		"first_name": first,
		"last_name":  last,
		"email":      email,
		"password":   password,
		"college":    college,
		"year":       year,
		"nuid":       nuid,
	}

	delete(body, "first_name")

	appAssert := CreateInvalidUser(t, body, expectedMessage)
	appAssert.App.DropDB()

	//TODO loop through all fields
}

// TODO test extra fields
