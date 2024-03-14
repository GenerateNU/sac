package tests

// import (
// 	"fmt"
// 	"net/http"
// 	"testing"

// 	"github.com/GenerateNU/sac/backend/src/errors"
// 	"github.com/GenerateNU/sac/backend/src/models"
// 	"github.com/GenerateNU/sac/backend/src/transactions"
// 	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
// 	"github.com/goccy/go-json"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"github.com/huandu/go-assert"
// 	"gorm.io/gorm"
// )

// func TestUpdatePOCFailsOnInvalidBody(t *testing.T) {
// 	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
// 	appAssert, _ := CreateSamplePOC(t, existingAppAssert)

// 	for _, invalidData := range []map[string]interface{}{
// 		{"email": "not.northeastern"},
// 		{"position": ""},
// 		{"name": ""},
// 	} {
// 		invalidData := invalidData
// 		appAssert.TestOnErrorAndTester(
// 			h.TestRequest{
// 				Method: fiber.MethodPut,
// 				Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
// 				Body:   &invalidData,
// 				Role:   &models.Super,
// 			},
// 			h.ErrorWithTester{
// 				Error:  errors.FailedToUpsertPointOfContact,
// 				Tester: TestNumPOCRemainsAt1,
// 			},
// 		)
// 	}
// 	appAssert.Close()
// }

// func SamplePOCFactory() *map[string]interface{} {
// 	return &map[string]interface{}{
// 		"name":     "Jane",
// 		"email":    "doe.jane@northeastern.edu",
// 		"position": "president",
// 	}
// }

// func BadEmailPOCFactory() *map[string]interface{} {
// 	return &map[string]interface{}{
// 		"name":     "Jane",
// 		"email":    "doe.jane",
// 		"position": "president",
// 	}
// }

// func AssertPOCWithBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) {
// 	var respPOC models.PointOfContact

// 	err := json.NewDecoder(resp.Body).Decode(&respPOC)
// 	eaa.Assert.NilError(err)

// 	var dbPOC models.PointOfContact

// 	err = eaa.App.Conn.Where("id = ?", respPOC.ID).First(&dbPOC).Error
// 	eaa.Assert.NilError(err)

// 	eaa.Assert.Equal(dbPOC.Name, respPOC.Name)
// 	eaa.Assert.Equal(dbPOC.Email, respPOC.Email)
// 	eaa.Assert.Equal(dbPOC.Position, respPOC.Position)

// 	eaa.Assert.Equal((*body)["name"].(string), dbPOC.Name)
// 	eaa.Assert.Equal((*body)["email"].(string), dbPOC.Email)
// 	eaa.Assert.Equal((*body)["position"].(string), dbPOC.Position)
// }

// // func AssertSamplePOCBodyRespDB(app h.TestApp, assert *assert.A, resp *http.Response) {
// // 	AssertPOCWithBodyRespDB(app, assert, resp, SamplePOCFactory())
// // }

// func CreateSamplePOC(t *testing.T, existingAppAssert h.ExistingAppAssert) (h.ExistingAppAssert, uuid.UUID) {
// 	_, _, clubUUID := CreateSampleClub(h.InitTest(t))
// 	var pocId uuid.UUID

// 	newAppAssert := existingAppAssert.TestOnStatusAndTester(
// 		h.TestRequest{
// 			Method: fiber.MethodPut,
// 			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
// 			Body:   SamplePOCFactory(),
// 		},
// 		h.TesterWithStatus{
// 			Status: fiber.StatusCreated,
// 			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
// 				var respPOC models.PointOfContact
// 				err := json.NewDecoder(resp.Body).Decode(&respPOC)
// 				eaa.Assert.NilError(err)
// 				pocId = respPOC.ID
// 			},
// 		},
// 	)
// 	return newAppAssert, pocId
// }

// func CreateInvalidEmailPOC(t *testing.T) h.ExistingAppAssert {
// 	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

