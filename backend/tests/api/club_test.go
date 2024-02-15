package tests

import (
	stdliberrors "errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

	return newAppAssert, newAppAssert.App.TestUser.UUID, sampleClubUUID
}

func TestCreateClubWorks(t *testing.T) {
	t.Parallel()
	existingAppAssert, _, _ := CreateSampleClub(h.InitTest(t))
	existingAppAssert.Close()
}

func TestGetClubsWorks(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	AssertCreateBadClubDataFails(t,
		"application_link",
		[]interface{}{
			"Not an URL",
			"@#139081#$Ad_Axf",
		},
	)
}

func TestCreateClubFailsOnInvalidLogo(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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

func TestUpdateClubFailsBadRequest(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
