package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AssertClubTagsRespDB(eaa h.ExistingAppAssert, resp *http.Response, id uuid.UUID) {
	var respTags []models.Tag

	err := json.NewDecoder(resp.Body).Decode(&respTags)

	eaa.Assert.NilError(err)

	var dbClub models.Club

	err = eaa.App.Conn.First(&dbClub, id).Error

	eaa.Assert.NilError(err)

	var dbTags []models.Tag

	err = eaa.App.Conn.Model(&dbClub).Association("Tag").Find(&dbTags)

	eaa.Assert.NilError(err)

	for i, respTag := range respTags {
		eaa.Assert.Equal(respTag.ID, dbTags[i].ID)
		eaa.Assert.Equal(respTag.Name, dbTags[i].Name)
		eaa.Assert.Equal(respTag.CategoryID, dbTags[i].CategoryID)
	}
}

func AssertSampleClubTagsRespDB(eaa h.ExistingAppAssert, resp *http.Response, uuid uuid.UUID) {
	AssertClubTagsRespDB(eaa, resp, uuid)
}

func TestCreateClubTagsFailsOnInvalidDataType(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	invalidTags := []interface{}{
		[]string{"1", "2", "34"},
		[]models.Tag{{Name: "Test", CategoryID: uuid.UUID{}}, {Name: "Test2", CategoryID: uuid.UUID{}}},
		[]float32{1.32, 23.5, 35.1},
	}

	for _, tag := range invalidTags {
		malformedTag := *SampleTagIDsFactory(nil)
		malformedTag["tags"] = tag

		appAssert = appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", clubUUID),
				Body:   &malformedTag,
				Role:   &models.Super,
			},
			errors.FailedToParseRequestBody,
		)
	}

	appAssert.Close()
}

func TestCreateClubTagsFailsOnInvalidUserID(t *testing.T) {
	appAssert := h.InitTest(t)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert = appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   fmt.Sprintf("/api/v1/clubs/%s/tags", badRequest),
				Body:   SampleTagIDsFactory(nil),
				Role:   &models.Student,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestCreateClubTagsFailsOnInvalidKey(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	invalidBody := []map[string]interface{}{
		{
			"tag": UUIDSlice{uuid.New(), uuid.New()},
		},
		{
			"tagIDs": []uint{1, 2, 3},
		},
	}

	for _, body := range invalidBody {
		body := body
		appAssert = appAssert.TestOnError(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", clubUUID),
				Body:   &body,
				Role:   &models.Student,
			},
			errors.FailedToValidateClubTags,
		)
	}

	appAssert.Close()
}

func TestCreateClubTagsFailsOnNonExistentClub(t *testing.T) {
	uuid := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", uuid),
			Body:   SampleTagIDsFactory(nil),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.ClubNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var dbClub models.Club

				err := eaa.App.Conn.First(&dbClub, uuid).Error

				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestCreateClubTagsWorks(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))
	tagUUIDs, appAssert := CreateSetOfTags(appAssert)

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", clubUUID),
			Body:   SampleTagIDsFactory(&tagUUIDs),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertSampleClubTagsRespDB(eaa, resp, clubUUID)
			},
		},
	)

	appAssert.Close()
}

func TestCreateClubTagsNoneAddedIfInvalid(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", clubUUID),
			Body:   SampleTagIDsFactory(nil),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var respTags []models.Tag

				err := json.NewDecoder(resp.Body).Decode(&respTags)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(respTags), 0)
			},
		},
	).Close()
}

func TestGetClubTagsFailsOnNonExistentClub(t *testing.T) {
	uuid := uuid.New()

	h.InitTest(t).TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", uuid),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error: errors.ClubNotFound,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var dbClub models.Club

				err := eaa.App.Conn.First(&dbClub, uuid).Error

				eaa.Assert.Assert(err != nil)
			},
		},
	).Close()
}

func TestGetClubTagsReturnsEmptyListWhenNoneAdded(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", clubUUID),
			Role:   &models.Student,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				var respTags []models.Tag

				err := json.NewDecoder(resp.Body).Decode(&respTags)

				eaa.Assert.NilError(err)

				eaa.Assert.Equal(len(respTags), 0)
			},
		},
	).Close()
}

func TestGetClubTagsReturnsCorrectList(t *testing.T) {
	appAssert, _, clubUUID := CreateSampleClub(h.InitTest(t))

	tagUUIDs, appAssert := CreateSetOfTags(appAssert)

	appAssert.TestOnStatus(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", clubUUID),
			Body:   SampleTagIDsFactory(&tagUUIDs),
			Role:   &models.Student,
		},
		fiber.StatusCreated,
	).TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/clubs/%s/tags/", clubUUID),
			Role:   &models.Student,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertSampleClubTagsRespDB(eaa, resp, clubUUID)
			},
		},
	).Close()
}