// 	appAssert.TestOnErrorAndTester(
// 		h.TestRequest{
// 			Method: fiber.MethodPut,
// 			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
// 			Body:   BadEmailPOCFactory(),
// 			Role:   &models.Super,
// 		},
// 		h.ErrorWithTester{
// 			Error:  errors.FailedToValidateEmail,
// 			Tester: TestNumPOCRemainsAt0,
// 		},
// 	).Close()
// 	return appAssert
// }

// // POINT OF CONTACT UPSERT
// func TestInsertPOCWorks(t *testing.T) {
// 	existingAppAssert, _, _ := CreateSampleClub(h.InitTest(t))
// 	appAssert, _ := CreateSamplePOC(t, existingAppAssert)
// 	appAssert.Close()
// }

// func TestCreatePOCFailsOnInvalidEmail(t *testing.T) {
// 	CreateInvalidEmailPOC(t).Close()
// }

// func TestUpdatePOCWorks(t *testing.T) {
// 	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
// 	appAssert, _ := CreateSamplePOC(t, existingAppAssert)

// 	newName := "Jane Austen"
// 	newPosition := "Executive Director"
// 	email := "doe.jane@northeastern.edu"

// 	requestBody := map[string]interface{}{
// 		"name":     newName,
// 		"position": newPosition,
// 		"email":    email,
// 	}

// 	updatedPOC := SamplePOCFactory()
// 	(*updatedPOC)["name"] = newName
// 	(*updatedPOC)["position"] = newPosition

// 	appAssert.TestOnStatusAndTester(
// 		h.TestRequest{
// 			Method: fiber.MethodPatch,
// 			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
// 			Body:   &requestBody,
// 			Role:   &models.Super,
// 		},
// 		h.TesterWithStatus{
// 			Status: fiber.StatusOK,
// 			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
// 				AssertPOCWithBodyRespDB(eaa, resp, updatedPOC)
// 			},
// 		},
// 	).Close()
// }
// func TestInsertPOCFailsOnMissingFields(t *testing.T) {
// 	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
// 	appAssert, _ := CreateSamplePOC(t, existingAppAssert)

// 	for _, missingField := range []string{
// 		"name",
// 		"email",
// 		"position",
// 	} {
// 		samplePOCPermutation := *SamplePOCFactory()
// 		delete(samplePOCPermutation, missingField)
// 		TestRequest{
// 			Method: fiber.MethodPut,
// 			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
// 			Body:   &samplePOCPermutation,
// 		}.TestOnStatusMessageAndDB(t, &appAssert,
// 			StatusMessageDBTester{
// 				MessageWithStatus: MessageWithStatus{
// 					Status:  400,
// 					Message: errors.FailedToValidatePointOfContact,
// 				},
// 				DBTester: TestNumPOCRemainsAt1,
// 			},
// 		)
// 	}
// 	appAssert.Close()
// }

// // GET ALL POC TEST CASES
// func TestGetAllPOCWorks(t *testing.T) {
// 	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
// 	appAssert, _ := CreateSamplePOC(t, existingAppAssert)

// 	TestRequest{
// 		Method: fiber.MethodGet,
// 		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
// 	}.TestOnStatusAndDB(t, &appAssert,
// 		DBTesterWithStatus{
// 			Status: 200,
// 			DBTester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
// 				var pointOfContacts []models.PointOfContact

// 				err := json.NewDecoder(resp.Body).Decode(&pointOfContacts)
// 				assert.NilError(err)
// 				assert.Equal(1, len(pointOfContacts))
// 				respPointOfContact := pointOfContacts[0]

// 				assert.Equal((*SamplePOCFactory())["name"].(string), respPointOfContact.Name)
// 				assert.Equal((*SamplePOCFactory())["email"].(string), respPointOfContact.Email)
// 				assert.Equal((*SamplePOCFactory())["position"].(string), respPointOfContact.Position)

