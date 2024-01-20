package tests

import (
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/huandu/go-assert"

	"github.com/goccy/go-json"
)

func TestCreateTagWorks(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t, "Science")
	TestRequest{
		Method: "POST",
		Path:   "/api/v1/tags/",
		Body: &map[string]interface{}{
			"name":        "Generate",
			"category_id": 1,
		},
	}.TestOnStatusAndDB(t, &existingAppAssert,
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

var AssertNoTagCreation = func(app TestApp, assert *assert.A, resp *http.Response) {
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
				DBTester: AssertNoTagCreation,
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
				DBTester: AssertNoTagCreation,
			},
		)
	}

}
