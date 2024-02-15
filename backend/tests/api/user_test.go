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
)

func TestGetUsersWorksForSuper(t *testing.T) {
	t.Parallel()
	h.InitTest(t).TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   "/api/v1/users/",
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var users []models.User

				err := json.NewDecoder(resp.Body).Decode(&users)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(users))

				respUser := users[0]

				eaa.Assert.Equal("SAC", respUser.FirstName)
				eaa.Assert.Equal("Super", respUser.LastName)
				eaa.Assert.Equal("generatesac@gmail.com", respUser.Email)
				eaa.Assert.Equal("000000000", respUser.NUID)
				eaa.Assert.Equal(models.College("KCCS"), respUser.College)
				eaa.Assert.Equal(models.Year(1), respUser.Year)

				dbUsers, err := transactions.GetUsers(eaa.App.Conn, 1, 0)

				eaa.Assert.NilError(&err)

				eaa.Assert.Equal(1, len(dbUsers))

				dbUser := dbUsers[0]

				eaa.Assert.Equal(dbUser, respUser)
			},
		},
	).Close()
}

func TestGetUsersFailsForStudent(t *testing.T) {
	t.Parallel()
	h.InitTest(t).TestOnError(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   "/api/v1/users/",
			Role:   &models.Student,
		},
		errors.Unauthorized,
	).Close()
}

func TestGetUserWorks(t *testing.T) {
	t.Parallel()
	h.InitTest(t).TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodGet,
			Path:               "/api/v1/users/:userID",
			Role:               &models.Student,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var respUser models.User

				err := json.NewDecoder(resp.Body).Decode(&respUser)

				eaa.Assert.NilError(err)

				sampleStudent, rawPassword := h.SampleStudentFactory()

				sampleUser := *h.SampleStudentJSONFactory(sampleStudent, rawPassword)

				eaa.Assert.Equal(sampleUser["first_name"].(string), respUser.FirstName)
				eaa.Assert.Equal(sampleUser["last_name"].(string), respUser.LastName)
				eaa.Assert.Equal(sampleUser["email"].(string), respUser.Email)
				eaa.Assert.Equal(sampleUser["nuid"].(string), respUser.NUID)
				eaa.Assert.Equal(models.College(sampleUser["college"].(string)), respUser.College)
				eaa.Assert.Equal(models.Year(sampleUser["year"].(int)), respUser.Year)

				dbUser, err := transactions.GetUser(eaa.App.Conn, eaa.App.TestUser.UUID)

				eaa.Assert.NilError(&err)

				eaa.Assert.Equal(dbUser, &respUser)
			},
		},
	).Close()
}

func TestGetUserFailsBadRequest(t *testing.T) {
	t.Parallel()
	appAssert := h.InitTest(t)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodGet,
				Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestGetUserFailsNotExist(t *testing.T) {
	t.Parallel()
	uuid := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", uuid).First(&user).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestUpdateUserWorks(t *testing.T) {
	t.Parallel()
	newFirstName := "Michael"
	newLastName := "Brennan"

	h.InitTest(t).TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   "/api/v1/users/:userID",
			Body: &map[string]interface{}{
				"first_name": newFirstName,
				"last_name":  newLastName,
			},
			Role:               &models.Student,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var respUser models.User

				err := json.NewDecoder(resp.Body).Decode(&respUser)

				eaa.Assert.NilError(err)

				sampleStudent, rawPassword := h.SampleStudentFactory()

				sampleStudentJSON := *h.SampleStudentJSONFactory(sampleStudent, rawPassword)

				eaa.Assert.Equal(newFirstName, respUser.FirstName)
				eaa.Assert.Equal(newLastName, respUser.LastName)
				eaa.Assert.Equal((sampleStudentJSON)["email"].(string), respUser.Email)
				eaa.Assert.Equal((sampleStudentJSON)["nuid"].(string), respUser.NUID)
				eaa.Assert.Equal(models.College((sampleStudentJSON)["college"].(string)), respUser.College)
				eaa.Assert.Equal(models.Year((sampleStudentJSON)["year"].(int)), respUser.Year)

				var dbUser models.User

				err = eaa.App.Conn.First(&dbUser, eaa.App.TestUser.UUID).Error

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(dbUser.FirstName, respUser.FirstName)
				eaa.Assert.Equal(dbUser.LastName, respUser.LastName)
				eaa.Assert.Equal(dbUser.Email, respUser.Email)
				eaa.Assert.Equal(dbUser.NUID, respUser.NUID)
				eaa.Assert.Equal(dbUser.College, respUser.College)
				eaa.Assert.Equal(dbUser.Year, respUser.Year)
			},
		},
	).Close()
}

