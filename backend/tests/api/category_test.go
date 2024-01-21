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

func CreateSampleCategory(t *testing.T, categoryName string) ExistingAppAssert {
	return TestRequest{
		Method: "POST",
		Path:   "/api/v1/categories/",
		Body: &map[string]interface{}{
			"category_name": categoryName,
		},
	}.TestOnStatusAndDBKeepDB(t, nil,
		DBTesterWithStatus{
			Status: 201,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {

				var respCategory models.Category

				err := json.NewDecoder(resp.Body).Decode(&respCategory)

				assert.NilError(err)

				dbCategory, err := transactions.GetCategory(app.Conn, respCategory.ID)

				assert.NilError(err)

				assert.Equal(dbCategory, &respCategory)
			},
		},
	)
}

func TestCreateCategoryWorks(t *testing.T) {
	CreateSampleCategory(t, "Science")
}

func TestCreateCategoryIgnoresid(t *testing.T) {
	TestRequest{
		Method: "POST",
		Path:   "/api/v1/categories/",
		Body: &map[string]interface{}{
			"id":            12,
			"category_name": "Science",
		},
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status: 201,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var respCategory models.Category

				err := json.NewDecoder(resp.Body).Decode(&respCategory)

				assert.NilError(err)

				dbCategory, err := transactions.GetCategory(app.Conn, respCategory.ID)

				assert.NilError(err)

				assert.NotEqual(12, dbCategory.ID)
			},
		},
	)
}

func AssertNoCategoryCreation(app TestApp, assert *assert.A, resp *http.Response) {
	AssertNumCategoriesRemainsAtN(app, assert, resp, 0)
}

func AssertNumCategoriesRemainsAtN(app TestApp, assert *assert.A, resp *http.Response, n int) {
	var categories []models.Category

	err := app.Conn.Find(&categories).Error

	assert.NilError(err)

	assert.Equal(n, len(categories))
}

func TestCreateCategoryFailsIfNameIsNotString(t *testing.T) {
	TestRequest{
		Method: "POST",
		Path:   "/api/v1/categories/",
		Body: &map[string]interface{}{
			"category_name": 1231,
		},
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  400,
				Message: "failed to process the request",
			},
			DBTester: AssertNoCategoryCreation,
		},
	)
}

func TestCreateCategoryFailsIfNameIsMissing(t *testing.T) {
	TestRequest{
		Method: "POST",
		Path:   "/api/v1/categories/",
		Body:   &map[string]interface{}{},
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  400,
				Message: "failed to validate the data",
			},
			DBTester: AssertNoCategoryCreation,
		},
	)
}

func TestCreateCategoryFailsIfCategoryWithThatNameAlreadyExists(t *testing.T) {
	categoryName := "Science"

	existingAppAssert := CreateSampleCategory(t, categoryName)

	var TestNumCategoriesRemainsAt1 = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertNumCategoriesRemainsAtN(app, assert, resp, 1)
	}

	for _, permutation := range AllCasingPermutations(categoryName) {
		fmt.Println(permutation)
		TestRequest{
			Method: "POST",
			Path:   "/api/v1/categories/",
			Body: &map[string]interface{}{
				"category_name": permutation,
			},
		}.TestOnStatusMessageAndDBKeepDB(t, &existingAppAssert,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400,
					Message: "category with that name already exists",
				},
				DBTester: TestNumCategoriesRemainsAt1,
			},
		)
	}
}
