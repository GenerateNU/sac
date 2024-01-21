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

func CreateSampleTag(t *testing.T, tagName string, categoryName string, existingAppAssert *ExistingAppAssert) ExistingAppAssert {
	var appAssert ExistingAppAssert

	if existingAppAssert == nil {
		appAssert = CreateSampleCategory(t, categoryName, existingAppAssert)
	} else {
		appAssert = CreateSampleCategory(t, categoryName, existingAppAssert)
	}
	return TestRequest{
		Method: "POST",
		Path:   "/api/v1/tags/",
		Body: &map[string]interface{}{
			"name":        tagName,
			"category_id": 1,
		},
	}.TestOnStatusAndDB(t, &appAssert,
		DBTesterWithStatus{
			Status: 201,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respTag models.Tag

				err := json.NewDecoder(resp.Body).Decode(&respTag)

				assert.NilError(err)

				dbTag, err := transactions.GetTag(app.Conn, respTag.ID)

				assert.NilError(err)

				assert.Equal(dbTag, &respTag)
			},
		},
	)
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
		)
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
		)
	}
}

func TestGetTagWorks(t *testing.T) {
	existingAppAssert := CreateSampleTag(t, "Generate", "Science", nil)

	TestRequest{
		Method: "GET",
		Path:   "/api/v1/tags/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respTag models.Tag

				err := json.NewDecoder(resp.Body).Decode(&respTag)

				assert.NilError(err)

				dbTag, err := transactions.GetTag(app.Conn, respTag.ID)

				assert.NilError(err)

				assert.Equal(dbTag, &respTag)
			},
		},
	)
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
		)
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
	)
}
func TestDeleteTagWorks(t *testing.T) {
	existingAppAssert := CreateSampleTag(t, "Generate", "Science")

	TestRequest{
		Method: "DELETE",
		Path:   "/api/v1/tags/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   204,
			DBTester: AssertNoTags,
		},
	)

	existingAppAssert.App.DropDB()
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
		)
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
	)
}
