package root

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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
	eaa.Assert.Equal(dbUser.Major0, respUser.Major0)
	eaa.Assert.Equal(dbUser.Major1, respUser.Major1)
	eaa.Assert.Equal(dbUser.Major2, respUser.Major2)
	eaa.Assert.Equal(dbUser.College, respUser.College)
	eaa.Assert.Equal(dbUser.GraduationCycle, respUser.GraduationCycle)
	eaa.Assert.Equal(dbUser.GraduationYear, respUser.GraduationYear)

	match, err := auth.CompareHash((*body)["password"].(string), dbUser.PasswordHash)

	eaa.Assert.NilError(err)

	eaa.Assert.Assert(match)

	eaa.Assert.Equal((*body)["first_name"].(string), dbUser.FirstName)
	eaa.Assert.Equal((*body)["last_name"].(string), dbUser.LastName)
	eaa.Assert.Equal((*body)["email"].(string), dbUser.Email)
	eaa.Assert.Equal(models.Major((*body)["major0"].(string)), dbUser.Major0)
	eaa.Assert.Equal(models.Major((*body)["major1"].(string)), dbUser.Major1)
	eaa.Assert.Equal(models.Major((*body)["major2"].(string)), dbUser.Major2)
	eaa.Assert.Equal(models.College((*body)["college"].(string)), dbUser.College)
	eaa.Assert.Equal(models.GraduationCycle((*body)["graduation_cycle"].(string)), dbUser.GraduationCycle)
	eaa.Assert.Equal(int16((*body)["graduation_year"].(int)), dbUser.GraduationYear)

	return dbUser.ID
}

func AssertSampleUserBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response) uuid.UUID {
	sampleStudent, rawPassword := h.SampleStudentFactory()

	return AssertUserWithIDBodyRespDB(eaa, resp, h.SampleStudentJSONFactory(sampleStudent, rawPassword))
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
