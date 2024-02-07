package tests

import (
	stdliberrors "errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/huandu/go-assert"
	"gorm.io/gorm"
)

func SampleClubFactory(userID uuid.UUID) *map[string]interface{} {
	return &map[string]interface{}{
		"user_id":           userID,
		"name":              "Generate",
		"preview":           "Generate is Northeastern's premier student-led product development studio.",
		"description":       "https://mongodb.com",
		"num_members":       1,
		"is_recruiting":     true,
		"recruitment_cycle": "always",
		"recruitment_type":  "application",
		"application_link":  "https://generatenu.com/apply",
		"logo":              "https://aws.amazon.com/s3/",
	}
}

func AssertClubBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respClub models.Club

	err := json.NewDecoder(resp.Body).Decode(&respClub)

	assert.NilError(err)

	var dbClubs []models.Club

	err = app.Conn.Order("created_at desc").Find(&dbClubs).Error

	assert.NilError(err)

	assert.Equal(2, len(dbClubs))

	dbClub := dbClubs[0]

	assert.Equal(dbClub.ID, respClub.ID)
	assert.Equal(dbClub.Name, respClub.Name)
	assert.Equal(dbClub.Preview, respClub.Preview)
	assert.Equal(dbClub.Description, respClub.Description)
	assert.Equal(dbClub.NumMembers, respClub.NumMembers)
	assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
	assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
	assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
	assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
	assert.Equal(dbClub.Logo, respClub.Logo)

	var dbAdmins []models.User

	err = app.Conn.Model(&dbClub).Association("Admin").Find(&dbAdmins)

	assert.NilError(err)

	assert.Equal(1, len(dbAdmins))

	assert.Equal((*body)["user_id"].(uuid.UUID), dbAdmins[0].ID)
	assert.Equal((*body)["name"].(string), dbClub.Name)
	assert.Equal((*body)["preview"].(string), dbClub.Preview)
	assert.Equal((*body)["description"].(string), dbClub.Description)
	assert.Equal((*body)["num_members"].(int), dbClub.NumMembers)
	assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
	assert.Equal(models.RecruitmentCycle((*body)["recruitment_cycle"].(string)), dbClub.RecruitmentCycle)
	assert.Equal(models.RecruitmentType((*body)["recruitment_type"].(string)), dbClub.RecruitmentType)
	assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
	assert.Equal((*body)["logo"].(string), dbClub.Logo)

	return dbClub.ID
}

func AssertClubWithBodyRespDBMostRecent(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respClub models.Club

	err := json.NewDecoder(resp.Body).Decode(&respClub)

	assert.NilError(err)

	var dbClub models.Club

	err = app.Conn.Order("created_at desc").First(&dbClub).Error

	assert.NilError(err)

	assert.Equal(dbClub.ID, respClub.ID)
	assert.Equal(dbClub.Name, respClub.Name)
	assert.Equal(dbClub.Preview, respClub.Preview)
	assert.Equal(dbClub.Description, respClub.Description)
	assert.Equal(dbClub.NumMembers, respClub.NumMembers)
	assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
	assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
	assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
	assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
	assert.Equal(dbClub.Logo, respClub.Logo)

	var dbAdmins []models.User

	err = app.Conn.Model(&dbClub).Association("Admins").Find(&dbAdmins)

	assert.NilError(err)

	assert.Equal(1, len(dbAdmins))

	dbAdmin := dbAdmins[0]

	assert.Equal((*body)["user_id"].(uuid.UUID), dbAdmin.ID)
	assert.Equal((*body)["name"].(string), dbClub.Name)
	assert.Equal((*body)["preview"].(string), dbClub.Preview)
	assert.Equal((*body)["description"].(string), dbClub.Description)
	assert.Equal((*body)["num_members"].(int), dbClub.NumMembers)
	assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
	assert.Equal((*body)["recruitment_cycle"].(string), dbClub.RecruitmentCycle)
	assert.Equal((*body)["recruitment_type"].(string), dbClub.RecruitmentType)
	assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
	assert.Equal((*body)["logo"].(string), dbClub.Logo)

	return dbClub.ID
}

func AssertSampleClubBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, userID uuid.UUID) uuid.UUID {
	return AssertClubBodyRespDB(app, assert, resp, SampleClubFactory(userID))
}

func CreateSampleClub(t *testing.T, existingAppAssert *ExistingAppAssert) (eaa ExistingAppAssert, userUUID uuid.UUID, clubUUID uuid.UUID) {
	appAssert, userID := CreateSampleUser(t, existingAppAssert)

	var sampleClubUUID uuid.UUID

	newAppAssert := TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/clubs/",
		Body:   SampleClubFactory(userID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusCreated,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				sampleClubUUID = AssertSampleClubBodyRespDB(app, assert, resp, userID)
			},
		},
	)

	if existingAppAssert == nil {
		return newAppAssert, userID, sampleClubUUID
	} else {
		return *existingAppAssert, userID, sampleClubUUID
	}
}

func TestCreateClubWorks(t *testing.T) {
	existingAppAssert, _, _ := CreateSampleClub(t, nil)
	existingAppAssert.Close()
}

func TestGetClubsWorks(t *testing.T) {
	TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/clubs/",
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respClubs []models.Club

				err := json.NewDecoder(resp.Body).Decode(&respClubs)

				assert.NilError(err)

				assert.Equal(1, len(respClubs))

				respClub := respClubs[0]

				var dbClubs []models.Club

				err = app.Conn.Order("created_at desc").Find(&dbClubs).Error

				assert.NilError(err)

				assert.Equal(1, len(dbClubs))

				dbClub := dbClubs[0]

				assert.Equal(dbClub.ID, respClub.ID)
				assert.Equal(dbClub.Name, respClub.Name)
				assert.Equal(dbClub.Preview, respClub.Preview)
				assert.Equal(dbClub.Description, respClub.Description)
				assert.Equal(dbClub.NumMembers, respClub.NumMembers)
				assert.Equal(dbClub.IsRecruiting, respClub.IsRecruiting)
				assert.Equal(dbClub.RecruitmentCycle, respClub.RecruitmentCycle)
				assert.Equal(dbClub.RecruitmentType, respClub.RecruitmentType)
				assert.Equal(dbClub.ApplicationLink, respClub.ApplicationLink)
				assert.Equal(dbClub.Logo, respClub.Logo)

				assert.Equal("SAC", dbClub.Name)
				assert.Equal("SAC", dbClub.Preview)
				assert.Equal("SAC", dbClub.Description)
				assert.Equal(1, dbClub.NumMembers)
				assert.Equal(true, dbClub.IsRecruiting)
				assert.Equal(models.RecruitmentCycle(models.Always), dbClub.RecruitmentCycle)
				assert.Equal(models.RecruitmentType(models.Application), dbClub.RecruitmentType)
				assert.Equal("https://generatenu.com/apply", dbClub.ApplicationLink)
				assert.Equal("https://aws.amazon.com/s3", dbClub.Logo)
			},
		},
	).Close()
}

func AssertNumClubsRemainsAtN(app TestApp, assert *assert.A, resp *http.Response, n int) {
	var dbClubs []models.Club

	err := app.Conn.Order("created_at desc").Find(&dbClubs).Error

	assert.NilError(err)

	assert.Equal(n, len(dbClubs))
}

var TestNumClubsRemainsAt1 = func(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumClubsRemainsAtN(app, assert, resp, 1)
}

func AssertCreateBadClubDataFails(t *testing.T, jsonKey string, badValues []interface{}) {
	appAssert, uuid := CreateSampleUser(t, nil)

	for _, badValue := range badValues {
		sampleClubPermutation := *SampleClubFactory(uuid)
		sampleClubPermutation[jsonKey] = badValue

		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/clubs/",
			Body:   &sampleClubPermutation,
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error:    errors.FailedToValidateClub,
				DBTester: TestNumClubsRemainsAt1,
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
			//"https://google.com", <-- TODO uncomment once we figure out s3 url validation
		},
	)
}

