package tests

import (
	"fmt"

	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/auth"
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
				var users []models.User

				err := json.NewDecoder(resp.Body).Decode(&users)

				assert.NilError(err)

				assert.Equal(1, len(users))

				respUser := users[0]

				assert.Equal("SAC", respUser.FirstName)
				assert.Equal("Super", respUser.LastName)
				assert.Equal("generatesac@gmail.com", respUser.Email)
				assert.Equal("000000000", respUser.NUID)
				assert.Equal(models.College("KCCS"), respUser.College)
				assert.Equal(models.Year(1), respUser.Year)

				dbUsers, err := transactions.GetAllUsers(app.Conn)

				assert.NilError(err)

				assert.Equal(1, len(dbUsers))

				dbUser := dbUsers[0]

				assert.Equal(dbUser, respUser)
			},
		},
	).Close()
}

func TestGetUserWorks(t *testing.T) {
	id := 1

	TestRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/api/v1/users/%d", id),
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respUser models.User

				err := json.NewDecoder(resp.Body).Decode(&respUser)

				assert.NilError(err)

				assert.Equal("SAC", respUser.FirstName)
				assert.Equal("Super", respUser.LastName)
				assert.Equal("generatesac@gmail.com", respUser.Email)
				assert.Equal("000000000", respUser.NUID)
				assert.Equal(models.College("KCCS"), respUser.College)
				assert.Equal(models.Year(1), respUser.Year)

				dbUser, err := transactions.GetUser(app.Conn, uint(id))

				assert.NilError(err)

				assert.Equal(dbUser, &respUser)
			},
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
			Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400, // should be 400 but the hardcoded error code is 500 @David
				Message: "Bad Request",
			},
		).Close()
	}
}

func TestGetUserFailsNotExist(t *testing.T) {
	id := uint(69)

	TestRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/api/v1/users/%d", id),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404, // should be 404 but the hardcoded error code is 500 @David
				Message: "Not Found",
			},
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var user models.User

				err := app.Conn.Where("id = ?", id).First(&user).Error

				assert.NilError(err)

				assert.Equal(nil, user)
			},
		},
	).Close()
}

func TestUpdateUserWorks(t *testing.T) {
	appAssert := CreateSampleUser(t)

	id := 2
	newFirstName := "Michael"
	newLastName := "Brennan"

	TestRequest{
		Method: "PATCH",
		Path:   fmt.Sprintf("/api/v1/users/%d", id),
		Body: &map[string]interface{}{
			"first_name": newFirstName,
			"last_name":  newLastName,
		},
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respUser models.User

				err := json.NewDecoder(resp.Body).Decode(&respUser)

				assert.NilError(err)

				assert.Equal(newFirstName, respUser.FirstName)
				assert.Equal(newLastName, respUser.LastName)
				assert.Equal((*SampleUserBody)["email"].(string), respUser.Email)
				assert.Equal((*SampleUserBody)["nuid"].(string), respUser.NUID)
				assert.Equal(models.College((*SampleUserBody)["college"].(string)), respUser.College)
				assert.Equal(models.Year((*SampleUserBody)["year"].(int)), respUser.Year)

				var dbUser models.User

				err = app.Conn.First(&dbUser, id).Error

				assert.NilError(err)

				assert.Equal(dbUser.FirstName, respUser.FirstName)
				assert.Equal(dbUser.LastName, respUser.LastName)
				assert.Equal(dbUser.Email, respUser.Email)
				assert.Equal(dbUser.NUID, respUser.NUID)
				assert.Equal(dbUser.College, respUser.College)
				assert.Equal(dbUser.Year, respUser.Year)
			},
		},
	).Close()
}

func TestUpdateUserFailsOnInvalidBody(t *testing.T) {
	appAssert := CreateSampleUser(t)

	for _, invalidData := range []map[string]interface{}{
		{"email": "not.northeastern@gmail.com"},
		{"nuid": "1800-123-4567"},
		{"password": "1234"},
		{"year": 1963},
		{"college": "UT-Austin"},
	} {
		TestRequest{
			Method: "PATCH",
			Path:   "/api/v1/users/2",
			Body:   &invalidData,
		}.TestOnStatusMessageAndDB(t, &appAssert,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400, // should be 400 but the hardcoded error code is 500 @David
					Message: "Bad Request",
				},
				DBTester: TestNumUsersRemainsAt2,
			},
		).Close()
	}
}

func TestUpdateUserFailsBadRequest(t *testing.T) {
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		TestRequest{
			Method: "PATCH",
			Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400, // should be 400 but the hardcoded error code is 500 @David
				Message: "Bad Request",
			},
		).Close()
	}
}

