package tests

import (
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-assert"

	"github.com/goccy/go-json"
)

var AssertRespCategorySameAsDBCategory = func(app TestApp, assert *assert.A, resp *http.Response) {
	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	dbCategory, err := transactions.GetCategory(app.Conn, respCategory.ID)

	assert.NilError(&err)

	assert.Equal(dbCategory, &respCategory)
}

func CreateSampleCategory(t *testing.T, categoryName string, existingAppAssert *ExistingAppAssert) ExistingAppAssert {
	return TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body: &map[string]interface{}{
			"category_name": categoryName,
		},
	}.TestOnStatusAndDB(t, existingAppAssert,
		DBTesterWithStatus{
			Status:   201,
			DBTester: AssertRespCategorySameAsDBCategory,
		},
	)
}

func TestCreateCategoryWorks(t *testing.T) {
	CreateSampleCategory(t, "Science", nil).Close()
}

func TestCreateCategoryIgnoresid(t *testing.T) {
	TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body: &map[string]interface{}{
			"id":            12,
			"category_name": "Science",
		},
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status:   201,
			DBTester: AssertRespCategorySameAsDBCategory,
		},
	).Close()
}

func AssertNoCategories(app TestApp, assert *assert.A, resp *http.Response) {
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
		Method: fiber.MethodPost,
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
			DBTester: AssertNoCategories,
		},
	).Close()
}

func TestCreateCategoryFailsIfNameIsMissing(t *testing.T) {
	TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body:   &map[string]interface{}{},
	}.TestOnStatusMessageAndDB(t, nil,
		StatusMessageDBTester{
			MessageWithStatus: MessageWithStatus{
				Status:  400,
				Message: "failed to validate category",
			},
			DBTester: AssertNoCategories,
		},
	).Close()
}

func TestCreateCategoryFailsIfCategoryWithThatNameAlreadyExists(t *testing.T) {
	categoryName := "foo"

	existingAppAssert := CreateSampleCategory(t, categoryName, nil)

	var TestNumCategoriesRemainsAt1 = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertNumCategoriesRemainsAtN(app, assert, resp, 1)
	}

	for _, permutation := range AllCasingPermutations(categoryName) {
		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body: &map[string]interface{}{
				"category_name": permutation,
			},
		}.TestOnStatusMessageAndDB(t, &existingAppAssert,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400,
					Message: "failed to validate ID",
				},
				DBTester: TestNumCategoriesRemainsAt1,
			},
		)
	}

	existingAppAssert.Close()
}
