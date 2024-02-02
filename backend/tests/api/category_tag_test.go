package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/huandu/go-assert"
)

func AssertTagsWithBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, body *[]map[string]interface{}) []uuid.UUID {
	var respTags []models.Tag

	err := json.NewDecoder(resp.Body).Decode(&respTags)

	assert.NilError(err)

	var dbTags []models.Tag

	err = app.Conn.Find(&dbTags).Error

	assert.NilError(err)

	assert.Equal(len(dbTags), len(respTags))

	for i, dbTag := range dbTags {
		assert.Equal(dbTag.ID, respTags[i].ID)
		assert.Equal(dbTag.Name, respTags[i].Name)
		assert.Equal(dbTag.CategoryID, respTags[i].CategoryID)
	}

	tagIDs := make([]uuid.UUID, len(dbTags))
	for i, dbTag := range dbTags {
		tagIDs[i] = dbTag.ID
	}

	return tagIDs
}

func TestGetCategoryTagsWorks(t *testing.T) {
	appAssert, categoryUUID, tagID := CreateSampleTag(t)

	body := SampleTagFactory(categoryUUID)
	(*body)["id"] = tagID

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/categories/%s/tags", categoryUUID),
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertTagsWithBodyRespDB(app, assert, resp, &[]map[string]interface{}{*body})
			},
		},
	).Close()
}

func TestGetCategoryTagsFailsCategoryBadRequest(t *testing.T) {
	appAssert, _ := CreateSampleCategory(t, nil)

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
			Path:   fmt.Sprintf("/api/v1/categories/%s/tags", badRequest),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID)
	}

	appAssert.Close()
}

func TestGetCategoryTagsFailsCategoryNotFound(t *testing.T) {
	appAssert, _ := CreateSampleCategory(t, nil)

	uuid := uuid.New()

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/categories/%s/tags", uuid),
	}.TestOnErrorAndDB(t, &appAssert, ErrorWithDBTester{
		Error: errors.CategoryNotFound,
		DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
			var category models.Category
			err := app.Conn.Where("id = ?", uuid).First(&category).Error
			assert.Error(err)

			var respBody []map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&respBody)
			assert.NilError(err)
			assert.Equal(0, len(respBody))
		},
	}).Close()
}

func TestGetCategoryTagWorks(t *testing.T) {
	existingAppAssert, categoryUUID, tagUUID := CreateSampleTag(t)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/categories/%s/tags/%s", categoryUUID, tagUUID),
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status: fiber.StatusOK,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				AssertTagWithBodyRespDB(app, assert, resp, SampleTagFactory(categoryUUID))
			},
		},
	).Close()
}

func TestGetCategoryTagFailsCategoryBadRequest(t *testing.T) {
	appAssert, _, tagUUID := CreateSampleTag(t)

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
			Path:   fmt.Sprintf("/api/v1/categories/%s/tags/%s", badRequest, tagUUID),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID)
	}

	appAssert.Close()
}

func TestGetCategoryTagFailsTagBadRequest(t *testing.T) {
	appAssert, categoryUUID := CreateSampleCategory(t, nil)

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
			Path:   fmt.Sprintf("/api/v1/categories/%s/tags/%s", categoryUUID, badRequest),
		}.TestOnError(t, &appAssert, errors.FailedToValidateID)
	}

	appAssert.Close()
}

func TestGetCategoryTagFailsCategoryNotFound(t *testing.T) {
	appAssert, _, tagUUID := CreateSampleTag(t)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/categories/%s/tags/%s", uuid.New(), tagUUID),
	}.TestOnError(t, &appAssert, errors.TagNotFound).Close()
}

func TestGetCategoryTagFailsTagNotFound(t *testing.T) {
	appAssert, categoryUUID := CreateSampleCategory(t, nil)

	TestRequest{
		Method: fiber.MethodGet,
		Path:   fmt.Sprintf("/api/v1/categories/%s/tags/%s", categoryUUID, uuid.New()),
	}.TestOnError(t, &appAssert, errors.TagNotFound).Close()
}
