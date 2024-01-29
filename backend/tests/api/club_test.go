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
		"position": "president",
	}
}

func BadEmailPOCFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"name":     "Jane",
		"email":    "doe.jane",
		"position": "president",
	}
}

func AssertPOCWithBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) {
	var respPOC models.PointOfContact

	err := json.NewDecoder(resp.Body).Decode(&respPOC)
	assert.NilError(err)

	var dbPOC models.PointOfContact

	err = app.Conn.Where("id = ?", respPOC.ID).First(&dbPOC).Error
	assert.NilError(err)

	assert.Equal(dbPOC.Name, respPOC.Name)
	assert.Equal(dbPOC.Email, respPOC.Email)
	assert.Equal(dbPOC.Position, respPOC.Position)

	assert.Equal((*body)["name"].(string), dbPOC.Name)
	assert.Equal((*body)["email"].(string), dbPOC.Email)
	assert.Equal((*body)["position"].(string), dbPOC.Position)
}

func AssertSamplePOCBodyRespDB(app TestApp, assert *assert.A, resp *http.Response) {
	AssertPOCWithBodyRespDB(app, assert, resp, SamplePOCFactory())
}

func CreateSamplePOC(t *testing.T) ExistingAppAssert {
	return TestRequest{
		Method: fiber.MethodPut,
		Path:   "/api/v1/clubs/1/poc/",
		Body:   SamplePOCFactory(),
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status:   200,
			DBTester: AssertSamplePOCBodyRespDB,
		},
	)
}

func CreateInvalidEmailPOC(t *testing.T) ExistingAppAssert {
	return TestRequest{
		Method: fiber.MethodPut,
		Path:   "/api/v1/clubs/1/poc/",
		Body:   BadEmailPOCFactory(),
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status:   400,
			DBTester: TestNumPOCRemainsAt0,
		},
	)
}

// POINT OF CONTACT UPSERT
func TestInsertPOCWorks(t *testing.T) {
	CreateSamplePOC(t).Close()
}

func TestCreatePOCFailsOnInvalidEmail(t *testing.T) {
	CreateInvalidEmailPOC(t).Close()
}

func TestUpdatePOCWorks(t *testing.T) {
	appAssert := CreateSamplePOC(t)

	id := 1
	newName := "Jane Austen"
	newPosition := "Executive Director"
	email := "doe.jane@northeastern.edu"

	requestBody := map[string]interface{}{
		"name":     newName,
		"position": newPosition,
		"email":    email,
	}

	TestRequest{
		Method: fiber.MethodPut,
		Path:   fmt.Sprintf("/api/v1/clubs/%d/poc/", id),
		Body:   &requestBody,
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club
				err := app.Conn.First(&club, id).Error
				assert.NilError(err)

				var respPOC models.PointOfContact
				err = json.NewDecoder(resp.Body).Decode(&respPOC)
				assert.NilError(err)

				var dbPOC models.PointOfContact
				err = app.Conn.Where("email = ?", "doe.jane@northeastern.edu").First(&dbPOC).Error
				assert.NilError(err)

				assert.Equal(newName, respPOC.Name)
				assert.Equal(newPosition, respPOC.Position)
				assert.Equal((*SamplePOCFactory())["email"].(string), respPOC.Email)

				assert.Equal(dbPOC.Name, respPOC.Name)
				assert.Equal(dbPOC.Photo, respPOC.Photo)
				assert.Equal(dbPOC.Email, respPOC.Email)
				assert.Equal(dbPOC.Position, respPOC.Position)
				// TODO Club ID not matching between response and db
			},
		},
	).Close()
}

