package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-assert"

	"github.com/goccy/go-json"
)

func SampleTagFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"name":        "Generate",
		"category_id": 1,
	}
}

func AssertTagWithIDBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, id uint, body *map[string]interface{}) {
	var respTag models.Tag

	err := json.NewDecoder(resp.Body).Decode(&respTag)

	assert.NilError(err)

	var dbTag models.Tag

	err = app.Conn.First(&dbTag, id).Error

	assert.NilError(err)

	assert.Equal(dbTag.ID, respTag.ID)
	assert.Equal(dbTag.Name, respTag.Name)
	assert.Equal(dbTag.CategoryID, respTag.CategoryID)

	assert.Equal((*body)["name"].(string), dbTag.Name)
	assert.Equal((*body)["category_id"].(int), int(dbTag.CategoryID))
}

func AssertSampleTagBodyRespDB(app TestApp, assert *assert.A, resp *http.Response) {
	AssertTagWithIDBodyRespDB(app, assert, resp, 1, SampleTagFactory())
}

func CreateSampleTag(t *testing.T) ExistingAppAssert {
	appAssert := CreateSampleCategory(t)

	return TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/tags/",
		Body:   SampleTagFactory(),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusCreated,
			DBTester: AssertSampleTagBodyRespDB,
		},
	)
}

func TestCreateTagWorks(t *testing.T) {
	CreateSampleTag(t).Close()
}

var AssertNoTags = func(app TestApp, assert *assert.A, resp *http.Response) {
	var tags []models.Tag

	err := app.Conn.Find(&tags).Error

	assert.NilError(err)

	assert.Equal(0, len(tags))
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
		}.TestOnStatusMessageAndDB(t, nil,
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
			"category_id": 1,
		},
		{},
	}

	for _, badBody := range badBodys {
		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/tags/",
			Body:   &badBody,
		}.TestOnStatusMessageAndDB(t, nil,
			ErrorWithDBTester{
				Error:    errors.FailedToValidateTag,
				DBTester: AssertNoTags,
			},
		).Close()
	}
}

func TestGetTagWorks(t *testing.T) {
	existingAppAssert := CreateSampleTag(t)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   "/api/v1/tags/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: AssertSampleTagBodyRespDB,
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
		Path:   "/api/v1/tags/1",
	}.TestOnError(t, nil, errors.TagNotFound).Close()
}

func TestUpdateTagWorksUpdateName(t *testing.T) {
	existingAppAssert := CreateSampleTag(t)

	generateNUTag := *SampleTagFactory()
	generateNUTag["name"] = "GenerateNU"

	var AssertUpdatedTagBodyRespDB = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertTagWithIDBodyRespDB(app, assert, resp, 1, &generateNUTag)
	}

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   "/api/v1/tags/1",
		Body:   &generateNUTag,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: AssertUpdatedTagBodyRespDB,
		},
	).Close()
}

func TestUpdateTagWorksUpdateCategory(t *testing.T) {
	existingAppAssert := CreateSampleTag(t)

	technologyCategory := *SampleCategoryFactory()
	technologyCategory["name"] = "Technology"

	var AssertNewCategoryBodyRespDB = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertCategoryWithIDBodyRespDB(app, assert, resp, 2, &technologyCategory)
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

	technologyTag := *SampleTagFactory()
	technologyTag["category_id"] = 2

	var AssertUpdatedTagBodyRespDB = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertTagWithIDBodyRespDB(app, assert, resp, 1, &technologyTag)
	}

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   "/api/v1/tags/1",
		Body:   &technologyTag,
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: AssertUpdatedTagBodyRespDB,
		},
	).Close()
}

func TestUpdateTagWorksWithSameDetails(t *testing.T) {
	existingAppAssert := CreateSampleTag(t)

	TestRequest{
		Method: fiber.MethodPatch,
		Path:   "/api/v1/tags/1",
		Body:   SampleTagFactory(),
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusOK,
			DBTester: AssertSampleTagBodyRespDB,
		},
	).Close()
}

func TestUpdateTagFailsBadRequest(t *testing.T) {
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
			Body:   SampleTagFactory(),
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestDeleteTagWorks(t *testing.T) {
	existingAppAssert := CreateSampleTag(t)

	TestRequest{
		Method: fiber.MethodDelete,
		Path:   "/api/v1/tags/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   fiber.StatusNoContent,
			DBTester: AssertNoTags,
		},
	).Close()
}

func TestDeleteTagFailsBadRequest(t *testing.T) {
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
		}.TestOnError(t, nil, errors.FailedToValidateID).Close()
	}
}

func TestDeleteTagFailsNotFound(t *testing.T) {
	TestRequest{
		Method: fiber.MethodDelete,
		Path:   "/api/v1/tags/1",
	}.TestOnError(t, nil, errors.TagNotFound).Close()
}
