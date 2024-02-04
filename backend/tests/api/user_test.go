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
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/goccy/go-json"
	"github.com/huandu/go-assert"
)

func TestGetUsersWorksForSuper(t *testing.T) {
	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/users/",
		Role:   &models.Super,
	}.TestOnStatusAndDB(t, nil,
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
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

func TestGetUsersFailsForStudent(t *testing.T) {
	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/users/",
		Role:   &models.Student,
	}.TestOnError(t, nil, errors.Unauthorized).Close()
}

func TestGetUserWorks(t *testing.T) {
	appAssert, uuid, _ := CreateSampleStudent(t, nil)

	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
	}.TestOnStatusAndDB(t, &appAssert,
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var respUser models.User

				err := json.NewDecoder(resp.Body).Decode(&respUser)

				assert.NilError(err)

				sampleStudent, rawPassword := h.SampleStudentFactory()

				sampleUser := *h.SampleStudentJSONFactory(sampleStudent, rawPassword)

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
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestGetUserFailsNotExist(t *testing.T) {
	uuid := uuid.New()

	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
	}.TestOnErrorAndDB(t, nil,
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var user models.User

				err := app.Conn.Where("id = ?", uuid).First(&user).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestUpdateUserWorks(t *testing.T) {
	appAssert, uuid, _ := CreateSampleStudent(t, nil)

	newFirstName := "Michael"
	newLastName := "Brennan"

	h.TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
		Body: &map[string]interface{}{
			"first_name": newFirstName,
			"last_name":  newLastName,
		},
	}.TestOnStatusAndDB(t, &appAssert,
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var respUser models.User

				err := json.NewDecoder(resp.Body).Decode(&respUser)

				assert.NilError(err)

				sampleStudent, rawPassword := h.SampleStudentFactory()

				sampleStudentJSON := *h.SampleStudentJSONFactory(sampleStudent, rawPassword)

				assert.Equal(newFirstName, respUser.FirstName)
				assert.Equal(newLastName, respUser.LastName)
				assert.Equal((sampleStudentJSON)["email"].(string), respUser.Email)
				assert.Equal((sampleStudentJSON)["nuid"].(string), respUser.NUID)
				assert.Equal(models.College((sampleStudentJSON)["college"].(string)), respUser.College)
				assert.Equal(models.Year((sampleStudentJSON)["year"].(int)), respUser.Year)

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
	appAssert, uuid, _ := CreateSampleStudent(t, nil)

	for _, invalidData := range []map[string]interface{}{
		{"email": "not.northeastern@gmail.com"},
		{"nuid": "1800-123-4567"},
		{"password": "1234"},
		{"year": 1963},
		{"college": "UT-Austin"},
	} {
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
			Body:   &invalidData,
		}.TestOnErrorAndDB(t, &appAssert,
			h.ErrorWithTester{
				Error:  errors.FailedToValidateUser,
				Tester: TestNumUsersRemainsAt2,
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

	sampleStudent, rawPassword := h.SampleStudentFactory()

	for _, badRequest := range badRequests {
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
			Body:   h.SampleStudentJSONFactory(sampleStudent, rawPassword),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestUpdateUserFailsOnIdNotExist(t *testing.T) {
	uuid := uuid.New()

	sampleStudent, rawPassword := h.SampleStudentFactory()

	h.TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
		Body:   h.SampleStudentJSONFactory(sampleStudent, rawPassword),
	}.TestOnErrorAndDB(t, nil,
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var user models.User

				err := app.Conn.Where("id = ?", uuid).First(&user).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeleteUserWorks(t *testing.T) {
	appAssert, uuid, _ := CreateSampleStudent(t, nil)

	h.TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
	}.TestOnStatusAndDB(t, &appAssert,
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: TestNumUsersRemainsAt1,
		},
	).Close()
}

func TestDeleteUserNotExist(t *testing.T) {
	uuid := uuid.New()
	h.TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
	}.TestOnErrorAndDB(t, nil,
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
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
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
		}.TestOnErrorAndDB(t, nil,
			h.ErrorWithTester{
				Error:  errors.FailedToValidateID,
				Tester: TestNumUsersRemainsAt1,
			},
		)
	}
}

func AssertUserWithIDBodyRespDB(app h.TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
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

func AssertSampleUserBodyRespDB(app h.TestApp, assert *assert.A, resp *http.Response) uuid.UUID {
	sampleStudent, rawPassword := h.SampleStudentFactory()

	return AssertUserWithIDBodyRespDB(app, assert, resp, h.SampleStudentJSONFactory(sampleStudent, rawPassword))
}

func CreateSampleStudent(t *testing.T, existingAppAssert *h.ExistingAppAssert) (h.ExistingAppAssert, uuid.UUID, *map[string]interface{}) {
	var uuid uuid.UUID

	sampleStudent, rawPassword := h.SampleStudentFactory()

	newAppAssert := h.TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/users/",
		Body:   h.SampleStudentJSONFactory(sampleStudent, rawPassword),
	}.TestOnStatusAndDB(t, existingAppAssert,
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				uuid = AssertSampleUserBodyRespDB(app, assert, resp)
			},
		},
	)

	if existingAppAssert == nil {
		return newAppAssert, uuid, h.SampleStudentJSONFactory(sampleStudent, rawPassword)
	} else {
		return *existingAppAssert, uuid, h.SampleStudentJSONFactory(sampleStudent, rawPassword)
	}
}

func AssertNumUsersRemainsAtN(app h.TestApp, assert *assert.A, resp *http.Response, n int) {
	var users []models.User

	err := app.Conn.Find(&users).Error

	assert.NilError(err)

	assert.Equal(n, len(users))
}

var TestNumUsersRemainsAt1 = func(app h.TestApp, assert *assert.A, resp *http.Response) {
	AssertNumUsersRemainsAtN(app, assert, resp, 1)
}

var TestNumUsersRemainsAt2 = func(app h.TestApp, assert *assert.A, resp *http.Response) {
	AssertNumUsersRemainsAtN(app, assert, resp, 2)
}

func TestCreateUserWorks(t *testing.T) {
	appAssert, _, _ := CreateSampleStudent(t, nil)
	appAssert.Close()
}

func TestCreateUserFailsIfUserWithEmailAlreadyExists(t *testing.T) {
	appAssert, studentUUID, body := CreateSampleStudent(t, nil)

	(*body)["id"] = studentUUID

	h.TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/users/",
		Body:   body,
	}.TestOnErrorAndDB(t, &appAssert,
		h.ErrorWithTester{
			Error:  errors.UserAlreadyExists,
			Tester: TestNumUsersRemainsAt2,
		},
	)

	appAssert.Close()
}