func TestUpdateUserFailsOnInvalidBody(t *testing.T) {
	t.Parallel()
	for _, invalidData := range []map[string]interface{}{
		{"email": "not.northeastern@gmail.com"},
		{"nuid": "1800-123-4567"},
		{"year": 1963},
		{"college": "UT-Austin"},
	} {
		invalidData := invalidData
		h.InitTest(t).TestOnErrorAndTester(
			h.TestRequest{
				Method:             fiber.MethodPatch,
				Path:               "/api/v1/users/:userID",
				Body:               &invalidData,
				Role:               &models.Student,
				TestUserIDReplaces: h.StringToPointer(":userID"),
			},
			h.ErrorWithTester{
				Error:  errors.FailedToValidateUser,
				Tester: TestNumUsersRemainsAt2,
			},
		).Close()
	}
}

func TestUpdateUserFailsBadRequest(t *testing.T) {
	t.Parallel()
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	sampleStudent, rawPassword := h.SampleStudentFactory()
	slightlyDifferentSampleStudentJSON := h.SampleStudentJSONFactory(sampleStudent, rawPassword)
	(*slightlyDifferentSampleStudentJSON)["first_name"] = "John"

	for _, badRequest := range badRequests {
		h.InitTest(t).TestOnError(h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
			Body:   slightlyDifferentSampleStudentJSON,
			Role:   &models.Student,
		},
			errors.FailedToValidateID,
		).Close()
	}
}

func TestUpdateUserFailsOnIdNotExist(t *testing.T) {
	t.Parallel()
	uuid := uuid.New()

	sampleStudent, rawPassword := h.SampleStudentFactory()

	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
			Body:   h.SampleStudentJSONFactory(sampleStudent, rawPassword),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", uuid).First(&user).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeleteUserWorks(t *testing.T) {
	t.Parallel()
	h.InitTest(t).TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodDelete,
			Path:               "/api/v1/users/:userID",
			Role:               &models.Student,
			TestUserIDReplaces: h.StringToPointer(":userID"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: TestNumUsersRemainsAt1,
		},
	).Close()
}

func TestDeleteUserNotExist(t *testing.T) {
	t.Parallel()
	uuid := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(h.TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/users/%s", uuid),
		Role:   &models.Super,
	},
		h.ErrorWithTester{
			Error: errors.UserNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var user models.User

				err := eaa.App.Conn.Where("id = ?", uuid).First(&user).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

				TestNumUsersRemainsAt1(eaa, resp)
			},
		},
	).Close()
}

func TestDeleteUserBadRequest(t *testing.T) {
	t.Parallel()
	appAssert := h.InitTest(t)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"hello",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodDelete,
				Path:   fmt.Sprintf("/api/v1/users/%s", badRequest),
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToValidateID,
				Tester: TestNumUsersRemainsAt1,
			},
		)
	}

	appAssert.Close()
}

func AssertUserWithIDBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respUser models.User

	err := json.NewDecoder(resp.Body).Decode(&respUser)

	eaa.Assert.NilError(err)

	var dbUsers []models.User

	err = eaa.App.Conn.Find(&dbUsers).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(2, len(dbUsers))

	dbUser := dbUsers[1]

	eaa.Assert.Equal(dbUser.FirstName, respUser.FirstName)
	eaa.Assert.Equal(dbUser.LastName, respUser.LastName)
	eaa.Assert.Equal(dbUser.Email, respUser.Email)
	eaa.Assert.Equal(dbUser.NUID, respUser.NUID)
	eaa.Assert.Equal(dbUser.College, respUser.College)
	eaa.Assert.Equal(dbUser.Year, respUser.Year)

	match, err := auth.ComparePasswordAndHash((*body)["password"].(string), dbUser.PasswordHash)

	eaa.Assert.NilError(err)

	eaa.Assert.Assert(match)

	eaa.Assert.Equal((*body)["first_name"].(string), dbUser.FirstName)
	eaa.Assert.Equal((*body)["last_name"].(string), dbUser.LastName)
	eaa.Assert.Equal((*body)["email"].(string), dbUser.Email)
	eaa.Assert.Equal((*body)["nuid"].(string), dbUser.NUID)
	eaa.Assert.Equal(models.College((*body)["college"].(string)), dbUser.College)
	eaa.Assert.Equal(models.Year((*body)["year"].(int)), dbUser.Year)

	return dbUser.ID
}

func AssertSampleUserBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response) uuid.UUID {
	sampleStudent, rawPassword := h.SampleStudentFactory()

	return AssertUserWithIDBodyRespDB(eaa, resp, h.SampleStudentJSONFactory(sampleStudent, rawPassword))
}

func CreateSampleStudent(t *testing.T, existingAppAssert *h.ExistingAppAssert) (h.ExistingAppAssert, uuid.UUID, *map[string]interface{}) {
	if existingAppAssert == nil {
		newAppAssert := h.InitTest(t)
		existingAppAssert = &newAppAssert
	}

	var uuid uuid.UUID

	sampleStudent, rawPassword := h.SampleStudentFactory()

	existingAppAssert.TestOnStatusAndTester(h.TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/users/",
		Body:   h.SampleStudentJSONFactory(sampleStudent, rawPassword),
		Role:   &models.Super,
	},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				uuid = AssertSampleUserBodyRespDB(eaa, resp)
			},
		},
	)

	return *existingAppAssert, uuid, h.SampleStudentJSONFactory(sampleStudent, rawPassword)
}

func AssertNumUsersRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var users []models.User

	err := eaa.App.Conn.Find(&users).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(users))
}

var TestNumUsersRemainsAt1 = func(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumUsersRemainsAtN(eaa, resp, 1)
}

var TestNumUsersRemainsAt2 = func(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumUsersRemainsAtN(eaa, resp, 2)
}

func TestCreateUserWorks(t *testing.T) {
	t.Parallel()
	appAssert, _, _ := CreateSampleStudent(t, nil)
	appAssert.Close()
}

func TestCreateUserFailsIfUserWithEmailAlreadyExists(t *testing.T) {
	t.Parallel()
	appAssert, studentUUID, body := CreateSampleStudent(t, nil)

	(*body)["id"] = studentUUID

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/users/",
			Body:   body,
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error:  errors.UserAlreadyExists,
			Tester: TestNumUsersRemainsAt2,
		},
	).Close()
}

func TestCreateUserFailsIfUserWithNUIDAlreadyExists(t *testing.T) {
	t.Parallel()
	appAssert, _, _ := CreateSampleStudent(t, nil)

	sampleStudent, rawPassword := h.SampleStudentFactory()

	slightlyDifferentSampleStudentJSON := h.SampleStudentJSONFactory(sampleStudent, rawPassword)

	(*slightlyDifferentSampleStudentJSON)["first_name"] = "John"
	(*slightlyDifferentSampleStudentJSON)["last_name"] = "Doe"
	(*slightlyDifferentSampleStudentJSON)["email"] = "doe.john@northeastern.edu"

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/users/",
			Body:   slightlyDifferentSampleStudentJSON,
			Role:   &models.Super,
		},
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

		appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/users/",
				Body:   &sampleUserPermutation,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToValidateUser,
				Tester: TestNumUsersRemainsAt2,
			},
		)
	}

	appAssert.Close()
}

func TestCreateUserFailsOnInvalidNUID(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	AssertCreateBadDataFails(t,
		"year",
		[]interface{}{
			0,
			7,
		})
}

func TestCreateUserFailsOnInvalidCollege(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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

		appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/users/",
				Body:   &sampleUserPermutation,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToValidateUser,
				Tester: TestNumUsersRemainsAt2,
			},
		)
	}

	appAssert.Close()
}