func TestUpdateUserFailsOnIdNotExist(t *testing.T) {
	id := uint(69)

	TestRequest{
		Method: "PATCH",
		Path:   fmt.Sprintf("/api/v1/users/%d", id),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404, // should be 404 but the hardcoded error code is 500 @David
				Message: "Not Found",
			},
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var user models.User

				err := app.Conn.Where("id = ?", id).First(&user).Error

				assert.NilError(err)

				assert.Equal(nil, user)
			},
		},
	).Close()
}

func TestDeleteUserWorks(t *testing.T) {
	appAssert := CreateSampleUser(t)

	TestRequest{
		Method: "DELETE",
		Path:   "/api/v1/users/2",
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   204,
			DBTester: TestNumUsersRemainsAt1,
		},
	).Close()
}

func TestDeleteUserNotExist(t *testing.T) {
	id := uint(69)

	TestRequest{
		Method: "DELTE",
		Path:   fmt.Sprintf("/api/v1/users/%d", id),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404, // should be 404 but the hardcoded error code is 500 @David
				Message: "Not Found",
			},
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var user models.User

				err := app.Conn.Where("id = ?", id).First(&user).Error

				assert.NilError(err)

				assert.Equal(nil, user)
			},
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

var SampleUserBody = &map[string]interface{}{
	"first_name": "Jane",
	"last_name":  "Doe",
	"email":      "doe.jane@northeastern.edu",
	"password":   "1234567890&",
	"nuid":       "001234567",
	"college":    "KCCS",
	"year":       3,
}

func AssertUserWithIDCreatedBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, id uint) {
	var respUser models.User

	err := json.NewDecoder(resp.Body).Decode(&respUser)

	assert.NilError(err)

	var dbUser models.User

	err = app.Conn.First(&dbUser, id).Error

	assert.NilError(err)

	assert.Equal(dbUser.FirstName, respUser.FirstName)
	assert.Equal(dbUser.LastName, respUser.LastName)
	assert.Equal(dbUser.Email, respUser.Email)
	assert.Equal(dbUser.NUID, respUser.NUID)
	assert.Equal(dbUser.College, respUser.College)
	assert.Equal(dbUser.Year, respUser.Year)

	match, err := auth.ComparePasswordAndHash((*SampleUserBody)["password"].(string), dbUser.PasswordHash)

	assert.NilError(err)

	assert.Assert(match)

	assert.Equal((*SampleUserBody)["first_name"].(string), dbUser.FirstName)
	assert.Equal((*SampleUserBody)["last_name"].(string), dbUser.LastName)
	assert.Equal((*SampleUserBody)["email"].(string), dbUser.Email)
	assert.Equal((*SampleUserBody)["nuid"].(string), dbUser.NUID)
	assert.Equal(models.College((*SampleUserBody)["college"].(string)), dbUser.College)
	assert.Equal(models.Year((*SampleUserBody)["year"].(int)), dbUser.Year)
}

func AssertSampleUserCreatedBodyRespDB(app TestApp, assert *assert.A, resp *http.Response) {
	AssertUserWithIDCreatedBodyRespDB(app, assert, resp, 2)
}

func CreateSampleUser(t *testing.T) ExistingAppAssert {
	return TestRequest{
		Method: "POST",
		Path:   "/api/v1/users/",
		Body:   SampleUserBody,
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status:   201,
			DBTester: AssertSampleUserCreatedBodyRespDB,
		},
	)
}

func AssertNumUsersRemainsAtN(app TestApp, assert *assert.A, resp *http.Response, n int) {
	var users []models.User

	err := app.Conn.Find(&users).Error

	assert.NilError(err)

	assert.Equal(n, len(users))
}

var TestNumUsersRemainsAt1 = func(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumUsersRemainsAtN(app, assert, resp, 1)
}

var TestNumUsersRemainsAt2 = func(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumUsersRemainsAtN(app, assert, resp, 2)
}

func TestCreateUserWorks(t *testing.T) {
	CreateSampleUser(t).Close()
}

func TestCreateUserFailsIfCategoryWithEmailAlreadyExists(t *testing.T) {
	appAssert := CreateSampleUser(t)

	for _, email := range AllCasingPermutations((*SampleUserBody)["email"].(string)) {
		sampleUserPermutation := *SampleUserBody
		sampleUserPermutation["email"] = email

		TestRequest{
			Method: "POST",
			Path:   "/api/v1/users/",
			Body:   &sampleUserPermutation,
		}.TestOnStatusMessageAndDB(t, &appAssert,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400,
					Message: "user with that email already exists",
				},
				DBTester: TestNumUsersRemainsAt1,
			},
		)

	}
	appAssert.Close()

}

