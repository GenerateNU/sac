package tests

import (
	stdliberrors "errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-assert"
	"gorm.io/gorm"
)

func SamplePOCFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"name":     "Jane",
		"email":    "doe.jane@northeastern.edu",
		"photo":    "google.com",
		"position": "president",
	}
}

func AssertPOCWithBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) {
	var respPOC models.PointOfContact

	err := json.NewDecoder(resp.Body).Decode(&respPOC)
	assert.NilError(err)

	var dbPOC models.PointOfContact

	// Assuming Email is a unique identifier in your model
	err = app.Conn.Where("email = ?", respPOC.Email).First(&dbPOC).Error
	assert.NilError(err)

	assert.Equal(dbPOC.Name, respPOC.Name)
	assert.Equal(dbPOC.Email, respPOC.Email)
	assert.Equal(dbPOC.Photo, respPOC.Photo)
	assert.Equal(dbPOC.Position, respPOC.Position)

	assert.Equal((*body)["name"].(string), dbPOC.Name)
	assert.Equal((*body)["email"].(string), dbPOC.Email)
	assert.Equal((*body)["photo"].(string), dbPOC.Photo)
	assert.Equal((*body)["position"].(string), dbPOC.Position)
}

func AssertSamplePOCBodyRespDB(app TestApp, assert *assert.A, resp *http.Response) {
    AssertPOCWithBodyRespDB(app, assert, resp, SamplePOCFactory())
}

func CreateSamplePOC(t *testing.T) ExistingAppAssert {
	return TestRequest{
		Method: fiber.MethodPut,
		Path:   "/api/v1/clubs/1/poc",
		Body:   SamplePOCFactory(),
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status:   201,
			DBTester: AssertSamplePOCBodyRespDB,
		},
	)
}

// POINT OF CONTACT UPSERT
func TestInsertPOCWorks(t *testing.T) {
	CreateSamplePOC(t).Close()
}

func TestUpdatePOCWorks(t *testing.T) {
	appAssert := CreateSamplePOC(t)

	clubId := 2
	newName := "Jane"
	newPosition := "Executive Director"
	email := "doe.jane@northeastern.edu"

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/clubs/%d/poc", clubId),
		Body: &map[string]interface{}{
			"name":     newName,
			"position": newPosition,
			"email":    email,
		},
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respPOC models.PointOfContact

				err := json.NewDecoder(resp.Body).Decode(&respPOC)

				assert.NilError(err)

				assert.Equal(newName, respPOC.Name)
				assert.Equal(newPosition, respPOC.Position)
				assert.Equal((*SamplePOCFactory())["name"].(string), respPOC.Name)
				assert.Equal((*SamplePOCFactory())["position"].(string), respPOC.Position)

				var dbClub models.Club

				err = app.Conn.First(&dbClub, clubId).Error
				assert.NilError(err)

				var dbPOC models.PointOfContact

				assert.Equal(dbPOC.Name, respPOC.Name)
				assert.Equal(dbPOC.Email, respPOC.Email)
				assert.Equal(dbPOC.Photo, respPOC.Photo)
				assert.Equal(dbPOC.Position, respPOC.Position)
			},
		},
	).Close()
}

func TestInsertPOCFailsOnInvalidEmail(t *testing.T) {
	AssertCreatePOCBadDataFails(t,
		"email",
		[]interface{}{
			"doe.jane@mail",
			"doe.jane",
			"doe.jane@",
			"doe.jane@northeastern.",
			"doe.jane@gmail.c",
			"",
		})
}

func TestInsertUserFailsOnMissingFields(t *testing.T) {
	appAssert := CreateSamplePOC(t)

	for _, missingField := range []string{
		"name",
		"email",
		"position",
	} {
		samplePOCPermutation := *SamplePOCFactory()
		delete(samplePOCPermutation, missingField)
		TestRequest{
			Method: fiber.MethodPut,
			Path:   "/api/v1/clubs/2/poc",
			Body:   &samplePOCPermutation,
		}.TestOnStatusMessageAndDB(t, &appAssert,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400,
					Message: errors.FailedToValidatePointOfContact,
				},
				DBTester: TestNumPOCRemainsAt1,
			},
		)
	}
	appAssert.Close()
}


