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
	"github.com/huandu/go-assert"

	"github.com/goccy/go-json"
)

func SampleTagFactory(categoryID uuid.UUID) *map[string]interface{} {
	return &map[string]interface{}{
		"name":        "Generate",
		"category_id": categoryID,
	}
}

func AssertTagWithBodyRespDB(app h.TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
	var respTag models.Tag

	err := json.NewDecoder(resp.Body).Decode(&respTag)

	assert.NilError(err)

	var dbTag models.Tag

	err = app.Conn.First(&dbTag).Error

	assert.NilError(err)

	assert.Equal(dbTag.ID, respTag.ID)
	assert.Equal(dbTag.Name, respTag.Name)
	assert.Equal(dbTag.CategoryID, respTag.CategoryID)

	assert.Equal((*body)["name"].(string), dbTag.Name)
	assert.Equal((*body)["category_id"].(uuid.UUID), dbTag.CategoryID)

	return dbTag.ID
}

func AssertSampleTagBodyRespDB(t *testing.T, app h.TestApp, assert *assert.A, resp *http.Response) uuid.UUID {
	appAssert, uuid := CreateSampleCategory(t,
		&h.ExistingAppAssert{
			App:    app,
			Assert: assert,
		})
	return AssertTagWithBodyRespDB(appAssert.App, appAssert.Assert, resp, SampleTagFactory(uuid))
}

func CreateSampleTag(t *testing.T) (appAssert h.ExistingAppAssert, categoryUUID uuid.UUID, tagUUID uuid.UUID) {
	appAssert, categoryUUID = CreateSampleCategory(t, nil)

	AssertSampleTagBodyRespDB := func(app h.TestApp, assert *assert.A, resp *http.Response) {
		tagUUID = AssertTagWithBodyRespDB(app, assert, resp, SampleTagFactory(categoryUUID))
	}

	appAssert.TestOnStatusAndDB(
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
	appAssert, _, _ := CreateSampleTag(t)
	appAssert.Close()
}

func AssertNumTagsRemainsAtN(app h.TestApp, assert *assert.A, resp *http.Response, n int) {
	var tags []models.Tag

	err := app.Conn.Find(&tags).Error

	assert.NilError(err)

	assert.Equal(n, len(tags))
}

func AssertNoTags(app h.TestApp, assert *assert.A, resp *http.Response) {
	AssertNumTagsRemainsAtN(app, assert, resp, 0)
}

func Assert1Tag(app h.TestApp, assert *assert.A, resp *http.Response) {
	AssertNumTagsRemainsAtN(app, assert, resp, 1)
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
		appAssert.TestOnErrorAndDB(
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
		appAssert.TestOnErrorAndDB(
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
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(t)

	existingAppAssert.TestOnStatusAndDB(
		h.TestRequest{
			Method: fiber.MethodGet,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				AssertTagWithBodyRespDB(app, assert, resp, SampleTagFactory(categoryUUID))
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
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(t)

	generateNUTag := *SampleTagFactory(categoryUUID)
	generateNUTag["name"] = "GenerateNU"

	AssertUpdatedTagBodyRespDB := func(app h.TestApp, assert *assert.A, resp *http.Response) {
		tagUUID = AssertTagWithBodyRespDB(app, assert, resp, &generateNUTag)
	}

	existingAppAssert.TestOnStatusAndDB(
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
	existingAppAssert, _, tagUUID := CreateSampleTag(t)

	technologyCategory := *SampleCategoryFactory()
	technologyCategory["name"] = "Technology"

	var technologyCategoryUUID uuid.UUID

	AssertNewCategoryBodyRespDB := func(app h.TestApp, assert *assert.A, resp *http.Response) {
		technologyCategoryUUID = AssertCategoryWithBodyRespDBMostRecent(app, assert, resp, &technologyCategory)
	}

	existingAppAssert.TestOnStatusAndDB(
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

	AssertUpdatedTagBodyRespDB := func(app h.TestApp, assert *assert.A, resp *http.Response) {
		AssertTagWithBodyRespDB(app, assert, resp, &technologyTag)
	}

	existingAppAssert.TestOnStatusAndDB(
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
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(t)

	existingAppAssert.TestOnStatusAndDB(
		h.TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
			Body:   SampleTagFactory(categoryUUID),
			Role:   &models.Super,
		},
		h.TesterWithStatus{
			Status: fiber.StatusOK,
			Tester: func(app h.TestApp, assert *assert.A, resp *http.Response) {
				AssertTagWithBodyRespDB(app, assert, resp, SampleTagFactory(categoryUUID))
			},
		},
	).Close()
}

func TestUpdateTagFailsBadRequest(t *testing.T) {
	appAssert, uuid := CreateSampleCategory(t, nil)

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
	existingAppAssert, _, tagUUID := CreateSampleTag(t)

	existingAppAssert.TestOnStatusAndDB(
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
	appAssert, _, _ := CreateSampleTag(t)

	badRequests := []string{
		"0",
		"-1",
		"1.1",
		"foo",
		"null",
	}

	for _, badRequest := range badRequests {
		appAssert.TestOnErrorAndDB(
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
	appAssert, _, _ := CreateSampleTag(t)

	appAssert.TestOnErrorAndDB(
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