func TestCreateUserFailsIfUserWithNUIDAlreadyExists(t *testing.T) {
	appAssert, _, _ := CreateSampleStudent(t, nil)

	sampleStudent, rawPassword := h.SampleStudentFactory()

	slightlyDifferentSampleStudentJSON := h.SampleStudentJSONFactory(sampleStudent, rawPassword)

	(*slightlyDifferentSampleStudentJSON)["first_name"] = "John"
	(*slightlyDifferentSampleStudentJSON)["last_name"] = "Doe"
	(*slightlyDifferentSampleStudentJSON)["email"] = "doe.john@northeastern.edu"

	h.TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/users/",
		Body:   slightlyDifferentSampleStudentJSON,
	}.TestOnErrorAndDB(t, &appAssert,
		h.ErrorWithTester{
			Error:  errors.UserAlreadyExists,
			Tester: TestNumUsersRemainsAt2,
		},
	).Close()
}

func AssertCreateBadDataFails(t *testing.T, jsonKey string, badValues []interface{}) {
	appAssert, _, _ := CreateSampleStudent(t, nil)

	sampleStudent, rawPassword := h.SampleStudentFactory()

	for _, badValue := range badValues {
		sampleUserPermutation := *h.SampleStudentJSONFactory(sampleStudent, rawPassword)
		sampleUserPermutation[jsonKey] = badValue

		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/users/",
			Body:   &sampleUserPermutation,
		}.TestOnErrorAndDB(t, &appAssert,
			h.ErrorWithTester{
				Error:  errors.FailedToValidateUser,
				Tester: TestNumUsersRemainsAt2,
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
	permutations := h.AllCasingPermutations(khouryAbbreviation)
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
	appAssert, _, _ := CreateSampleStudent(t, nil)

	sampleStudent, rawPassword := h.SampleStudentFactory()

	for _, missingField := range []string{
		"first_name",
		"last_name",
		"email",
		"password",
		"nuid",
		"college",
		"year",
	} {
		sampleUserPermutation := *h.SampleStudentJSONFactory(sampleStudent, rawPassword)
		delete(sampleUserPermutation, missingField)

		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/users/",
			Body:   &sampleUserPermutation,
		}.TestOnErrorAndDB(t, &appAssert,
			h.ErrorWithTester{
				Error:  errors.FailedToValidateUser,
				Tester: TestNumUsersRemainsAt2,
			},
		)
	}
	appAssert.Close()
}