// 				dbPointOfContacts, err := transactions.GetAllPointOfContacts(app.Conn, 1)
// 				assert.NilError(&err)
// 				assert.Equal(1, len(dbPointOfContacts))
// 				// dbPOC := dbPointOfContacts[0]
// 				// assert.Equal(dbPOC, respPointOfContact)
// 				// TODO: Club ID not matching between response and db
// 			},
// 		},
// 	).Close()
// }

// func TestGetAllPOCClubNotFound(t *testing.T) {
// 	clubId := uuid.New()

// 	h.TestRequest{
// 		Method: fiber.MethodGet,
// 		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubId),
// 	}.TestOnStatusMessageAndDB(t, nil,
// 		StatusMessageDBTester{
// 			MessageWithStatus: MessageWithStatus{
// 				Status:  404,
// 				Message: errors.ClubNotFound,
// 			},
// 			DBTester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
// 				var club models.Club
// 				err := app.Conn.Where("id = ?", clubId).First(&club).Error
// 				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
// 			},
// 		},
// 	).Close()
// }

// func TestGetPOCWorks(t *testing.T) {
// 	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
// 	appAssert, pocUUID := CreateSamplePOC(t, existingAppAssert)

// 	h.TestRequest{
// 		Method: fiber.MethodGet,
// 		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, pocUUID),
// 	}.TestOnStatusAndDB(t, &appAssert,
// 		DBTesterWithStatus{
// 			Status: 200,
// 			DBTester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
// 				var respPOC models.PointOfContact

// 				err := json.NewDecoder(resp.Body).Decode(&respPOC)

// 				assert.NilError(err)

// 				assert.Equal("Jane", respPOC.Name)
// 				assert.Equal("president", respPOC.Position)
// 				assert.Equal("doe.jane@northeastern.edu", respPOC.Email)

// 				// dbPOC, err := transactions.GetPointOfContact(app.Conn, uint(id), uint(1))
// 				// assert.NilError(&err)
// 				// assert.Equal(dbPOC, &respPOC)
// 				// TODO: Club ID not matching between response and db
// 			},
// 		},
// 	).Close()
// }

// func TestGetPOCFailsBadRequest(t *testing.T) {
// 	_, _, clubUUID := CreateSampleClub(h.InitTest(t))

// 	badRequests := []string{
// 		"0",
// 		"-1",
// 		"1.1",
// 		"foo",
// 		"null",
// 	}

// 	for _, badRequest := range badRequests {
// 		TestRequest{
// 			Method: fiber.MethodGet,
// 			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, badRequest),
// 		}.TestOnStatusAndMessage(t, nil,
// 			MessageWithStatus{
// 				Status:  400,
// 				Message: errors.FailedToValidatePointOfContactId,
// 			},
// 		).Close()
// 	}
// }

// func TestGetPOCFailsNotExist(t *testing.T) {
// 	_, _, clubUUID := CreateSampleClub(h.InitTest(t))
// 	pocUUIDNotExist := uuid.New()

// 	TestRequest{
// 		Method: fiber.MethodGet,
// 		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, pocUUIDNotExist),
// 	}.TestOnStatusMessageAndDB(t, nil,
// 		StatusMessageDBTester{
// 			MessageWithStatus: MessageWithStatus{
// 				Status:  404,
// 				Message: errors.PointOfContactNotFound,
// 			},
// 			DBTester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
// 				var pointOfContact models.PointOfContact
// 				err := app.Conn.Where("id = ?", pocUUIDNotExist).First(&pointOfContact).Error
// 				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
// 			},
// 		},
// 	).Close()
// }

// // DELETE TEST CASES
// func TestDeletePointOfContactWorks(t *testing.T) {
// 	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
// 	appAssert, pocUUID := CreateSamplePOC(t, existingAppAssert)

