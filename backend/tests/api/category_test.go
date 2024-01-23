package tests

import (
	"net/http"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-assert"

	"github.com/goccy/go-json"
)

func SampleCategoryFactory() *map[string]interface{} {
	return &map[string]interface{}{
		"category_name": "Foo",
	}
}

func AssertCategoryWithIDBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, id uint, body *map[string]interface{}) {
	var respCategory models.Category

	err := json.NewDecoder(resp.Body).Decode(&respCategory)

	assert.NilError(err)

	var dbCategory models.Category

	err = app.Conn.First(&dbCategory, id).Error

	assert.NilError(err)

	assert.Equal(dbCategory.ID, respCategory.ID)
	assert.Equal(dbCategory.Name, respCategory.Name)

	assert.Equal((*body)["category_name"].(string), dbCategory.Name)

}

func AssertSampleCategoryBodyRespDB(app TestApp, assert *assert.A, resp *http.Response) {
	AssertCategoryWithIDBodyRespDB(app, assert, resp, 1, SampleCategoryFactory())
}

func CreateSampleCategory(t *testing.T) ExistingAppAssert {
	return TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body:   SampleCategoryFactory(),
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status:   201,
			DBTester: AssertSampleCategoryBodyRespDB,
		},
	)
}

func TestCreateCategoryWorks(t *testing.T) {
	CreateSampleCategory(t).Close()
}

func TestCreateCategoryIgnoresid(t *testing.T) {
	TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/categories/",
		Body: &map[string]interface{}{
			"id":            12,
			"category_name": "Foo",
		},
	}.TestOnStatusAndDB(t, nil,
		DBTesterWithStatus{
			Status:   201,
			DBTester: AssertSampleCategoryBodyRespDB,
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
		ErrorWithDBTester{
			Error:    errors.FailedToValidateCategory,
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
		ErrorWithDBTester{
			Error:    errors.FailedToValidateCategory,
			DBTester: AssertNoCategories,
		},
	).Close()
}

func TestCreateCategoryFailsIfCategoryWithThatNameAlreadyExists(t *testing.T) {

	existingAppAssert := CreateSampleCategory(t)

	var TestNumCategoriesRemainsAt1 = func(app TestApp, assert *assert.A, resp *http.Response) {
		AssertNumCategoriesRemainsAtN(app, assert, resp, 1)
	}

	for _, permutation := range AllCasingPermutations((*SampleCategoryFactory())["category_name"].(string)) {
		modifiedSampleCategoryBody := *SampleCategoryFactory()
		modifiedSampleCategoryBody["category_name"] = permutation

		TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/categories/",
			Body:   &modifiedSampleCategoryBody,
		}.TestOnStatusMessageAndDB(t, &existingAppAssert,
			ErrorWithDBTester{
				Error:    errors.CategoryAlreadyExists,
				DBTester: TestNumCategoriesRemainsAt1,
			},
		)
	}

	existingAppAssert.Close()
}
