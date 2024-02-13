package tests

import (
	stdliberrors "errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func AssertPOCWithBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) {
	var respPOC models.PointOfContact

	err := json.NewDecoder(resp.Body).Decode(&respPOC)
	eaa.Assert.NilError(err)

	var dbPOC models.PointOfContact

	err = eaa.App.Conn.Where("id = ?", respPOC.ID).First(&dbPOC).Error
	eaa.Assert.NilError(err)

	eaa.Assert.Equal(dbPOC.Name, respPOC.Name)
	eaa.Assert.Equal(dbPOC.Email, respPOC.Email)
	eaa.Assert.Equal(dbPOC.Position, respPOC.Position)

	eaa.Assert.Equal((*body)["name"].(string), dbPOC.Name)
	eaa.Assert.Equal((*body)["email"].(string), dbPOC.Email)
	eaa.Assert.Equal((*body)["position"].(string), dbPOC.Position)
}

// func AssertSamplePOCBodyRespDB(app h.TestApp, assert *assert.A, resp *http.Response) {
// 	AssertPOCWithBodyRespDB(app, assert, resp, SamplePOCFactory())
// }

func CreateSamplePOC(t *testing.T, existingAppAssert h.ExistingAppAssert) (h.ExistingAppAssert, uuid.UUID) {
	_, _, clubUUID := CreateSampleClub(h.InitTest(t))
	var pocId uuid.UUID

	newAppAssert := existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
			Body:   SamplePOCFactory(),
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var respPOC models.PointOfContact
				err := json.NewDecoder(resp.Body).Decode(&respPOC)
				eaa.Assert.NilError(err)
				pocId = respPOC.ID
			},
		},
	)
	return newAppAssert, pocId
}

func CreateInvalidEmailPOC(t *testing.T) h.ExistingAppAssert {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
			Body:   BadEmailPOCFactory(),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error:  errors.FailedToValidateEmail,
			Tester: TestNumPOCRemainsAt0,
		},
	).Close()
	return appAssert
}

// POINT OF CONTACT UPSERT
func TestInsertPOCWorks(t *testing.T) {
	existingAppAssert, _, _ := CreateSampleClub(h.InitTest(t))
	appAssert, _ := CreateSamplePOC(t, existingAppAssert)
	appAssert.Close()
}

func TestCreatePOCFailsOnInvalidEmail(t *testing.T) {
	CreateInvalidEmailPOC(t).Close()
}

func TestUpdatePOCWorks(t *testing.T) {
	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
	appAssert, _ := CreateSamplePOC(t, existingAppAssert)

	newName := "Jane Austen"
	newPosition := "Executive Director"
	email := "doe.jane@northeastern.edu"

	requestBody := map[string]interface{}{
		"name":     newName,
		"position": newPosition,
		"email":    email,
	}

	updatedPOC := SamplePOCFactory()
	(*updatedPOC)["name"] = newName
	(*updatedPOC)["position"] = newPosition

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
			Body:   &requestBody,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertPOCWithBodyRespDB(eaa, resp, updatedPOC)
			},
		},
	).Close()
}

func SampleClubFactory(userID *uuid.UUID) *map[string]interface{} {
	return &map[string]interface{}{
		"user_id":           userID,
		"name":              "Generate",
		"preview":           "Generate is Northeastern's premier student-led product development studio.",
		"description":       "https://mongodb.com",
		"is_recruiting":     true,
		"recruitment_cycle": "always",
		"recruitment_type":  "application",
		"application_link":  "https://generatenu.com/apply",
		"logo":              "https://aws.amazon.com/s3/",
	}
}

func AssertClubBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respClub models.Club

	err := json.NewDecoder(resp.Body).Decode(&respClub)

	eaa.Assert.NilError(err)

	var dbClubs []models.Club

	err = eaa.App.Conn.Order("created_at desc").Find(&dbClubs).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(2, len(dbClubs))

	dbClub := dbClubs[0]

	eaa.Assert.Equal(dbClub.ID, respClub.ID)
	eaa.Assert.Equal(dbClub.Name, respClub.Name)
	eaa.Assert.Equal(dbClub.Preview, respClub.Preview)
	eaa.Assert.Equal(dbClub.Description, respClub.Description)
	eaa.Assert.Equal(dbClub.NumMembers, respClub.NumMembers)
	eaa.Assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
	eaa.Assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
	eaa.Assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
	eaa.Assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
	eaa.Assert.Equal(dbClub.Logo, respClub.Logo)

	var dbAdmins []models.User

	err = eaa.App.Conn.Model(&dbClub).Association("Admin").Find(&dbAdmins)

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(1, len(dbAdmins))

	eaa.Assert.Equal(*(*body)["user_id"].(*uuid.UUID), dbAdmins[0].ID)
	eaa.Assert.Equal((*body)["name"].(string), dbClub.Name)
	eaa.Assert.Equal((*body)["preview"].(string), dbClub.Preview)
	eaa.Assert.Equal((*body)["description"].(string), dbClub.Description)
	eaa.Assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
	eaa.Assert.Equal(models.RecruitmentCycle((*body)["recruitment_cycle"].(string)), dbClub.RecruitmentCycle)
	eaa.Assert.Equal(models.RecruitmentType((*body)["recruitment_type"].(string)), dbClub.RecruitmentType)
	eaa.Assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
	eaa.Assert.Equal((*body)["logo"].(string), dbClub.Logo)

	return dbClub.ID
}