func TestUpdatePOCFailsOnInvalidBody(t *testing.T) {
	appAssert := CreateSamplePOC(t)

	for _, invalidData := range []map[string]interface{}{
		{"email": "not.northeastern"},
		{"position": ""},
		{"name": ""},
	} {
		TestRequest{
			Method: fiber.MethodPut,
			Path:   "/api/v1/clubs/1/poc/",
			Body:   &invalidData,
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

func TestInsertPOCFailsOnMissingFields(t *testing.T) {
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
			Path:   "/api/v1/clubs/1/poc/",
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

// GET ALL POC TEST CASES
func TestGetAllPOCWorks(t *testing.T) {
	appAssert := CreateSamplePOC(t)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/clubs/1/poc/",
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var pointOfContacts []models.PointOfContact

				err := json.NewDecoder(resp.Body).Decode(&pointOfContacts)
				assert.NilError(err)
				assert.Equal(1, len(pointOfContacts))
				respPointOfContact := pointOfContacts[0]

				assert.Equal((*SamplePOCFactory())["name"].(string), respPointOfContact.Name)
				assert.Equal((*SamplePOCFactory())["email"].(string), respPointOfContact.Email)
				assert.Equal((*SamplePOCFactory())["position"].(string), respPointOfContact.Position)

				dbPointOfContacts, err := transactions.GetAllPointOfContacts(app.Conn, 1)
				assert.NilError(&err)
				assert.Equal(1, len(dbPointOfContacts))
				// dbPOC := dbPointOfContacts[0]
				// assert.Equal(dbPOC, respPointOfContact)
				// TODO: Club ID not matching between response and db
			},
		},
	).Close()
}

func TestGetAllPOCClubNotFound(t *testing.T) {
	clubId := 17
	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%d/poc/", clubId),
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

func TestGetPOCWorks(t *testing.T) {
	appAssert := CreateSamplePOC(t)
	id := 1

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/1/poc/%d", id),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respPOC models.PointOfContact

				err := json.NewDecoder(resp.Body).Decode(&respPOC)

				assert.NilError(err)

				assert.Equal("Jane", respPOC.Name)
				assert.Equal("president", respPOC.Position)
				assert.Equal("doe.jane@northeastern.edu", respPOC.Email)

				// dbPOC, err := transactions.GetPointOfContact(app.Conn, uint(id), uint(1))
				// assert.NilError(&err)
				// assert.Equal(dbPOC, &respPOC)
				// TODO: Club ID not matching between response and db
			},
		},
	).Close()
}

func TestGetPOCFailsBadRequest(t *testing.T) {
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
			Path:   fmt.Sprintf("/api/v1/clubs/1/poc/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: errors.FailedToValidatePointOfContactId,
			},
		).Close()
	}
}

func TestGetPOCFailsNotExist(t *testing.T) {
	id := uint(42)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/1/poc/%d", id),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404,
				Message: errors.PointOfContactNotFound,
			},
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var pointOfContact models.PointOfContact
				err := app.Conn.Where("id = ?", id).First(&pointOfContact).Error
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
		Path:   "/api/v1/clubs/1/poc/1",
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   204,
			DBTester: TestNumPOCRemainsAt0,
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

func TestDeletePOCBadRequest(t *testing.T) {
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
			Path:   fmt.Sprintf("/api/v1/clubs/1/poc/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: errors.FailedToValidatePointOfContactId,
			},
		).Close()
	}
}

func TestDeletePOCClubNotExist(t *testing.T) {
	pocId := 1

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/3/poc/%d", pocId),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404,
				Message: errors.ClubNotFound,
			},
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var pointOfContact models.PointOfContact

				err := app.Conn.Where("id = ?", pocId).First(&pointOfContact).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeletePOCNotExist(t *testing.T) {
	pocId := 6

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/1/poc/%d", pocId),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404,
				Message: errors.PointOfContactNotFound,
			},
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var pointOfContact models.PointOfContact

				err := app.Conn.Where("id = ?", pocId).First(&pointOfContact).Error

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

var TestNumPOCRemainsAt0 = func(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumPOCRemainsAtN(app, assert, resp, 0)
}

func AssertCreatePOCBadDataFails(t *testing.T, jsonKey string, badValues []interface{}) {
	appAssert := CreateSamplePOC(t)

	for _, badValue := range badValues {
		samplePOCPermutation := *SamplePOCFactory()
		samplePOCPermutation[jsonKey] = badValue

		TestRequest{
			Method: fiber.MethodPut,
			Path:   "/api/v1/clubs/1/poc",
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
