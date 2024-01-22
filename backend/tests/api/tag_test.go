package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/huandu/go-assert"

	"github.com/goccy/go-json"
)

var AssertRespTagSameAsDBTag = func(app TestApp, assert *assert.A, resp *http.Response) {
	var respTag models.Tag

	err := json.NewDecoder(resp.Body).Decode(&respTag)

	assert.NilError(err)

	dbTag, err := transactions.GetTag(app.Conn, respTag.ID)

	assert.NilError(err)

	assert.Equal(dbTag, &respTag)
}

func CreateSampleTag(t *testing.T, tagName string, categoryName string, existingAppAssert *ExistingAppAssert) ExistingAppAssert {
	appAssert := CreateSampleCategory(t, categoryName, existingAppAssert)

	return TestRequest{
		Method: "POST",
		Path:   "/api/v1/tags/",
		Body: &map[string]interface{}{
			"name":        tagName,
			"category_id": 1,
		},
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status:   201,
			DBTester: AssertRespTagSameAsDBTag,
		},
	)
}

func TestCreateTagWorks(t *testing.T) {
	CreateSampleTag(t, "Generate", "Science", nil).Close()
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
			Method: "POST",
			Path:   "/api/v1/tags/",
			Body:   &badBody,
		}.TestOnStatusMessageAndDB(t, nil,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400,
					Message: "failed to process the request",
				},
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
			Method: "POST",
			Path:   "/api/v1/tags/",
			Body:   &badBody,
		}.TestOnStatusMessageAndDB(t, nil,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400,
					Message: "failed to validate the data",
				},
				DBTester: AssertNoTags,
			},
		).Close()
	}
}

func TestGetTagWorks(t *testing.T) {
	existingAppAssert := CreateSampleTag(t, "Generate", "Science", nil)

	TestRequest{
		Method: "GET",
		Path:   "/api/v1/tags/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   200,
			DBTester: AssertRespTagSameAsDBTag,
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
			Method: "GET",
			Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: "failed to validate id",
			},
		).Close()
	}
}

func TestGetTagFailsNotFound(t *testing.T) {
	TestRequest{
		Method: "GET",
		Path:   "/api/v1/tags/1",
	}.TestOnStatusAndMessage(t, nil,
		MessageWithStatus{
			Status:  404,
			Message: "failed to find tag",
		},
	).Close()
}

func TestUpdateTagWorksUpdateName(t *testing.T) {
	existingAppAssert := CreateSampleTag(t, "Generate", "Science", nil)

	TestRequest{
		Method: "PATCH",
		Path:   "/api/v1/tags/1",
		Body: &map[string]interface{}{
			"name":        "GenerateNU",
			"category_id": 1,
		},
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   200,
			DBTester: AssertRespTagSameAsDBTag,
		},
	).Close()
}

func TestUpdateTagWorksUpdateCategory(t *testing.T) {
	existingAppAssert := CreateSampleTag(t, "Generate", "Science", nil)
	existingAppAssert = CreateSampleCategory(t, "Technology", &existingAppAssert)

	TestRequest{
		Method: "PATCH",
		Path:   "/api/v1/tags/1",
		Body: &map[string]interface{}{
			"name":        "Generate",
			"category_id": 2,
		},
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   200,
			DBTester: AssertRespTagSameAsDBTag,
		},
	).Close()
}

func TestUpdateTagWorksWithSameDetails(t *testing.T) {
	existingAppAssert := CreateSampleTag(t, "Generate", "Science", nil)

	TestRequest{
		Method: "PATCH",
		Path:   "/api/v1/tags/1",
		Body: &map[string]interface{}{
			"name":        "Generate",
			"category_id": 1,
		},
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   200,
			DBTester: AssertRespTagSameAsDBTag,
		},
	).Close()
}

func TestUpdateTagFailsBadRequest(t *testing.T) {
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
			Method: "PATCH",
			Path:   "/api/v1/tags/1",
			Body:   &badBody,
		}.TestOnStatusMessageAndDB(t, nil,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400,
					Message: "failed to process the request",
				},
				DBTester: AssertNoTags,
			},
		).Close()
	}
}

func TestDeleteTagWorks(t *testing.T) {
	existingAppAssert := CreateSampleTag(t, "Generate", "Science", nil)

	TestRequest{
		Method: "DELETE",
		Path:   "/api/v1/tags/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   204,
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
			Method: "DELETE",
			Path:   fmt.Sprintf("/api/v1/tags/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: "failed to validate id",
			},
		).Close()
	}
}

func TestDeleteTagFailsNotFound(t *testing.T) {
	TestRequest{
		Method: "DELETE",
		Path:   "/api/v1/tags/1",
	}.TestOnStatusAndMessage(t, nil,
		MessageWithStatus{
			Status:  404,
			Message: "failed to find tag",
		},
	).Close()
}