func TestUpdateClubWorks(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)

	updatedClub := SampleClubFactory(userUUID)
	(*updatedClub)["name"] = "Updated Name"
	(*updatedClub)["preview"] = "Updated Preview"

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/clubs/%s", clubUUID),
		Body:   updatedClub,
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertClubBodyRespDB(app, assert, resp, updatedClub)
			},
		},
	).Close()
}

func TestUpdateClubFailsOnInvalidBody(t *testing.T) {
	appAssert, userUUID, clubUUID := CreateSampleClub(t, nil)

	body := SampleClubFactory(userUUID)

	for _, invalidData := range []map[string]interface{}{
		{"description": "Not a URL"},
		{"recruitment_cycle": "1234"},
		{"recruitment_type": "ALLLLWAYSSSS"},
		{"application_link": "Not an URL"},
		{"logo": "@12394X_2"},
	} {
		TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/clubs/%s", clubUUID),
			Body:   &invalidData,
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error: errors.FailedToValidateClub,
				DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
					var dbClubs []models.Club

					err := app.Conn.Order("created_at desc").Find(&dbClubs).Error

					assert.NilError(err)

					assert.Equal(2, len(dbClubs))

					dbClub := dbClubs[0]

					var dbAdmins []models.User

					err = app.Conn.Model(&dbClub).Association("Admin").Find(&dbAdmins)

					assert.NilError(err)

					assert.Equal(1, len(dbAdmins))

					assert.Equal((*body)["user_id"].(uuid.UUID), dbAdmins[0].ID)
					assert.Equal((*body)["name"].(string), dbClub.Name)
					assert.Equal((*body)["preview"].(string), dbClub.Preview)
					assert.Equal((*body)["description"].(string), dbClub.Description)
					assert.Equal((*body)["num_members"].(int), dbClub.NumMembers)
					assert.Equal((*body)["is_recruiting"].(bool), dbClub.IsRecruiting)
					assert.Equal(models.RecruitmentCycle((*body)["recruitment_cycle"].(string)), dbClub.RecruitmentCycle)
					assert.Equal(models.RecruitmentType((*body)["recruitment_type"].(string)), dbClub.RecruitmentType)
					assert.Equal((*body)["application_link"].(string), dbClub.ApplicationLink)
					assert.Equal((*body)["logo"].(string), dbClub.Logo)
				},
			},
		)
	}
	appAssert.Close()
}

func TestUpdateClubFailsBadRequest(t *testing.T) {
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
			Path:   fmt.Sprintf("/api/v1/clubs/%s", badRequest),
			Body:   SampleUserFactory(),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestUpdateClubFailsOnClubIdNotExist(t *testing.T) {
	appAssert, userUUID := CreateSampleUser(t, nil)

	uuid := uuid.New()

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/clubs/%s", uuid),
		Body:   SampleClubFactory(userUUID),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error: errors.ClubNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club

				err := app.Conn.Where("id = ?", uuid).First(&club).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))
			},
		},
	).Close()
}

func TestDeleteClubWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(t, nil)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s", clubUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusNoContent,
			DBTester: TestNumClubsRemainsAt1,
		},
	).Close()
}

func TestDeleteClubNotExist(t *testing.T) {
	uuid := uuid.New()
	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/clubs/%s", uuid),
	}.TestOnErrorAndDB(t, nil,
		ErrorWithDBTester{
			Error: errors.ClubNotFound,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var club models.Club

				err := app.Conn.Where("id = ?", uuid).First(&club).Error

				assert.Assert(stdliberrors.Is(err, gorm.ErrRecordNotFound))

				AssertNumClubsRemainsAtN(app, assert, resp, 1)
			},
		},
	).Close()
}

func TestDeleteClubBadRequest(t *testing.T) {
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
			Path:   fmt.Sprintf("/api/v1/clubs/%s", badRequest),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}

}