func AssertClubWithBodyRespDBMostRecent(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respClub models.Club

	err := json.NewDecoder(resp.Body).Decode(&respClub)

	eaa.Assert.NilError(err)

	var dbClub models.Club

	err = eaa.App.Conn.Order("created_at desc").First(&dbClub).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(dbClub.ID, respClub.ID)
	eaa.Assert.Equal(dbClub.Name, respClub.Name)
	eaa.Assert.Equal(dbClub.Preview, respClub.Preview)
	eaa.Assert.Equal(dbClub.Description, respClub.Description)
	eaa.Assert.Equal(dbClub.NumMembers, respClub.NumMembers)
	eaa.Assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
	eaa.Assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
	eaa.Assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
	eaa.Assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
	eaa.Assert.Equal(dbClub.Logo, respClub.Logo)

	var dbAdmins []models.User

	err = eaa.App.Conn.Model(&dbClub).Association("Admins").Find(&dbAdmins)

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(1, len(dbAdmins))

	dbAdmin := dbAdmins[0]

	eaa.Assert.Equal((*body)["user_id"].(uuid.UUID), dbAdmin.ID)
	eaa.Assert.Equal((*body)["name"].(string), dbClub.Name)
	eaa.Assert.Equal((*body)["preview"].(string), dbClub.Preview)
	eaa.Assert.Equal((*body)["description"].(string), dbClub.Description)
	eaa.Assert.Equal((*body)["num_members"].(int), dbClub.NumMembers)
	eaa.Assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
	eaa.Assert.Equal((*body)["recruitment_cycle"].(string), dbClub.RecruitmentCycle)
	eaa.Assert.Equal((*body)["recruitment_type"].(string), dbClub.RecruitmentType)
	eaa.Assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
	eaa.Assert.Equal((*body)["logo"].(string), dbClub.Logo)

	return dbClub.ID
}

func AssertSampleClubBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, userID uuid.UUID) uuid.UUID {
	sampleClub := SampleClubFactory(&userID)
	(*sampleClub)["num_members"] = 1

	return AssertClubBodyRespDB(eaa, resp, sampleClub)
}

func CreateSampleClub(existingAppAssert h.ExistingAppAssert) (eaa h.ExistingAppAssert, studentUUID uuid.UUID, clubUUID uuid.UUID) {
	var sampleClubUUID uuid.UUID

	newAppAssert := existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method:             fiber.MethodPost,
			Path:               "/api/v1/clubs/",
			Body:               SampleClubFactory(nil),
			Role:               &models.Super,
			TestUserIDReplaces: h.StringToPointer("user_id"),
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				sampleClubUUID = AssertSampleClubBodyRespDB(eaa, resp, eaa.App.TestUser.UUID)
			},
		},
	)

	return existingAppAssert, newAppAssert.App.TestUser.UUID, sampleClubUUID
}

func TestCreateClubWorks(t *testing.T) {
	existingAppAssert, _, _ := CreateSampleClub(h.InitTest(t))
	existingAppAssert.Close()
}

