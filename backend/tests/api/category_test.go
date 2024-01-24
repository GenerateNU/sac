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

func TestGetCategoryWorks(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t, "Science", nil)

	TestRequest{
		Method: "GET",
		Path:   "/api/v1/categories/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   200,
			DBTester: AssertRespCategorySameAsDBCategory,
		},
	).Close()
}

func TestGetCategoryFailsBadRequest(t *testing.T) {
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
			Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: "failed to validate id",
			},
		).Close()
	}
}

func TestGetCategoryFailsNotFound(t *testing.T) {
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

func TestGetCategoriesWork(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t, "Science", nil)

	TestRequest{
		Method: "GET",
		Path:   "/api/v1/categories/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   200,
			DBTester: AssertRespCategorySameAsDBCategory,
		},
	).Close()
}

func TestUpdateCategoryWorksUpdateName(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t, "Science", nil)

	TestRequest{
		Method: "PATCH",
		Path:   "/api/v1/categories/1",
		Body: &map[string]interface{}{
			"category_name": "History",
		},
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status: 200,
			DBTester: func(app TestApp, assert *assert.A, resp *http.Response) {
				var category models.Category

				err := json.NewDecoder(resp.Body).Decode(&category)

				assert.NilError(err)

				// get the user directly from
				dbCategory, err := transactions.GetCategory(app.Conn, category.ID)

				assert.NilError(err)

				// assert that the user returned from the database is the same as the user returned from the API
				assert.Equal(dbCategory.Name, category.Name)
			},
		},
	).Close()
}

func TestUpdateCategoryWorksWithSameDetails(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t, "Science", nil)

	TestRequest{
		Method: "PATCH",
		Path:   "/api/v1/categories/1",
		Body: &map[string]interface{}{
			"category_name": "Science",
		},
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   200,
			DBTester: AssertRespCategorySameAsDBCategory,
		},
	).Close()
}

func TestUpdateCategoryFailsWithExistingName(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t, "Science", nil)
	CreateSampleCategory(t, "History", &existingAppAssert)

	TestRequest{
		Method: "PATCH",
		Path:   "/api/v1/categories/1",
		Body: &map[string]interface{}{
			"category_name": "History",
		},
	}.TestOnStatusAndMessage(t, &existingAppAssert,
		MessageWithStatus{
			Status:  400,
			Message: "category with that name already exists",
		},
	).Close()
}

func TestUpdateCategoryBadRequest(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t, "Science", nil)

	badBodys := []map[string]interface{}{
		{
			"name": 1,
		},
		{
			"name": models.Category{},
		},
	}

	for _, badBody := range badBodys {
		TestRequest{
			Method: "PATCH",
			Path:   "/api/v1/categories/1",
			Body:   &badBody,
		}.TestOnStatusMessageAndDB(t, &existingAppAssert,
			StatusMessageDBTester{
				MessageWithStatus: MessageWithStatus{
					Status:  400,
					Message: "failed to validate the data",
				},
				DBTester: AssertNoTags,
			},
		)
	}

	existingAppAssert.Close()
}

func TestDeleteCategoryWorks(t *testing.T) {
	existingAppAssert := CreateSampleCategory(t, "Science", nil)

	TestRequest{
		Method: "DELETE",
		Path:   "/api/v1/categories/1",
	}.TestOnStatusAndDB(t, &existingAppAssert,
		DBTesterWithStatus{
			Status:   204,
			DBTester: AssertNoTags,
		},
	).Close()
}

func TestDeleteCategoryFailsBadRequest(t *testing.T) {
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
			Path:   fmt.Sprintf("/api/v1/categories/%s", badRequest),
		}.TestOnStatusAndMessage(t, nil,
			MessageWithStatus{
				Status:  400,
				Message: "failed to validate id",
			},
		).Close()
	}
}

func TestDeleteCategoryFailsNotFound(t *testing.T) {
	TestRequest{
		Method: "DELETE",
		Path:   "/api/v1/categories/1",
	}.TestOnStatusAndMessage(t, nil,
		MessageWithStatus{
			Status:  404,
			Message: "failed to find category",
		},
	).Close()
}
