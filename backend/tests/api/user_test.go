package tests

import (
	"bytes"
	"fmt"
	"math/rand"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/goccy/go-json"
	"github.com/huandu/go-assert"
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
		Path:   "/api/v1/users/1000",
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
	perms := 0

	for _, permutation := range AllCasingPermutations(email) {
		if perms == 20 {
			break
		}
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
		perms++
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
	// To use for testing invalid users that should fail to be created
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
	// tests that:
	// 		 if nuid is not 9 digits, the Post Request should fail and return a 400
	// 		 if nuid is not a number, the Post Request should fail and return a 400

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
	// tests that:
	// 		 if email is not a northeastern email (ends in @northeastern.edu), the Post Request should fail and return a 400
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
	// tests that:
	// 		 if password is not at least 10 characters, the Post Request should fail and return a 400
	// 		 TODO create better password requirements
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
	// tests that:
	// 		 if year is not within range [1,6], the Post Request should fail and return a 400
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
	// tests that:
	// 		 if college is not one of CAMD DMSB KCCS CE BCHS SL CPS CS CSSH, the Post Request should fail and return a 400

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
	//		 if a field is missing, the Post Request should fail and return a 400
	// 		 if a field is present but empty, the Post Request should fail and return a 400

	password := "123456789"
	first := "Jermaine"
	last := "Cole"
	email := "test@northeastern.edu"
	college := "CS"
	year := 6
	nuid := "001159263"

	body := map[string]interface{}{
		"first_name": first,
		"last_name":  last,
		"email":      email,
		"password":   password,
		"college":    college,
		"year":       year,
		"nuid":       nuid,
	}

	// map from json field name to struct field name
	fields := map[string]string{"first_name": "FirstName", "last_name": "LastName", "email": "Email", "password": "Password", "college": "College", "year": "Year", "nuid": "NUID"}

	//loops through each field and removes it from the body then tests that the post fails and returns a 400
	for structKey, jsonKey := range fields {
		// Create a copy of the body without the current field
		updatedBody := make(map[string]interface{})
		for key, value := range body {
			if key != structKey {
				updatedBody[key] = value
			}
		}

		expectedMesage := "Key: 'CreateUserRequestBody." + jsonKey + "' Error:Field validation for '" + jsonKey + "' failed on the 'required' tag"

		appAssert := CreateInvalidUser(t, updatedBody, expectedMesage)
		appAssert.App.DropDB()
	}

}

/*
Dear TLs David, Garrett,

I can not for the life of me figure out how to do this test. It has become the bane of me
I also have not touched nor smelled the notion of "grass" in the last 27 hours
Please forgive me. I am but a humble student
- Olivier
*/
func TestCreateUserFailsOnExtraFields(t *testing.T) {
	// tests that:
	// 		 if extra fields are present, the Post Request should fail and return a 400

	password := "123456789"
	first := "Jermaine"
	last := "Cole"
	email := "jermaine@northeastern.edu"
	college := "CS"
	year := 6
	nuid := "001159263"
	someField := "someField"

	body := map[string]interface{}{
		"first_name": first,
		"last_name":  last,
		"email":      email,
		"password":   password,
		"college":    college,
		"year":       year,
		"nuid":       nuid,
		// foreign fields should not be allowed
		"extra": someField,
	}

	appAssert := CreateInvalidUser(t, body, "expectedMessage")
	appAssert.App.DropDB()

}