func TestGetClubsWorks(t *testing.T) {
	h.InitTest(t).TestOnStatusAndTester(h.TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/clubs/",
		Role:   &models.Super,
	},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var respClubs []models.Club

				err := json.NewDecoder(resp.Body).Decode(&respClubs)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(respClubs))

				respClub := respClubs[0]

				var dbClubs []models.Club

				err = eaa.App.Conn.Order("created_at desc").Find(&dbClubs).Error

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(1, len(dbClubs))

				dbClub := dbClubs[0]

				eaa.Assert.Equal(dbClub.ID, respClub.ID)
				eaa.Assert.Equal(dbClub.Name, respClub.Name)
				eaa.Assert.Equal(dbClub.Preview, respClub.Preview)
				eaa.Assert.Equal(dbClub.Description, respClub.Description)
				eaa.Assert.Equal(dbClub.NumMembers, respClub.NumMembers)
				eaa.Assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
				eaa.Assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
				eaa.Assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
				eaa.Assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
				eaa.Assert.Equal(dbClub.Logo, respClub.Logo)

				eaa.Assert.Equal("SAC", dbClub.Name)
				eaa.Assert.Equal("SAC", dbClub.Preview)
				eaa.Assert.Equal("SAC", dbClub.Description)
				eaa.Assert.Equal(1, dbClub.NumMembers)
				eaa.Assert.Equal(true, dbClub.IsRecruiting)
				eaa.Assert.Equal(models.Always, dbClub.RecruitmentCycle)
				eaa.Assert.Equal(models.Application, dbClub.RecruitmentType)
				eaa.Assert.Equal("https://generatenu.com/apply", dbClub.ApplicationLink)
				eaa.Assert.Equal("https://aws.amazon.com/s3", dbClub.Logo)
			},
		},
	).Close()
}

func TestUpdatePOCFailsOnInvalidBody(t *testing.T) {
	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
	appAssert, _ := CreateSamplePOC(t, existingAppAssert)

	for _, invalidData := range []map[string]interface{}{
		{"email": "not.northeastern"},
		{"position": ""},
		{"name": ""},
	} {
		invalidData := invalidData
		appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPut,
				Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
				Body:   &invalidData,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToUpsertPointOfContact,
				Tester: TestNumPOCRemainsAt1,
			},
		)
	}
	appAssert.Close()
}

func AssertNumClubsRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var dbClubs []models.Club

	err := eaa.App.Conn.Order("created_at desc").Find(&dbClubs).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(dbClubs))
}

var TestNumClubsRemainsAt1 = func(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumClubsRemainsAtN(eaa, resp, 1)
}

func AssertCreateBadClubDataFails(t *testing.T, jsonKey string, badValues []interface{}) {
	appAssert, uuid, _ := CreateSampleStudent(t, nil)

	for _, badValue := range badValues {
		sampleClubPermutation := *SampleClubFactory(&uuid)
		sampleClubPermutation[jsonKey] = badValue

		appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/clubs/",
				Body:   &sampleClubPermutation,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToValidateClub,
				Tester: TestNumClubsRemainsAt1,
			},
		)
	}
	appAssert.Close()
}

func TestCreateClubFailsOnInvalidDescription(t *testing.T) {
	AssertCreateBadClubDataFails(t,
		"description",
		[]interface{}{
			"Not an URL",
			"@#139081#$Ad_Axf",
			// "https://google.com", <-- TODO fix once we handle mongo urls
		},
	)
}

func TestCreateClubFailsOnInvalidRecruitmentCycle(t *testing.T) {
	AssertCreateBadClubDataFails(t,
		"recruitment_cycle",
		[]interface{}{
			"1234",
			"garbanzo",
			"@#139081#$Ad_Axf",
			"https://google.com",
		},
	)
}

func TestCreateClubFailsOnInvalidRecruitmentType(t *testing.T) {
	AssertCreateBadClubDataFails(t,
		"recruitment_type",
		[]interface{}{
			"1234",
			"garbanzo",
			"@#139081#$Ad_Axf",
			"https://google.com",
		},
	)
}

func TestCreateClubFailsOnInvalidApplicationLink(t *testing.T) {
	AssertCreateBadClubDataFails(t,
		"application_link",
		[]interface{}{
			"Not an URL",
			"@#139081#$Ad_Axf",
		},
	)
}

func TestCreateClubFailsOnInvalidLogo(t *testing.T) {
	AssertCreateBadClubDataFails(t,
		"logo",
		[]interface{}{
			"Not an URL",
			"@#139081#$Ad_Axf",
			// "https://google.com", <-- TODO uncomment once we figure out s3 url validation
		},
	)
}

func TestUpdateClubWorks(t *testing.T) {
	appAssert, studentUUID, clubUUID := CreateSampleClub(h.InitTest(t))

	updatedClub := SampleClubFactory(&studentUUID)
	(*updatedClub)["name"] = "Updated Name"
	(*updatedClub)["preview"] = "Updated Preview"

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/clubs/%s", clubUUID),
			Body:   updatedClub,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertClubBodyRespDB(eaa, resp, updatedClub)
			},
		},
	).Close()
}