// GET ALL TEST CASES
func TestGetAllPOCWorks(t *testing.T) {
	TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/clubs/1/poc",
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var pointOfContacts []models.PointOfContact

				err := json.NewDecoder(resp.Body).Decode(&pointOfContacts)

				assert.NilError(err)

				assert.Equal(1, len(pointOfContacts))

				respPointOfContact := pointOfContacts[0]

				assert.Equal("Jane", respPointOfContact.Name)
				assert.Equal("doe.jane@northeastern.edu", respPointOfContact.Email)
				assert.Equal("google.com", respPointOfContact.Photo)
				assert.Equal("president", respPointOfContact.Position)

				dbPointOfContacts, err := transactions.GetAllPointOfContacts(app.Conn, 2)

				assert.NilError(&err)

				assert.Equal(1, len(dbPointOfContacts))

				dbPointOfContact := dbPointOfContacts[0]

				assert.Equal(dbPointOfContact, respPointOfContact)
			},
		},
	).Close()
}

func TestGetAllPOCClubNotFound(t *testing.T) {
	clubId := 17

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/%d/poc", clubId),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404,
				Message: errors.ClubNotFound,
			},
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club
				err := app.Conn.Where("id = ?", clubId).First(&club).Error
				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

// DELETE TEST CASES
func TestDeletePointOfContactWorks(t *testing.T) {
	appAssert := CreateSamplePOC(t)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   "/api/v1/users/1/poc/doe.jane@northeastern.edu",
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   204,
			DBTester: TestNumPOCRemainsAt1,
		},
	).Close()
}

func TestDeletePOCClubIDBadRequest(t *testing.T) {
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
			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/doe.jane@northeastern.edu", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: errors.FailedToValidateClubId,
			},
		).Close()
	}
}

func TestDeletePOCEmailBadRequest(t *testing.T) {
	badRequests := []string{
		"1",
		"hello@gmail",
		"hello",
		"1.1",
		"null",
	}

	for _, badRequest := range badRequests {
		TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/clubs/1/poc/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: errors.FailedToValidateEmail,
			},
		).Close()
	}
}

func TestDeletePOCClubNotExist(t *testing.T) {
	email := "doe.jane@northeastern.edu"

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/3/poc/%s", email),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404,
				Message: errors.ClubNotFound,
			},
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var pointOfContact models.PointOfContact

				err := app.Conn.Where("email = ?", email).First(&pointOfContact).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeletePOCNotExist(t *testing.T) {
	email := "doe.john@northeastern.edu"

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/1/poc/%s", email),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404,
				Message: errors.PointOfContactNotFound,
			},
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var pointOfContact models.PointOfContact

				err := app.Conn.Where("email = ?", email).First(&pointOfContact).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}


// assert remaining numbers of POC
func AssertNumPOCRemainsAtN(app TestApp, assert *assert.A, resp *http.Response, n int) {
	var pointOfContact []models.PointOfContact

	err := app.Conn.Find(&pointOfContact).Error

	assert.NilError(err)

	assert.Equal(n, len(pointOfContact))
}

// assert remaining POC = 1
var TestNumPOCRemainsAt1 = func(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumPOCRemainsAtN(app, assert, resp, 1)
}

// assert remaining POC = 2
var TestNumPOCRemainsAt2 = func(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumPOCRemainsAtN(app, assert, resp, 2)
}

func AssertCreatePOCBadDataFails(t *testing.T, jsonKey string, badValues []interface{}) {
	appAssert := CreateSamplePOC(t)

	for _, badValue := range badValues {
		samplePOCPermutation := *SamplePOCFactory()
		samplePOCPermutation[jsonKey] = badValue

		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/clubs/1/poc",
			Body:   &samplePOCPermutation,
		}.TestOnStatusMessageAndDB(t, &appAssert,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400,
					Message: errors.FailedToValidatePointOfContact,
				},
				DBTester: TestNumPOCRemainsAt2,
			},
		)
	}
	appAssert.Close()
}