func TestCreateUserFailsIfUserWithNUIDAlreadyExists(t *testing.T) {
	appAssert := CreateSampleUser(t)

	slightlyDifferentSampleUser := &map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"email":      "doe.john@northeastern.edu",
		"password":   "1234567890&",
		"nuid":       "001234567",
		"college":    "KCCS",
		"year":       3,
	}

	TestRequest{
		Method: "POST",
		Path:   "/api/v1/users/",
		Body:   slightlyDifferentSampleUser,
	}.TestOnStatusMessageAndDB(t, &appAssert,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  400,                     // should be a 400 @David
				Message: "Internal Server Error", // not the biggest fan of this error message @David
			},
			DBTester: TestNumUsersRemainsAt1,
		},
	).Close()
}

func TestCreateUserFailsIfUserWithEmailAlreadyExists(t *testing.T) {
	appAssert := CreateSampleUser(t)

	slightlyDifferentSampleUser := &map[string]interface{}{
		"first_name": "Jane",
		"last_name":  "Doe",
		"email":      "doe.jane@northeastern.edu",
		"password":   "1234567890&",
		"nuid":       "001234568",
		"college":    "KCCS",
		"year":       3,
	}

	TestRequest{
		Method: "POST",
		Path:   "/api/v1/users/",
		Body:   slightlyDifferentSampleUser,
	}.TestOnStatusMessageAndDB(t, &appAssert,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  400,           // should be a 400 @David
				Message: "Bad Request", // not the biggest fan of this error message @David
			},
			DBTester: TestNumUsersRemainsAt1,
		},
	).Close()
}

func AssertCreateBadDataFails(t *testing.T, jsonKey string, badValues []interface{}) {
	appAssert := CreateSampleUser(t)

	for _, badValue := range badValues {
		sampleUserPermutation := *SampleUserBody
		sampleUserPermutation[jsonKey] = badValue

		TestRequest{
			Method: "POST",
			Path:   "/api/v1/users/",
			Body:   &sampleUserPermutation,
		}.TestOnStatusMessageAndDB(t, &appAssert,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400, // should be a 400 @David
					Message: "Bad Request",
				},
				DBTester: TestNumUsersRemainsAt1,
			},
		)
	}
	appAssert.Close()
}

func TestCreateUserFailsOnInvalidNUID(t *testing.T) {
	AssertCreateBadDataFails(t,
		"nuid",
		[]interface{}{
			"00123456",
			"0012345678",
			"00123456a",
			"00123456!",
			""})
}

func TestCreateUserFailsOnInvalidEmail(t *testing.T) {
	AssertCreateBadDataFails(t,
		"email",
		[]interface{}{
			"doe.jane@northeastern",
			"doe.jane",
			"doe.jane@",
			"doe.jane@northeastern.",
			"doe.jane@northeastern.e",
			"",
		})
}

func TestCreateUserFailsOnInvalidPassword(t *testing.T) {
	AssertCreateBadDataFails(t,
		"password",
		[]interface{}{
			"",
			"foo",
			"abcdefg",
			"abcdefg0",
			"abcdefg@",
		})
}

func TestCreateUserFailsOnInvalidYear(t *testing.T) {
	AssertCreateBadDataFails(t,
		"year",
		[]interface{}{
			-1,
			0,
			7,
		})
}

func TestCreateUserFailsOnInvalidCollege(t *testing.T) {
	khouryAbbreviation := "KCCS"
	permutations := AllCasingPermutations(khouryAbbreviation)
	permutationsWithoutKhoury := make([]interface{}, len(permutations)-1)
	for _, permutation := range permutations {
		if permutation != khouryAbbreviation {
			permutationsWithoutKhoury = append(permutationsWithoutKhoury, permutation)
		}

	}

	AssertCreateBadDataFails(t,
		"college",
		permutationsWithoutKhoury)
}

func TestCreateUserFailsOnMissingFields(t *testing.T) {
	appAssert := CreateSampleUser(t)

	for _, missingField := range []string{
		"first_name",
		"last_name",
		"email",
		"password",
		"nuid",
		"college",
		"year",
	} {
		sampleUserPermutation := *SampleUserBody
		delete(sampleUserPermutation, missingField)

		TestRequest{
			Method: "POST",
			Path:   "/api/v1/users/",
			Body:   &sampleUserPermutation,
		}.TestOnStatusMessageAndDB(t, &appAssert,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400, // should be a 400 @David
					Message: "Bad Request",
				},
				DBTester: TestNumUsersRemainsAt1,
			},
		)
	}
	appAssert.Close()
}