func TestUpdateClubFailsOnInvalidBody(t *testing.T) {
	appAssert, studentUUID, clubUUID := CreateSampleClub(h.InitTest(t))

	body := SampleClubFactory(&studentUUID)

	for _, invalidData := range []map[string]interface{}{
		{"description": "Not a URL"},
		{"recruitment_cycle": "1234"},
		{"recruitment_type": "ALLLLWAYSSSS"},
		{"application_link": "Not an URL"},
		{"logo": "@12394X_2"},
	} {
		invalidData := invalidData
		appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPatch,
				Path:   fmt.Sprintf("/api/v1/clubs/%s", clubUUID),
				Body:   &invalidData,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error: errors.FailedToValidateClub,
				Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
					var dbClubs []models.Club

					err := eaa.App.Conn.Order("created_at desc").Find(&dbClubs).Error

					eaa.Assert.NilError(err)

					eaa.Assert.Equal(2, len(dbClubs))

					dbClub := dbClubs[0]

					var dbAdmins []models.User

					err = eaa.App.Conn.Model(&dbClub).Association("Admin").Find(&dbAdmins)

					eaa.Assert.NilError(err)

					eaa.Assert.Equal(1, len(dbAdmins))

					eaa.Assert.Equal(*(*body)["user_id"].(*uuid.UUID), dbAdmins[0].ID)
					eaa.Assert.Equal((*body)["name"].(string), dbClub.Name)
					eaa.Assert.Equal((*body)["preview"].(string), dbClub.Preview)
					eaa.Assert.Equal((*body)["description"].(string), dbClub.Description)
					eaa.Assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
					eaa.Assert.Equal(models.RecruitmentCycle((*body)["recruitment_cycle"].(string)), dbClub.RecruitmentCycle)
					eaa.Assert.Equal(models.RecruitmentType((*body)["recruitment_type"].(string)), dbClub.RecruitmentType)
					eaa.Assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
					eaa.Assert.Equal((*body)["logo"].(string), dbClub.Logo)
				},
			},
		)
	}
	appAssert.Close()
}

func TestInsertPOCFailsOnMissingFields(t *testing.T) {
	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
	appAssert, _ := CreateSamplePOC(t, existingAppAssert)

	for _, missingField := range []string{
		"name",
		"email",
		"position",
	} {
		samplePOCPermutation := *SamplePOCFactory()
		delete(samplePOCPermutation, missingField)
		TestRequest{
			Method: fiber.MethodPut,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
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
	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
	appAssert, _ := CreateSamplePOC(t, existingAppAssert)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
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
	clubId := uuid.New()

	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/", clubId),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404,
				Message: errors.ClubNotFound,
			},
			DBTester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club
				err := app.Conn.Where("id = ?", clubId).First(&club).Error
				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestGetPOCWorks(t *testing.T) {
	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
	appAssert, pocUUID := CreateSamplePOC(t, existingAppAssert)

	h.TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, pocUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
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
	_, _, clubUUID := CreateSampleClub(h.InitTest(t))

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
			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: errors.FailedToValidatePointOfContactId,
			},
		).Close()
	}
}

func TestGetPOCFailsNotExist(t *testing.T) {
	_, _, clubUUID := CreateSampleClub(h.InitTest(t))
	pocUUIDNotExist := uuid.New()

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, pocUUIDNotExist),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404,
				Message: errors.PointOfContactNotFound,
			},
			DBTester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var pointOfContact models.PointOfContact
				err := app.Conn.Where("id = ?", pocUUIDNotExist).First(&pointOfContact).Error
				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

// DELETE TEST CASES
func TestDeletePointOfContactWorks(t *testing.T) {
	existingAppAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
	appAssert, pocUUID := CreateSamplePOC(t, existingAppAssert)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, pocUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   204,
			DBTester: TestNumPOCRemainsAt0,
		},
	).Close()
}

func TestDeletePOCClubIDBadRequest(t *testing.T) {
	existingAppAssert, _, _ := CreateSampleClub(h.InitTest(t))
	_, pocUUID := CreateSamplePOC(t, existingAppAssert)
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
			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", badRequest, pocUUID),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: errors.FailedToValidateClub,
			},
		).Close()
	}
}

func TestDeletePOCBadRequest(t *testing.T) {
	_, _, clubUUID := CreateSampleClub(h.InitTest(t))

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
			Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: errors.FailedToValidatePointOfContactId,
			},
		).Close()
	}
}

