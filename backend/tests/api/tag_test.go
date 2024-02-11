package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	h "github.com/GenerateNU/sac/backend/tests/api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/goccy/go-json"
)

func SampleTagFactory(categoryID uuid.UUID) *map[string]interface{} {
	return &map[string]interface{}{
		"name":        "Generate",
		"category_id": categoryID,
	}
}

func AssertTagWithBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respTag models.Tag

	err := json.NewDecoder(resp.Body).Decode(&respTag)

	eaa.Assert.NilError(err)

	var dbTag models.Tag

	err = eaa.App.Conn.First(&dbTag).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(dbTag.ID, respTag.ID)
	eaa.Assert.Equal(dbTag.Name, respTag.Name)
	eaa.Assert.Equal(dbTag.CategoryID, respTag.CategoryID)

	eaa.Assert.Equal((*body)["name"].(string), dbTag.Name)
	eaa.Assert.Equal((*body)["category_id"].(uuid.UUID), dbTag.CategoryID)

	return dbTag.ID
}

func AssertSampleTagBodyRespDB(t *testing.T, eaa h.ExistingAppAssert, resp *http.Response) uuid.UUID {
	appAssert, uuid := CreateSampleCategory(eaa)
	return AssertTagWithBodyRespDB(appAssert, resp, SampleTagFactory(uuid))
}

func CreateSampleTag(appAssert h.ExistingAppAssert) (existingAppAssert h.ExistingAppAssert, categoryUUID uuid.UUID, tagUUID uuid.UUID) {
	appAssert, categoryUUID = CreateSampleCategory(appAssert)

	AssertSampleTagBodyRespDB := func(eaa h.ExistingAppAssert, resp *http.Response) {
		tagUUID = AssertTagWithBodyRespDB(appAssert, resp, SampleTagFactory(categoryUUID))
	}

	appAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/tags/",
			Body:   SampleTagFactory(categoryUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: AssertSampleTagBodyRespDB,
		},
	)

	return appAssert, categoryUUID, tagUUID
}

func TestCreateTagWorks(t *testing.T) {
	appAssert, _, _ := CreateSampleTag(h.InitTest(t))
	appAssert.Close()
}

func AssertNumTagsRemainsAtN(eaa h.ExistingAppAssert, resp *http.Response, n int) {
	var tags []models.Tag

	err := eaa.App.Conn.Find(&tags).Error

	eaa.Assert.NilError(err)

	eaa.Assert.Equal(n, len(tags))
}

func AssertNoTags(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumTagsRemainsAtN(eaa, resp, 0)
}

func Assert1Tag(eaa h.ExistingAppAssert, resp *http.Response) {
	AssertNumTagsRemainsAtN(eaa, resp, 1)
}

func TestCreateTagFailsBadRequest(t *testing.T) {
	appAssert := h.InitTest(t)

	badBodys := []map[string]interface{}{
		{
			"name":        "Generate",
			"category_id": "1",
		},
		{
			"name":        1,
			"category_id": 1,
		},
	}

	for _, badBody := range badBodys {
		badBody := badBody
		appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/tags/",
				Body:   &badBody,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToParseRequestBody,
				Tester: AssertNoTags,
			},
		)
	}

	appAssert.Close()
}

func TestCreateTagFailsValidation(t *testing.T) {
	appAssert := h.InitTest(t)

	badBodys := []map[string]interface{}{
		{
			"name": "Generate",
		},
		{
			"category_id": uuid.New(),
		},
		{},
	}

	for _, badBody := range badBodys {
		badBody := badBody
		appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodPost,
				Path:   "/api/v1/tags/",
				Body:   &badBody,
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToValidateTag,
				Tester: AssertNoTags,
			},
		)
	}

	appAssert.Close()
}

func TestGetTagWorks(t *testing.T) {
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(h.InitTest(t))

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertTagWithBodyRespDB(eaa, resp, SampleTagFactory(categoryUUID))
			},
		},
	).Close()
}

func TestGetTagFailsBadRequest(t *testing.T) {
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
				Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestGetTagFailsNotFound(t *testing.T) {
	h.InitTest(t).TestOnError(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/tags/%s", uuid.New()),
			Role:   &models.Super,
		},
		errors.TagNotFound,
	).Close()
}

func TestUpdateTagWorksUpdateName(t *testing.T) {
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(h.InitTest(t))

	generateNUTag := *SampleTagFactory(categoryUUID)
	generateNUTag["name"] = "GenerateNU"

	AssertUpdatedTagBodyRespDB := func(eaa h.ExistingAppAssert, resp *http.Response) {
		tagUUID = AssertTagWithBodyRespDB(eaa, resp, &generateNUTag)
	}

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Body:   &generateNUTag,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: AssertUpdatedTagBodyRespDB,
		},
	).Close()
}

func TestUpdateTagWorksUpdateCategory(t *testing.T) {
	existingAppAssert, _, tagUUID := CreateSampleTag(h.InitTest(t))

	technologyCategory := *SampleCategoryFactory()
	technologyCategory["name"] = "Technology"

	var technologyCategoryUUID uuid.UUID

	AssertNewCategoryBodyRespDB := func(eaa h.ExistingAppAssert, resp *http.Response) {
		technologyCategoryUUID = AssertCategoryWithBodyRespDBMostRecent(eaa, resp, &technologyCategory)
	}

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body:   &technologyCategory,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusCreated,
			Tester: AssertNewCategoryBodyRespDB,
		},
	)

	technologyTag := *SampleTagFactory(technologyCategoryUUID)

	AssertUpdatedTagBodyRespDB := func(eaa h.ExistingAppAssert, resp *http.Response) {
		AssertTagWithBodyRespDB(eaa, resp, &technologyTag)
	}

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Body:   &technologyTag,
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: AssertUpdatedTagBodyRespDB,
		},
	).Close()
}

func TestUpdateTagWorksWithSameDetails(t *testing.T) {
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(h.InitTest(t))

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Body:   SampleTagFactory(categoryUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
				AssertTagWithBodyRespDB(eaa, resp, SampleTagFactory(categoryUUID))
			},
		},
	).Close()
}

func TestUpdateTagFailsBadRequest(t *testing.T) {
	appAssert, uuid := CreateSampleCategory(h.InitTest(t))

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
				Method: fiber.MethodPatch,
				Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
				Body:   SampleTagFactory(uuid),
				Role:   &models.Super,
			},
			errors.FailedToValidateID,
		)
	}

	appAssert.Close()
}

func TestDeleteTagWorks(t *testing.T) {
	existingAppAssert, _, tagUUID := CreateSampleTag(h.InitTest(t))

	existingAppAssert.TestOnStatusAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusNoContent,
			Tester: AssertNoTags,
		},
	).Close()
}

func TestDeleteTagFailsBadRequest(t *testing.T) {
	appAssert, _, _ := CreateSampleTag(h.InitTest(t))

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert.TestOnErrorAndTester(
			h.TestRequest{
				Method: fiber.MethodDelete,
				Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
				Role:   &models.Super,
			},
			h.ErrorWithTester{
				Error:  errors.FailedToValidateID,
				Tester: Assert1Tag,
			},
		)
	}

	appAssert.Close()
}

func TestDeleteTagFailsNotFound(t *testing.T) {
	appAssert, _, _ := CreateSampleTag(h.InitTest(t))

	appAssert.TestOnErrorAndTester(
		h.TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/tags/%s", uuid.New()),
			Role:   &models.Super,
		},
		h.ErrorWithTester{
			Error:  errors.TagNotFound,
			Tester: Assert1Tag,
		},
	).Close()
}