// 	TestRequest{
// 		Method: fiber.MethodDelete,
// 		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, pocUUID),
// 	}.TestOnStatusAndDB(t, &appAssert,
// 		DBTesterWithStatus{
// 			Status:   204,
// 			DBTester: TestNumPOCRemainsAt0,
// 		},
// 	).Close()
// }

// func TestDeletePOCClubIDBadRequest(t *testing.T) {
// 	existingAppAssert, _, _ := CreateSampleClub(h.InitTest(t))
// 	_, pocUUID := CreateSamplePOC(t, existingAppAssert)
// 	badRequests := []string{
// 		"0",
// 		"-1",
// 		"1.1",
// 		"hello",
// 		"null",
// 	}

// 	for _, badRequest := range badRequests {
// 		TestRequest{
// 			Method: fiber.MethodDelete,
// 			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", badRequest, pocUUID),
// 		}.TestOnStatusAndMessage(t, nil,
// 			MessageWithStatus{
// 				Status:  400,
// 				Message: errors.FailedToValidateClub,
// 			},
// 		).Close()
// 	}
// }

// func TestDeletePOCBadRequest(t *testing.T) {
// 	_, _, clubUUID := CreateSampleClub(h.InitTest(t))

// 	badRequests := []string{
// 		"0",
// 		"-1",
// 		"1.1",
// 		"hello",
// 		"null",
// 	}

// 	for _, badRequest := range badRequests {
// 		TestRequest{
// 			Method: fiber.MethodDelete,
// 			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, badRequest),
// 		}.TestOnStatusAndMessage(t, nil,
// 			MessageWithStatus{
// 				Status:  400,
// 				Message: errors.FailedToValidatePointOfContactId,
// 			},
// 		).Close()
// 	}
// }

// func TestDeletePOCClubNotExist(t *testing.T) {
// 	existingAppAssert, _, _ := CreateSampleClub(h.InitTest(t))
// 	_, pocUUID := CreateSamplePOC(t, existingAppAssert)
// 	clubUUID := uuid.New()

// 	TestRequest{
// 		Method: fiber.MethodDelete,
// 		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, pocUUID),
// 	}.TestOnStatusMessageAndDB(t, nil,
// 		StatusMessageDBTester{
// 			MessageWithStatus: MessageWithStatus{
// 				Status:  404,
// 				Message: errors.ClubNotFound,
// 			},
// 			DBTester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
// 				var club models.Club

// 				err := app.Conn.Where("id = ?", clubUUID).First(&club).Error

// 				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
// 			},
// 		},
// 	).Close()
// }

// func TestDeletePOCNotExist(t *testing.T) {
// 	_, _, clubUUID := CreateSampleClub(h.InitTest(t))
// 	pocUUID := uuid.New()

// 	TestRequest{
// 		Method: fiber.MethodDelete,
// 		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, pocUUID),
// 	}.TestOnStatusMessageAndDB(t, nil,
// 		StatusMessageDBTester{
// 			MessageWithStatus: MessageWithStatus{
// 				Status:  404,
// 				Message: errors.PointOfContactNotFound,
// 			},
// 			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
// 				var pointOfContact models.PointOfContact

// 				err := app.Conn.Where("id = ?", pocUUID).First(&pointOfContact).Error

// 				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
// 			},
// 		},
// 	).Close()
// }

// // assert remaining numbers of POC
// func AssertNumPOCRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
// 	var pointOfContact []models.PointOfContact

// 	err := eaa.App.Conn.Find(&pointOfContact).Error

// 	eaa.Assert.NilError(err)

// 	eaa.Assert.Equal(n, len(pointOfContact))
// }

// // assert remaining POC = 1
// var TestNumPOCRemainsAt1 = func(eaa h.ExistingAppAssert, resp *http.Response) {
// 	AssertNumPOCRemainsAtN(eaa, resp, 1)
// }

// var TestNumPOCRemainsAt0 = func(eaa h.ExistingAppAssert, resp *http.Response) {
// 	AssertNumPOCRemainsAtN(eaa, resp, 0)
// }