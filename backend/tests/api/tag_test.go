package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
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

func AssertTagWithBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *map[string]interface{}) uuid.UUID {
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

func AssertSampleTagBodyRespDB(t *testing.T, app TestApp, assert *assert.A, resp *http.Response) uuid.UUID {
	appAssert, uuid := CreateSampleCategory(t, &ExistingAppAssert{App: app,
		Assert: assert})
	return AssertTagWithBodyRespDB(appAssert.App, appAssert.Assert, resp, SampleTagFactory(uuid))
}

func CreateSampleTag(t *testing.T) (appAssert ExistingAppAssert, categoryUUID uuid.UUID, tagUUID uuid.UUID) {
	appAssert, categoryUUID = CreateSampleCategory(t, nil)

	var AssertSampleTagBodyRespDB = func(app TestApp, assert *assert.A, resp *http.Response) {
		tagUUID = AssertTagWithBodyRespDB(app, assert, resp, SampleTagFactory(categoryUUID))
	}

	TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/tags/",
		Body:   SampleTagFactory(categoryUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusCreated,
			DBTester: AssertSampleTagBodyRespDB,
		},
	)

	return appAssert, categoryUUID, tagUUID
}

func TestCreateTagWorks(t *testing.T) {
	appAssert, _, _ := CreateSampleTag(t)
	appAssert.Close()
}

func AssertNumTagsRemainsAtN(app TestApp, assert *assert.A, resp *http.Response, n int) {
	var tags []models.Tag

	err := app.Conn.Find(&tags).Error

	assert.NilError(err)

	assert.Equal(n, len(tags))
}

func AssertNoTags(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumTagsRemainsAtN(app, assert, resp, 0)
}

func Assert1Tag(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumTagsRemainsAtN(app, assert, resp, 1)
}

func TestCreateTagFailsBadRequest(t *testing.T) {
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
		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/tags/",
			Body:   &badBody,
		}.TestOnErrorAndDB(t, nil,
			ErrorWithDBTester{
				Error:    errors.FailedToParseRequestBody,
				DBTester: AssertNoTags,
			},
		).Close()
	}
}

func TestCreateTagFailsValidation(t *testing.T) {
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
		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/tags/",
			Body:   &badBody,
		}.TestOnErrorAndDB(t, nil,
			ErrorWithDBTester{
				Error:    errors.FailedToValidateTag,
				DBTester: AssertNoTags,
			},
		).Close()
	}
}

func TestGetTagWorks(t *testing.T) {
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(t)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertTagWithBodyRespDB(app, assert, resp, SampleTagFactory(categoryUUID))
			},
		},
	).Close()
}

func TestGetTagFailsBadRequest(t *testing.T) {
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
			Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestGetTagFailsNotFound(t *testing.T) {
	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/tags/%s", uuid.New()),
	}.TestOnError(t, nil, errors.TagNotFound).Close()
}

func TestUpdateTagWorksUpdateName(t *testing.T) {
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(t)

	generateNUTag := *SampleTagFactory(categoryUUID)
	generateNUTag["name"] = "GenerateNU"

	var AssertUpdatedTagBodyRespDB = func(app TestApp, assert *assert.A, resp *http.Response) {
		tagUUID = AssertTagWithBodyRespDB(app, assert, resp, &generateNUTag)
	}

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
		Body:   &generateNUTag,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: AssertUpdatedTagBodyRespDB,
		},
	).Close()
}

func TestUpdateTagWorksUpdateCategory(t *testing.T) {
	existingAppAssert, _, tagUUID := CreateSampleTag(t)

	technologyCategory := *SampleCategoryFactory()
	technologyCategory["name"] = "Technology"

	var technologyCategoryUUID uuid.UUID

	var AssertNewCategoryBodyRespDB = func(app TestApp, assert *assert.A, resp *http.Response) {
		technologyCategoryUUID = AssertCategoryWithBodyRespDBMostRecent(app, assert, resp, &technologyCategory)
	}

	TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body:   &technologyCategory,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusCreated,
			DBTester: AssertNewCategoryBodyRespDB,
		},
	)

	technologyTag := *SampleTagFactory(technologyCategoryUUID)

	var AssertUpdatedTagBodyRespDB = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertTagWithBodyRespDB(app, assert, resp, &technologyTag)
	}

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
		Body:   &technologyTag,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: AssertUpdatedTagBodyRespDB,
		},
	).Close()
}

func TestUpdateTagWorksWithSameDetails(t *testing.T) {
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(t)

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
		Body:   SampleTagFactory(categoryUUID),
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
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
		TestRequest{
			Method: fiber.MethodPatch,
			Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
			Body:   SampleTagFactory(uuid),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID).Close()
	}
}

func TestDeleteTagWorks(t *testing.T) {
	existingAppAssert, _, tagUUID := CreateSampleTag(t)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/tags/%s", tagUUID),
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusNoContent,
			DBTester: AssertNoTags,
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
		TestRequest{
			Method: fiber.MethodDelete,
			Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
		}.TestOnErrorAndDB(t, &appAssert,
			ErrorWithDBTester{
				Error:    errors.TagNotFound,
				DBTester: Assert1Tag,
			},
		)
	}

	appAssert.Close()
}

func TestDeleteTagFailsNotFound(t *testing.T) {
	appAssert, _, _ := CreateSampleTag(t)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   fmt.Sprintf("/api/v1/tags/%s", uuid.New()),
	}.TestOnErrorAndDB(t, &appAssert,
		ErrorWithDBTester{
			Error:    errors.TagNotFound,
			DBTester: Assert1Tag,
		},
	).Close()
}