func TestDeletePOCClubNotExist(t *testing.T) {
	existingAppAssert, _, _ := CreateSampleClub(h.InitTest(t))
	_, pocUUID := CreateSamplePOC(t, existingAppAssert)
	clubUUID := uuid.New()

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, pocUUID),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404,
				Message: errors.ClubNotFound,
			},
			DBTester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club

				err := app.Conn.Where("id = ?", clubUUID).First(&club).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeletePOCNotExist(t *testing.T) {
	_, _, clubUUID := CreateSampleClub(h.InitTest(t))
	pocUUID := uuid.New()

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s/poc/%s", clubUUID, pocUUID),
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  404,
				Message: errors.PointOfContactNotFound,
			},
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var pointOfContact models.PointOfContact

				err := app.Conn.Where("id = ?", pocUUID).First(&pointOfContact).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

// assert remaining numbers of POC
func AssertNumPOCRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var pointOfContact []models.PointOfContact

	err := eaa.App.Conn.Find(&pointOfContact).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(pointOfContact))
}

// assert remaining POC = 1
var TestNumPOCRemainsAt1 = func(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumPOCRemainsAtN(eaa, resp, 1)
}

var TestNumPOCRemainsAt0 = func(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumPOCRemainsAtN(eaa, resp, 0)
}

func TestUpdateClubFailsBadRequest(t *testing.T) {
	appAssert := h.InitTest(t)
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}
	sampleStudent, rawPassword := h.SampleStudentFactory()

	for _, badRequest := range badRequests {
		appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodPatch,
				Path:   fmt.Sprintf("/api/v1/clubs/%s", badRequest),
				Body:   h.SampleStudentJSONFactory(sampleStudent, rawPassword),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}
	appAssert.Close()
}

func TestUpdateClubFailsOnClubIdNotExist(t *testing.T) {
	uuid := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(h.TestRequest{
		Method:             fiber.MethodPatch,
		Path:               fmt.Sprintf("/api/v1/clubs/%s", uuid),
		Body:               SampleClubFactory(nil),
		Role:               &models.Super,
		TestUserIDReplaces: h.StringToPointer("user_id"),
	},
		h.ErrorWithTester{
			Error: errors.ClubNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var club models.Club

				err := eaa.App.Conn.Where("id = ?", uuid).First(&club).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeleteClubWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/clubs/%s", clubUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: TestNumClubsRemainsAt1,
		},
	).Close()
}

func TestDeleteClubNotExist(t *testing.T) {
	uuid := uuid.New()
	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/clubs/%s", uuid),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.ClubNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var club models.Club

				err := eaa.App.Conn.Where("id = ?", uuid).First(&club).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

				AssertNumClubsRemainsAtN(eaa, resp, 1)
			},
		},
	).Close()
}

func TestDeleteClubBadRequest(t *testing.T) {
	appAssert := h.InitTest(t)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"hello",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodDelete,
				Path:   fmt.Sprintf("/api/v1/clubs/%s", badRequest),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}
	appAssert.Close()
}

func TestUpdateClubFailsBadRequest(t *testing.T) {
	appAssert := h.InitTest(t)
	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}
	sampleStudent, rawPassword := h.SampleStudentFactory()

	for _, badRequest := range badRequests {
		appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodPatch,
				Path:   fmt.Sprintf("/api/v1/clubs/%s", badRequest),
				Body:   h.SampleStudentJSONFactory(sampleStudent, rawPassword),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}
	appAssert.Close()
}

func TestUpdateClubFailsOnClubIdNotExist(t *testing.T) {
	uuid := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(h.TestRequest{
		Method:             fiber.MethodPatch,
		Path:               fmt.Sprintf("/api/v1/clubs/%s", uuid),
		Body:               SampleClubFactory(nil),
		Role:               &models.Super,
		TestUserIDReplaces: h.StringToPointer("user_id"),
	},
		h.ErrorWithTester{
			Error: errors.ClubNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var club models.Club

				err := eaa.App.Conn.Where("id = ?", uuid).First(&club).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeleteClubWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/clubs/%s", clubUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: TestNumClubsRemainsAt1,
		},
	).Close()
}

func TestDeleteClubNotExist(t *testing.T) {
	uuid := uuid.New()
	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/clubs/%s", uuid),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.ClubNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var club models.Club

				err := eaa.App.Conn.Where("id = ?", uuid).First(&club).Error

				eaa.Assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

				AssertNumClubsRemainsAtN(eaa, resp, 1)
			},
		},
	).Close()
}

func TestDeleteClubBadRequest(t *testing.T) {
	appAssert := h.InitTest(t)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"hello",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodDelete,
				Path:   fmt.Sprintf("/api/v1/clubs/%s", badRequest),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}
	appAssert.Close()
}
