package tests

import (
	"fmt"
	"net/http"
	"testing"

	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/goccy/go-json"
	"github.com/huandu/go-assert"
)

func TestGetUsersWorks(t *testing.T) {
	TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/users/",
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
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

				dbUsers, err := transactions.GetUsers(app.Conn, 1, 0)

				assert.NilError(&err)

				assert.Equal(1, len(dbUsers))

				dbUser := dbUsers[0]

				assert.Equal(dbUser, respUser)
			},
		},
	).Close()
}

func TestGetUserWorks(t *testing.T) {
	appAssert, uuid := CreateSampleUser(t, nil)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respUser models.User

				err := json.NewDecoder(resp.Body).Decode(&respUser)

				assert.NilError(err)

				sampleUser := *SampleUserFactory()

				assert.Equal(sampleUser["first_name"].(string), respUser.FirstName)
				assert.Equal(sampleUser["last_name"].(string), respUser.LastName)
				assert.Equal(sampleUser["email"].(string), respUser.Email)
				assert.Equal(sampleUser["nuid"].(string), respUser.NUID)
				assert.Equal(models.College(sampleUser["college"].(string)), respUser.College)
				assert.Equal(models.Year(sampleUser["year"].(int)), respUser.Year)

				dbUser, err := transactions.GetUser(app.Conn, uuid)

				assert.NilError(&err)

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
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestGetUserFailsNotExist(t *testing.T) {
	uuid := uuid.New()

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
	}.TestOnErrorAndDB(t, nil,
		ErrorWithDBTester{
			Error: errors.UserNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var user models.User

				err := app.Conn.Where("id = ?", uuid).First(&user).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestUpdateUserWorks(t *testing.T) {
	appAssert, uuid := CreateSampleUser(t, nil)

	newFirstName := "Michael"
	newLastName := "Brennan"

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
		Body: &map[string]interface{}{
			"first_name": newFirstName,
			"last_name":  newLastName,
		},
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respUser models.User

				err := json.NewDecoder(resp.Body).Decode(&respUser)

				assert.NilError(err)

				assert.Equal(newFirstName, respUser.FirstName)
				assert.Equal(newLastName, respUser.LastName)
				assert.Equal((*SampleUserFactory())["email"].(string), respUser.Email)
				assert.Equal((*SampleUserFactory())["nuid"].(string), respUser.NUID)
				assert.Equal(models.College((*SampleUserFactory())["college"].(string)), respUser.College)
				assert.Equal(models.Year((*SampleUserFactory())["year"].(int)), respUser.Year)

				var dbUser models.User

				err = app.Conn.First(&dbUser, uuid).Error

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
	appAssert, uuid := CreateSampleUser(t, nil)

	for _, invalidData := range []map[string]interface{}{
		{"email": "not.northeastern@gmail.com"},
		{"nuid": "1800-123-4567"},
		{"password": "1234"},
		{"year": 1963},
		{"college": "UT-Austin"},
	} {
		TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
			Body:   &invalidData,
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error:    errors.FailedToValidateUser,
				DBTester: TestNumUsersRemainsAt2,
			},
		)
	}
	appAssert.Close()
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
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
			Body:   SampleUserFactory(),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestUpdateUserFailsOnIdNotExist(t *testing.T) {
	uuid := uuid.New()

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
		Body:   SampleUserFactory(),
	}.TestOnErrorAndDB(t, nil,
		ErrorWithDBTester{
			Error: errors.UserNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var user models.User

				err := app.Conn.Where("id = ?", uuid).First(&user).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeleteUserWorks(t *testing.T) {
	appAssert, uuid := CreateSampleUser(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusNoContent,
			DBTester: TestNumUsersRemainsAt1,
		},
	).Close()
}

func TestDeleteUserNotExist(t *testing.T) {
	uuid := uuid.New()
	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
	}.TestOnErrorAndDB(t, nil,
		ErrorWithDBTester{
			Error: errors.UserNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var user models.User

				err := app.Conn.Where("id = ?", uuid).First(&user).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

				TestNumUsersRemainsAt1(app, assert, resp)
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
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
		}.TestOnErrorAndDB(t, nil,
			ErrorWithDBTester{
				Error:    errors.FailedToValidateID,
				DBTester: TestNumUsersRemainsAt1,
			},
		)
	}
}

func SampleUserFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"first_name": "Jane",
		"last_name":  "Doe",
		"email":      "doe.jane@northeastern.edu",
		"password":   "1234567890&",
		"nuid":       "001234567",
		"college":    "KCCS",
		"year":       3,
	}
}

func AssertUserWithIDBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respUser models.User

	err := json.NewDecoder(resp.Body).Decode(&respUser)

	assert.NilError(err)

	var dbUsers []models.User

	err = app.Conn.Find(&dbUsers).Error

	assert.NilError(err)

	assert.Equal(2, len(dbUsers))

	dbUser := dbUsers[1]

	assert.Equal(dbUser.FirstName, respUser.FirstName)
	assert.Equal(dbUser.LastName, respUser.LastName)
	assert.Equal(dbUser.Email, respUser.Email)
	assert.Equal(dbUser.NUID, respUser.NUID)
	assert.Equal(dbUser.College, respUser.College)
	assert.Equal(dbUser.Year, respUser.Year)

	match, err := auth.ComparePasswordAndHash((*body)["password"].(string), dbUser.PasswordHash)

	assert.NilError(err)

	assert.Assert(match)

	assert.Equal((*body)["first_name"].(string), dbUser.FirstName)
	assert.Equal((*body)["last_name"].(string), dbUser.LastName)
	assert.Equal((*body)["email"].(string), dbUser.Email)
	assert.Equal((*body)["nuid"].(string), dbUser.NUID)
	assert.Equal(models.College((*body)["college"].(string)), dbUser.College)
	assert.Equal(models.Year((*body)["year"].(int)), dbUser.Year)

	return dbUser.ID
}

func AssertSampleUserBodyRespDB(app TestApp, assert *assert.A, resp *http.Response) uuid.UUID {
	return AssertUserWithIDBodyRespDB(app, assert, resp, SampleUserFactory())
}

func CreateSampleUser(t *testing.T, existingAppAssert *ExistingAppAssert) (ExistingAppAssert, uuid.UUID) {
	var uuid uuid.UUID

	newAppAssert := TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/users/",
		Body:   SampleUserFactory(),
	}.TestOnStatusAndDB(t, existingAppAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				uuid = AssertSampleUserBodyRespDB(app, assert, resp)
			},
		},
	)

	if existingAppAssert == nil {
		return newAppAssert, uuid
	} else {
		return *existingAppAssert, uuid
	}
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
	appAssert, _ := CreateSampleUser(t, nil)
	appAssert.Close()
}

func TestCreateUserFailsIfUserWithEmailAlreadyExists(t *testing.T) {
	appAssert, _ := CreateSampleUser(t, nil)

	TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/users/",
		Body:   SampleUserFactory(),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error:    errors.UserAlreadyExists,
			DBTester: TestNumUsersRemainsAt2,
		},
	)

	appAssert.Close()
}

func TestCreateUserFailsIfUserWithNUIDAlreadyExists(t *testing.T) {
	appAssert, _ := CreateSampleUser(t, nil)

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
		Method: fiber.MethodPost,
		Path:   "/api/v1/users/",
		Body:   slightlyDifferentSampleUser,
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error:    errors.UserAlreadyExists,
			DBTester: TestNumUsersRemainsAt2,
		},
	).Close()
}

func AssertCreateBadDataFails(t *testing.T, jsonKey string, badValues []interface{}) {
	appAssert, _ := CreateSampleUser(t, nil)

	for _, badValue := range badValues {
		sampleUserPermutation := *SampleUserFactory()
		sampleUserPermutation[jsonKey] = badValue

		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/users/",
			Body:   &sampleUserPermutation,
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error:    errors.FailedToValidateUser,
				DBTester: TestNumUsersRemainsAt2,
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
		})
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
	appAssert, _ := CreateSampleUser(t, nil)

	for _, missingField := range []string{
		"first_name",
		"last_name",
		"email",
		"password",
		"nuid",
		"college",
		"year",
	} {
		sampleUserPermutation := *SampleUserFactory()
		delete(sampleUserPermutation, missingField)

		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/users/",
			Body:   &sampleUserPermutation,
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error:    errors.FailedToValidateUser,
				DBTester: TestNumUsersRemainsAt2,
			},
		)
	}
	appAssert.Close()
}

// test create user following works
func TestCreateUserFollowingWorks(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var dbUser models.User
				err := app.Conn.Preload("Follower").First(&dbUser, userUUID).Error
				assert.NilError(err)

				assert.Equal(len(dbUser.Follower), 1)
			},
		},
	)
	appAssert.Close()
}

func TestCreateUserFollowingFailsClubIdBadRequest(t *testing.T) {
	appAssert, userUUID := CreateSampleUser(t, nil)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, badRequest),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID).Close()
	}
}

func TestCreateUserFollowingFailsUserIdBadRequest(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(t, nil)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", badRequest, clubUUID),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID).Close()
	}
}

func TestCreateUserFollowingFailsUserNotExist(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(t, nil)
	userUUIDNotExist := uuid.New()

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUIDNotExist, clubUUID),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.UserNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var user models.User
				err := app.Conn.Where("id = ?", userUUIDNotExist).First(&user).Error
				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestCreateUserFollowingFailsClubNotExist(t *testing.T) {
	appAssert, userUUID := CreateSampleUser(t, nil)
	clubUUIDNotExist := uuid.New()

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/users/%s/follower/%s", userUUID, clubUUIDNotExist),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.ClubNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club
				err := app.Conn.Where("id = ?", clubUUIDNotExist).First(&club).Error
				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}
